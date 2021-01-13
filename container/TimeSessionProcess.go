package container

import (
	build "github.com/expproletariy/pip-timers-service/build"
	cproc "github.com/pip-services3-go/pip-services3-container-go/container"
	rbuild "github.com/pip-services3-go/pip-services3-rpc-go/build"
)

type TimeSessionProcess struct {
	cproc.ProcessContainer
}

func NewTimeSessionProcess() *TimeSessionProcess {
	c := &TimeSessionProcess{
		ProcessContainer: *cproc.NewProcessContainer("pip-timers-service", "One more timers microservice"),
	}
	c.AddFactory(build.NewTimeSessionServiceFactory())
	c.AddFactory(rbuild.NewDefaultRpcFactory())
	return c
}
