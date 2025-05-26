package config

type AuthServiceConfig struct {
	Server ServerConfig `koanf:"server"`
	JWT    JWTConfig    `koanf:"jwt"`
	Log    LogConfig    `koanf:"logging"`
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
