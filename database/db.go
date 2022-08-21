package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host        string
	Port        string
	User        string
	Pass        string
	Name        string
	SSLMode     string
	SSLCert     string
	SSLKey      string
	SSLRootCert string
}

func (c *Config) url() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Pass, c.Host, c.Port, c.Name)
}

func Connect(cfg Config) (*pgxpool.Pool, error) {
	return pgxpool.Connect(context.Background(), cfg.url())
}
