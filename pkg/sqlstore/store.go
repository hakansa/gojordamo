package sqlstore

import (
	"strings"

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

// parseDSN splits up a connection string into a driver name and data source name.
//
// For example:
//	mysql://gojordomo:pass@localhost:5432/gojordomo
// returns
//	driverName = mysql
//	dataSourceName = gojordomo:pass@localhost:5432/gojordomo
//
// By contrast, a Postgres DSN is returned unmodified.
func parseDSN(dsn string) (string, string, error) {
	// Treat the DSN as the URL that it is.
	s := strings.SplitN(dsn, "://", 2)
	if len(s) != 2 {
		return "", "", errors.New("failed to parse DSN as URL")
	}

	scheme := s[0]
	switch scheme {
	case "mysql":
		// Strip off the mysql:// for the dsn with which to connect.
		dsn = s[1]

	case "postgres":
		// No changes required

	default:
		return "", "", errors.Errorf("unsupported scheme %s", scheme)
	}

	return scheme, dsn, nil
}
