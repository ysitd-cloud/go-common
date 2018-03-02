package db // import "code.ysitd.cloud/common/go/db"

import "github.com/tonyhhyip/go-di-container"

// NewPool create new db opener
func NewPool(driver, dsn string) DBOpener {
	return &opener{
		driver: driver,
		dsn: dsn,
	}
}

// NewServiceProvider create db opener service provider
func NewServiceProvider(app container.Container) container.ServiceProvider {
	sp := &serviceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}
	sp.SetContainer(app)
	return sp
}