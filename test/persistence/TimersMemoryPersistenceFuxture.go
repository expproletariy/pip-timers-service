package test_persistence

import (
	dataV1 "github.com/expproletariy/pip-timers-service/data/version1"
	persist "github.com/expproletariy/pip-timers-service/persistence"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type TimersPersistenceFixture struct {
	TIMERS1     *dataV1.TimeSession
	TIMERS2     *dataV1.TimeSession
	TIMERS3     *dataV1.TimeSession
	persistence persist.ITimeSessionPersistence
}

func NewTimersPersistenceFixture(persistence persist.ITimeSessionPersistence) *TimersPersistenceFixture {
	t := TimersPersistenceFixture{}

	t.TIMERS1 = &dataV1.TimeSession{
		Id:        "timer1",
		Name:      "timer1",
		User:      "user1",
		Tags:      nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    dataV1.SessionStatusCreated,
		Timers:    nil,
	}

	t.TIMERS2 = &dataV1.TimeSession{
		Id:        "timer2",
		Name:      "timer2",
		User:      "user1",
		Tags:      nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    dataV1.SessionStatusCreated,
		Timers:    nil,
	}

	t.TIMERS3 = &dataV1.TimeSession{
		Id:        "timer3",
		Name:      "timer3",
		User:      "user1",
		Tags:      nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    dataV1.SessionStatusCreated,
		Timers:    nil,
	}

	t.persistence = persistence
	return &t
}

func (c *TimersPersistenceFixture) testCreateTimers(t *testing.T) {
	timerRes, err := c.persistence.Create("", c.TIMERS1)
	assert.Nil(t, err)
	assert.Equal(t, timerRes.Id, c.TIMERS1.Id)
	assert.Equal(t, timerRes.Name, c.TIMERS1.Name)
	assert.Equal(t, timerRes.User, c.TIMERS1.User)
	assert.Equal(t, timerRes.Status, c.TIMERS1.Status)

	timerRes, err = c.persistence.Create("", c.TIMERS2)
	assert.Nil(t, err)
	assert.Equal(t, timerRes.Id, c.TIMERS2.Id)
	assert.Equal(t, timerRes.Name, c.TIMERS2.Name)
	assert.Equal(t, timerRes.User, c.TIMERS2.User)
	assert.Equal(t, timerRes.Status, c.TIMERS2.Status)

	timerRes, err = c.persistence.Create("", c.TIMERS3)
	assert.Nil(t, err)
	assert.Equal(t, timerRes.Id, c.TIMERS3.Id)
	assert.Equal(t, timerRes.Name, c.TIMERS3.Name)
	assert.Equal(t, timerRes.User, c.TIMERS3.User)
	assert.Equal(t, timerRes.Status, c.TIMERS3.Status)
}

func (c *TimersPersistenceFixture) TestCrudOperations(t *testing.T) {

	c.testCreateTimers(t)

	page, err := c.persistence.GetPageByFilter("", cdata.NewEmptyFilterParams(), cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 3)

	timerSession, err := c.persistence.GetOneById("", c.TIMERS1.Id)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS1.Id)
	assert.Equal(t, timerSession.Name, c.TIMERS1.Name)
	assert.Equal(t, timerSession.User, c.TIMERS1.User)

	timerSession.Status = dataV1.SessionStatusInUse
	timerSessionUpdated, err := c.persistence.Update("", timerSession)
	assert.Nil(t, err)
	assert.NotNil(t, timerSessionUpdated)
	assert.Equal(t, timerSessionUpdated.Status, timerSession.Status)

	timerSessionDeleted, err := c.persistence.DeleteById("", timerSession.Id)
	assert.Nil(t, err)
	assert.NotNil(t, timerSessionDeleted)
	assert.Equal(t, timerSessionDeleted.Id, timerSession.Id)

	timerSessionDeleted, err = c.persistence.GetOneById("", timerSession.Id)
	assert.Nil(t, err)
	assert.Nil(t, timerSessionDeleted)
}

func (c *TimersPersistenceFixture) TestGetWithFilters(t *testing.T) {
	c.testCreateTimers(t)

	const status = dataV1.SessionStatusCreated
	page, err := c.persistence.GetPageByFilter(
		"",
		cdata.NewFilterParamsFromTuples(
			"status", status,
		),
		cdata.NewEmptyPagingParams(),
	)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 3)
	assert.Equal(t, page.Data[0].Status, status)

	const user = "user1"
	page, err = c.persistence.GetPageByFilter(
		"",
		cdata.NewFilterParamsFromTuples(
			"user", user,
		),
		cdata.NewEmptyPagingParams(),
	)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 3)
	assert.Equal(t, page.Data[0].User, user)

	const id = "timer1"
	page, err = c.persistence.GetPageByFilter(
		"",
		cdata.NewFilterParamsFromTuples(
			"id", id,
		),
		cdata.NewEmptyPagingParams(),
	)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 1)
	assert.Equal(t, page.Data[0].Id, id)

}
