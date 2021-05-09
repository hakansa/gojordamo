package store

// DBDriver declares the db drivers
type DBDriver string

const (
	// DBDriverMySQL is an enum for mysql driver
	DBDriverMySQL DBDriver = "mysql"
	// DBDriverPostgres is an enum for postgres driver
	DBDriverPostgres = "postgres"
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
