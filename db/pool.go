package db // import "code.ysitd.cloud/common/go/db"

import "database/sql"

func (p *pool) Acquire() (*sql.DB, error) {
	return sql.Open(p.driver, p.dsn)
}
