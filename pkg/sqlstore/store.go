package sqlstore

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
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

		//case DBDriverPostgres: // postgres
		//	builder = builder.PlaceholderFormat(sq.Dollar)

	}

	return &SQLStore{
		cfg:     cfg,
		db:      db,
		builder: builder,
	}, nil
}

// queryer is an interface describing a resource that can query.
//
// It exactly matches sqlx.Queryer, existing simply to constrain sqlx usage to this file.
type queryer interface {
	sqlx.Queryer
}

// builder is an interface describing a resource that can construct SQL and arguments.
//
// It exists to allow consuming any squirrel.*Builder type.
type builder interface {
	ToSql() (string, []interface{}, error)
}

// get queries for a single row, building the sql, and writing the result into dest.
//
// Use this to simplify querying for a single row or column. Dest may be a pointer to a simple
// type, or a struct with fields to be populated from the returned columns.
func (sqlStore *SQLStore) getBuilder(q sqlx.Queryer, dest interface{}, b builder) error {
	sqlString, args, err := b.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build sql")
	}

	sqlString = sqlStore.db.Rebind(sqlString)

	return sqlx.Get(q, dest, sqlString, args...)
}

// selectBuilder queries for one or more rows, building the sql, and writing the result into dest.
//
// Use this to simplify querying for multiple rows (and possibly columns). Dest may be a slice of
// a simple, or a slice of a struct with fields to be populated from the returned columns.
func (sqlStore *SQLStore) selectBuilder(q sqlx.Queryer, dest interface{}, b builder) error {
	sqlString, args, err := b.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build sql")
	}

	sqlString = sqlStore.db.Rebind(sqlString)

	return sqlx.Select(q, dest, sqlString, args...)
}

// execer is an interface describing a resource that can execute write queries.
//
// It allows the use of *sqlx.Db and *sqlx.Tx.
type execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	DriverName() string
}

type queryExecer interface {
	queryer
	execer
}

// exec executes the given query using positional arguments, automatically rebinding for the db.
func (sqlStore *SQLStore) exec(e execer, sqlString string, args ...interface{}) (sql.Result, error) {
	sqlString = sqlStore.db.Rebind(sqlString)
	return e.Exec(sqlString, args...)
}

// exec executes the given query, building the necessary sql.
func (sqlStore *SQLStore) execBuilder(e execer, b builder) (sql.Result, error) {
	sqlString, args, err := b.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build sql")
	}

	return sqlStore.exec(e, sqlString, args...)
}

// finalizeTransaction ensures a transaction is closed after use, rolling back if not already committed.
func (sqlStore *SQLStore) finalizeTransaction(tx *sqlx.Tx) {
	// Rollback returns sql.ErrTxDone if the transaction was already closed.
	if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
		logrus.Errorf("Failed to rollback transaction; err: %v", err)
	}
}
