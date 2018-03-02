package db // import "code.ysitd.cloud/common/go/db"

import "database/sql"

func (p *opener) Acquire() (*sql.DB, error) {
	return sql.Open(p.driver, p.dsn)
}
