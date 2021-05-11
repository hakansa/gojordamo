package sqlstore

import (
	"github.com/blang/semver"
	"github.com/jmoiron/sqlx"
)

// Migration declares the SQL migrations for updates
type Migration struct {
	fromVersion   semver.Version
	toVersion     semver.Version
	migrationFunc func(sqlx.Ext, *SQLStore) error
}

var migrations = []Migration{
	{
		fromVersion: semver.MustParse("0.0.0"),
		toVersion:   semver.MustParse("0.1.0"),
		migrationFunc: func(e sqlx.Ext, sqlStore *SQLStore) error {

			return nil
		},
	},
}
