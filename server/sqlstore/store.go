package sqlstore

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type SQLStore struct {
	db      *sqlx.DB
	builder sq.StatementBuilderType
}

// New constructs a new instance of SQLStore.
func New() (*SQLStore, error) {
	var db *sqlx.DB

	db = sqlx.NewDb(origDB, pluginAPI.Store.DriverName())

	builder := sq.StatementBuilder.PlaceholderFormat(sq.Question)
	if store.DriverName() == model.DATABASE_DRIVER_POSTGRES {
		builder = builder.PlaceholderFormat(sq.Dollar)
	}

	if store.DriverName() == model.DATABASE_DRIVER_MYSQL {
		db.MapperFunc(func(s string) string { return s })
	}

	return &SQLStore{
		log,
		db,
		builder,
	}, nil
}
