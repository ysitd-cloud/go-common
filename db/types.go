package db // import "code.ysitd.cloud/common/go/db"

import (
	"database/sql"

	"github.com/tonyhhyip/go-di-container"
)

// Pool provide way to acquire db connection
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
