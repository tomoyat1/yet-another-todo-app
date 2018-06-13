package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type PgConfig struct {
	User string `envconfig:"db_user"`
	Password string `envconfig:"db_passwd"`
	Host string `envconfig:"db_host"`
	Port uint16 `envconfig:"db_port"`
	Name string `envconfig:"db_name"`
}

func DBConfigFromEnv() (*PgConfig, error) {
	var c PgConfig
	err := envconfig.Process("todo", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *PgConfig) GenerateConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}
