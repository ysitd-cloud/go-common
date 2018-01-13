package db // import "code.ysitd.cloud/common/go/db"

import "github.com/tonyhhyip/go-di-container"

func (*serviceProvider) Provides() []string {
	return []string {
		"db.pool",
	}
}

func (*serviceProvider) Register(app container.Container) {
	app.Singleton("db.pool", func(app container.Container) interface{} {
		return NewPool(
			app.Make("db.driver").(string),
			app.Make("db.url").(string),
			)
	})
}
