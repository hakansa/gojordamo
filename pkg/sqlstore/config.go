package sqlstore

// DBDriver declares the db drivers
type DBDriver string

const (
	// DBDriverMySQL is an enum for mysql driver
	DBDriverMySQL DBDriver = "mysql"
	// DBDriverPostgres is an enum for postgres driver
	//DBDriverPostgres = "postgres"
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

type SqlSettings struct {
	DriverName                  string
	DataSource                  string
	MaxIdleConns                int
	ConnMaxLifetimeMilliseconds int
	ConnMaxIdleTimeMilliseconds int
	MaxOpenConns                int
	QueryTimeout                int
}

func (s *SqlSettings) SetDefaults(isUpdate bool) {
	if s.DriverName == "" {
		s.DriverName = string(DBDriverMySQL)
	}

	if s.DataSource == "" {
		s.DataSource = "" // TODO: DEFAULT DATASOURCE
	}

	if s.MaxIdleConns == 0 {
		s.MaxIdleConns = 20
	}

	if s.MaxOpenConns == 0 {
		s.MaxOpenConns = 300
	}

	if s.ConnMaxLifetimeMilliseconds == 0 {
		s.ConnMaxLifetimeMilliseconds = 3600000
	}

	if s.ConnMaxIdleTimeMilliseconds == 0 {
		s.ConnMaxIdleTimeMilliseconds = 300000
	}

	if s.QueryTimeout == 0 {
		s.QueryTimeout = 30
	}

}
