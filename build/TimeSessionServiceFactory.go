package build

import (
	logic "github.com/expproletariy/pip-timers-service/logic"
	persist "github.com/expproletariy/pip-timers-service/persistence"
	serviceV1 "github.com/expproletariy/pip-timers-service/services/version1"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type TimeSessionServiceFactory struct {
	cbuild.Factory
}

func NewTimeSessionServiceFactory() *TimeSessionServiceFactory {
	c := &TimeSessionServiceFactory{
		Factory: *cbuild.NewFactory(),
	}

	c.RegisterType(
		cref.NewDescriptor("pip-timers-service", "persistence", "memory", "*", "1.0"),
		persist.NewTimersMemoryPersistence,
	)

	c.RegisterType(
		cref.NewDescriptor("pip-timers-service", "controller", "default", "*", "1.0"),
		logic.NewTimeSessionController,
	)

	c.RegisterType(
		cref.NewDescriptor("pip-timers-service", "service", "commandable-http", "*", "1.0"),
		serviceV1.NewTimeSessionCommandableHttpService,
	)

	return c
}
