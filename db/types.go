package db // import "code.ysitd.cloud/common/go/db"

import (
	"database/sql"

	"github.com/tonyhhyip/go-di-container"
)

// Alias for backward compatible. Will remove soon
type Pool = DBOpener

// Pool provide way to acquire db connection
type DBOpener interface {
	Acquire() (*sql.DB, error)
}

type opener struct {
	driver string
	dsn    string
}

type serviceProvider struct {
	*container.AbstractServiceProvider
}
