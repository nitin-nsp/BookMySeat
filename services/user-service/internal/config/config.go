package config

type Config struct {
	Name     string
	Host     string
	Port     int
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

type JWTConfig struct {
	Secret string
	Expiry string
}
