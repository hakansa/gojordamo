package sqlstore

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/influxdata/influxdb/kit/errors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
