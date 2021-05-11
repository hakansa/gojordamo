package sqlstore

// DBDriver declares the db drivers
type DBDriver string

const (
	// DBDriverMySQL is an enum for mysql driver
	DBDriverMySQL DBDriver = "mysql"
	// DBDriverPostgres is an enum for postgres driver
	//DBDriverPostgres = "postgres"
)

// Default SQLSettings
const (
	DefaultDBDriver                    = DBDriverMySQL
	DefaultDataSource                  = ""
	DefaultMaxIdleConns                = 20
	DefaultMaxOpenConns                = 300
	DefaultConnMaxLifetimeMilliseconds = 3600000
	DefaultConnMaxIdleTimeMilliseconds = 300000
	DefaultQueryTimeout                = 30
)

// MaxWriteLength defines the maximum length accepted for write to the Configurations or
// ConfigurationFiles table.
//
// It is imposed by MySQL's default max_allowed_packet value of 4Mb.
const MaxWriteLength = 4 * 1024 * 1024

// Config declares the config parameters for store
type Config struct {
	Driver     DBDriver
	DataSource string
}

// SQLSettings declares the sql database settings
type SQLSettings struct {
	DriverName                  string
	DataSource                  string
	MaxIdleConns                int
	ConnMaxLifetimeMilliseconds int
	ConnMaxIdleTimeMilliseconds int
	MaxOpenConns                int
	QueryTimeout                int
}

// SetDefaults sets the default settings if they are not setted
func (s *SQLSettings) SetDefaults(isUpdate bool) {
	if s.DriverName == "" {
		s.DriverName = string(DBDriverMySQL)
	}

	if s.DataSource == "" {
		s.DataSource = DefaultDataSource
	}

	if s.MaxIdleConns == 0 {
		s.MaxIdleConns = DefaultMaxIdleConns
	}

	if s.MaxOpenConns == 0 {
		s.MaxOpenConns = DefaultMaxOpenConns
	}

	if s.ConnMaxLifetimeMilliseconds == 0 {
		s.ConnMaxLifetimeMilliseconds = DefaultConnMaxLifetimeMilliseconds
	}

	if s.ConnMaxIdleTimeMilliseconds == 0 {
		s.ConnMaxIdleTimeMilliseconds = DefaultConnMaxIdleTimeMilliseconds
	}

	if s.QueryTimeout == 0 {
		s.QueryTimeout = DefaultQueryTimeout
	}

}
