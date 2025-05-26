package env

import (
	"flag"
	"os"
)

type EnvConfig struct {
	Env        string
	ConfigPath string
}

// ResolveConfig processes CLI flags and env vars in order of precedence:
// CLI flags > environment variables > defaults.
func ResolveEnvConfig(defaultConfigPath string, defaultEnv string) *EnvConfig {
	envFlag := flag.String("env", "", "Environment to run the service in (e.g., development, production)")
	configPathFlag := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	configPath := *configPathFlag
	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	if configPath == "" {
		configPath = defaultConfigPath
	}

	env := *envFlag
	if env == "" {
		env = os.Getenv("ENV")
	}
	if env == "" {
		env = defaultEnv
	}

	return &EnvConfig{
		Env:        env,
		ConfigPath: configPath,
	}
}
