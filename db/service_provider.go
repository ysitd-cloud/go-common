package db // import "code.ysitd.cloud/common/go/db"

import "github.com/tonyhhyip/go-di-container"

func (*serviceProvider) Provides() []string {
	return []string {
		"db.opener",
	}
}

func (*serviceProvider) Register(app container.Container) {
	app.Alias("db.opener", "db.opener")
	app.Singleton("db.opener", func(app container.Container) interface{} {
		return NewPool(
			app.Make("db.driver").(string),
			app.Make("db.url").(string),
			)
	})
}
