package config

type DatabaseDriver string

var (
	DatabaseDriverMemory   DatabaseDriver = "memory"
	DatabaseDriverPostgres DatabaseDriver = "postgres"
	DatabaseDriverMongo    DatabaseDriver = "mongo"
)

type DatabaseConfig struct {
	Driver DatabaseDriver `env:"DATABASE_DRIVER" yaml:"driver"`
}

type DatabasePostgresConfig struct {
	Host   string `env:"POSTGRES_HOST" yaml:"host"`
	Port   string `env:"POSTGRES_PORT" yaml:"port"`
	User   string `env:"POSTGRES_USER" yaml:"user"`
	Pass   string `env:"POSTGRES_PASS" yaml:"pass"`
	DbName string `env:"POSTGRES_DBNAME" yaml:"dbname"`
}

type DatabaseMongoConfig struct {
	Host   string `env:"MONGO_HOST" yaml:"host"`
	Port   string `env:"MONGO_PORT" yaml:"port"`
	User   string `env:"MONGO_USER" yaml:"user"`
	Pass   string `env:"MONGO_PASS" yaml:"pass"`
	DbName string `env:"MONGO_DBNAME" yaml:"dbname"`
}
