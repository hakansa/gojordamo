package sqlstore

import (
	"github.com/blang/semver"
	"github.com/pkg/errors"
)

// Migrate runs migrate function for latest version upgrade
func (s *SQLStore) Migrate(originalSchemaVersion semver.Version) error {
	currentSchemaVersion := originalSchemaVersion
	for _, migration := range migrations {
		if !currentSchemaVersion.EQ(migration.fromVersion) {
			continue
		}

		if err := s.migrate(migration); err != nil {
			return err
		}

		currentSchemaVersion = migration.toVersion
	}

	return nil
}

func (s *SQLStore) migrate(migration Migration) (err error) {
	tx, err := sqlStore.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "could not begin transaction")
	}
	defer s.finalizeTransaction(tx)

	if err := migration.migrationFunc(tx, s); err != nil {
		return errors.Wrapf(err, "error executing migration from version %s to version %s", migration.fromVersion.String(), migration.toVersion.String())
	}

	if err := s.SetCurrentVersion(tx, migration.toVersion); err != nil {
		return errors.Wrapf(err, "failed to set the current version to %s", migration.toVersion.String())
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "could not commit transaction")
	}
	return nil
}

// RunMigrations will run the migrations (if any).
func (s *SQLStore) RunMigrations() error {
	currentSchemaVersion, err := s.GetCurrentVersion()
	if err != nil {
		return errors.Wrapf(err, "failed to get the current schema version")
	}

	if currentSchemaVersion.LT(LatestVersion()) {
		if err := s.Migrate(currentSchemaVersion); err != nil {
			return errors.Wrapf(err, "failed to complete migrations")
		}
	}

	return nil
}
