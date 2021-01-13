package persistence

import (
	dataV1 "github.com/expproletariy/pip-timers-service/data/version1"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
)

type ITimeSessionPersistence interface {
	GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *dataV1.TimeSessionDataPage, err error)

	GetOneById(correlationId string, id string) (res *dataV1.TimeSession, err error)

	Create(correlationId string, item *dataV1.TimeSession) (res *dataV1.TimeSession, err error)

	Update(correlationId string, item *dataV1.TimeSession) (res *dataV1.TimeSession, err error)

	DeleteById(correlationId string, id string) (res *dataV1.TimeSession, err error)
}
