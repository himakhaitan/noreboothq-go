package config

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
var k = koanf.New(".")

// LoadConfig loads a configuration file from the specified path and unmarshals it into the provided type T.
// It returns a pointer to T and an error if any occurs during loading or unmarshalling.
// T can be any struct type that matches the structure of the configuration file.
func LoadConfig[T any](basePath string, env string) (*T, error) {
	if basePath == "" {
		return nil, fmt.Errorf("basePath cannot be empty")
	}

	// Loading the base configuration file.
	if err := k.Load(file.Provider(fmt.Sprintf("%s/base.yaml", basePath)), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("failed to load config file from %s: %w", basePath, err)
	}

	// Loading the environment-specific configuration file.
	configPath := fmt.Sprintf("%s/%s.yaml", basePath, env)

	if configPath != "" {
		if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
			return nil, fmt.Errorf("failed to load environment config file from %s: %w", configPath, err)
		}
	}

	var cfg T
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}