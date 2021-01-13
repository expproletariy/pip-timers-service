package persistence

import (
	dataV1 "github.com/expproletariy/pip-timers-service/data/version1"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cpersist "github.com/pip-services3-go/pip-services3-data-go/persistence"
	"reflect"
)

type TimersMemoryPersistence struct {
	cpersist.IdentifiableMemoryPersistence
}

func NewTimersMemoryPersistence() *TimersMemoryPersistence {
	proto := reflect.TypeOf(&dataV1.TimeSession{})
	c := TimersMemoryPersistence{
		IdentifiableMemoryPersistence: *cpersist.NewIdentifiableMemoryPersistence(proto),
	}
	c.MaxPageSize = 1000
	return &c
}

func (c *TimersMemoryPersistence) composeFilter(filter *cdata.FilterParams) func(beacon interface{}) bool {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	id := filter.GetAsString("id")
	userId := filter.GetAsString("user")
	createdAt := filter.GetAsDateTime("created_at")
	updatedAt := filter.GetAsDateTime("updated_at")
	status, isValidStatus := dataV1.SessionStatusFromString(filter.GetAsString("status"))

	return func(timerSession interface{}) bool {
		if item, ok := timerSession.(dataV1.TimeSession); ok {
			if id != "" && item.Id != id {
				return false
			}
			if userId != "" && item.User != userId {
				return false
			}

			if !createdAt.IsZero() && !item.CreatedAt.Equal(createdAt) {
				return false
			}
			if !updatedAt.IsZero() && !item.UpdatedAt.Equal(updatedAt) {
				return false
			}
			if isValidStatus && item.Status != status {
				return false
			}
		} else {
			return false
		}

		return true
	}
}

func (c *TimersMemoryPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (*dataV1.TimeSessionDataPage, error) {
	tempPage, err := c.IdentifiableMemoryPersistence.GetPageByFilter(correlationId, c.composeFilter(filter), paging, nil, nil)
	if tempPage == nil || err != nil {
		return nil, err
	}

	dataLen := len(tempPage.Data)
	data := make([]*dataV1.TimeSession, dataLen)
	for i, v := range tempPage.Data {
		data[i] = v.(*dataV1.TimeSession)
	}
	page := dataV1.NewTimeSessionDataPage(tempPage.Total, data)

	return page, nil
}

func (c *TimersMemoryPersistence) GetOneById(correlationId string, id string) (*dataV1.TimeSession, error) {
	result, err := c.IdentifiableMemoryPersistence.GetOneById(correlationId, id)

	if result == nil || err != nil {
		return nil, err
	}

	// Convert to BeaconV1
	item, _ := result.(*dataV1.TimeSession)
	return item, err
}

func (c *TimersMemoryPersistence) Create(correlationId string, item *dataV1.TimeSession) (*dataV1.TimeSession, error) {
	value, err := c.IdentifiableMemoryPersistence.Create(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	result, _ := value.(*dataV1.TimeSession)
	return result, nil
}

func (c *TimersMemoryPersistence) Update(correlationId string, item *dataV1.TimeSession) (*dataV1.TimeSession, error) {
	value, err := c.IdentifiableMemoryPersistence.Update(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	// Convert to BeaconV1
	result, _ := value.(*dataV1.TimeSession)
	return result, nil
}

func (c *TimersMemoryPersistence) DeleteById(correlationId string, id string) (*dataV1.TimeSession, error) {
	result, err := c.IdentifiableMemoryPersistence.DeleteById(correlationId, id)

	if result == nil || err != nil {
		return nil, err
	}

	// Convert to BeaconV1
	item, _ := result.(*dataV1.TimeSession)
	return item, nil
}
