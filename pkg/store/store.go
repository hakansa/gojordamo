package store

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/influxdata/influxdb/kit/errors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// SQLStore ..
type SQLStore struct {
	cfg     Config
	db      *sqlx.DB
	builder sq.StatementBuilderType
}

// New constructs a new instance of SQLStore.
func New(cfg Config) (*SQLStore, error) {

	db, err := sqlx.Open(string(cfg.Driver), cfg.DataSource)
	if err != nil {
		return nil, errors.Wrap("unable to open db connection", err)
	}

	builder := sq.StatementBuilder.PlaceholderFormat(sq.Question)

	switch cfg.Driver {
	case DBDriverMySQL: // mysql
		db.MapperFunc(func(s string) string { return s })

	case DBDriverPostgres: // postgres
		builder = builder.PlaceholderFormat(sq.Dollar)

	}

	return &SQLStore{
		cfg:     cfg,
		db:      db,
		builder: builder,
	}, nil
}

// finalizeTransaction ensures a transaction is closed after use, rolling back if not already committed.
func (sqlStore *SQLStore) finalizeTransaction(tx *sqlx.Tx) {
	// Rollback returns sql.ErrTxDone if the transaction was already closed.
	if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
		logrus.Errorf("Failed to rollback transaction; err: %v", err)
	}
}
