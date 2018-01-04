package db

import "database/sql"

func (p *pool) Acquire() (*sql.DB, error) {
	return sql.Open(p.driver, p.dsn)
}
