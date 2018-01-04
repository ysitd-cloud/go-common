package db

import "github.com/tonyhhyip/go-di-container"

func NewPool(driver, dsn string) Pool {
	return &pool{
		driver: driver,
		dsn: dsn,
	}
}

func NewServiceProvider(app container.Container) container.ServiceProvider {
	sp := &serviceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}
	sp.SetContainer(app)
	return sp
}