package server

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPServeMux struct {
		Patterns struct {
			GraphQL           string
			GraphQLPlayground string
		}
	}

	HTTPServer struct {
		Port int
	}

	GRPCServer struct {
		Port int
	}

	Logger struct {
		Level       string
		Development bool
	}
}

func (c *Config) LoadFile(configFile string) error {
	v := viper.New()
	v.SetConfigFile(configFile)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("load config file %s: %w", configFile, err)
	}
	if err := v.UnmarshalExact(c); err != nil {
		return fmt.Errorf("load config file %s: %w", configFile, err)
	}
	return nil
}
