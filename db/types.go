package db

import (
	"database/sql"

	"github.com/tonyhhyip/go-di-container"
)

type Pool interface {
	Acquire() (*sql.DB, error)
}

type pool struct {
	driver string
	dsn    string
}

type serviceProvider struct {
	*container.AbstractServiceProvider
}
