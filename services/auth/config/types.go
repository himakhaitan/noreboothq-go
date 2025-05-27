package config

type AuthServiceConfig struct {
	Server ServerConfig   `koanf:"server"`
	JWT    JWTConfig      `koanf:"jwt"`
	Log    LogConfig      `koanf:"logging"`
	DB     DatabaseConfig `koanf:"database"`
}

type DatabaseConfig struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	DBName   string `koanf:"db_name"`
	SSLMode  string `koanf:"ssl_mode"`
}

type ServerConfig struct {
	Port int `koanf:"port"`
}

type JWTConfig struct {
	SecretKey string `koanf:"secret_key"`
}

type LogConfig struct {
	Level string `koanf:"level"`
}
