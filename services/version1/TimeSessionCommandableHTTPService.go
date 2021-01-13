package version1

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cservices "github.com/pip-services3-go/pip-services3-rpc-go/services"
)

type TimeSessionCommandableHttpService struct {
	*cservices.CommandableHttpService
}

func NewTimeSessionCommandableHttpService() *TimeSessionCommandableHttpService {
	c := &TimeSessionCommandableHttpService{
		CommandableHttpService: cservices.NewCommandableHttpService("v1/time_sessions"),
	}
	c.CommandableHttpService.IRegisterable = c
	c.DependencyResolver.Put("controller", cref.NewDescriptor("pip-timers-service", "controller", "*", "*", "1.0"))
	return c
}

func (c *TimeSessionCommandableHttpService) Register() {
	c.CommandableHttpService.Register()
}
