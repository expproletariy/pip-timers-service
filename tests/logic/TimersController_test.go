package test_logic

import (
	dataV1 "github.com/expproletariy/pip-timers-service/data/version1"
	"github.com/expproletariy/pip-timers-service/logic"
	"github.com/expproletariy/pip-timers-service/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type timersControllerTest struct {
	TIMERS1     *dataV1.TimeSession
	TIMERS2     *dataV1.TimeSession
	persistence *persistence.TimersMemoryPersistence
	controller  *logic.TimeSessionController
}

func newTimersControllerTest() *timersControllerTest {
	TIMERS1 := &dataV1.TimeSession{
		Id:        "timer1",
		Name:      "timer1",
		User:      "user1",
		Tags:      nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    dataV1.SessionStatusCreated,
		Timers:    nil,
	}

	TIMERS2 := &dataV1.TimeSession{
		Id:        "timer2",
		Name:      "timer2",
		User:      "user1",
		Tags:      nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    dataV1.SessionStatusCreated,
		Timers:    nil,
	}

	persist := persistence.NewTimersMemoryPersistence()
	persist.Configure(cconf.NewEmptyConfigParams())

	controller := logic.NewTimeSessionController()
	controller.Configure(cconf.NewEmptyConfigParams())
	controller.SetReferences(cref.NewReferencesFromTuples(
		cref.NewDescriptor("pip-timers-service", "persistence", "memory", "default", "1.0"), persist,
	))

	return &timersControllerTest{
		TIMERS1:     TIMERS1,
		TIMERS2:     TIMERS2,
		persistence: persist,
		controller:  controller,
	}
}

func (c *timersControllerTest) setup(t *testing.T) {
	err := c.persistence.Open("")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear("")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *timersControllerTest) teardown(t *testing.T) {
	err := c.persistence.Close("")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func (c timersControllerTest) testCrudOperations(t *testing.T) {

	// Test Timers session logic
	timerSession, err := c.controller.CreateTimeSession("", c.TIMERS1)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS1.Id)

	timerSession, err = c.controller.CreateTimeSession("", c.TIMERS2)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS2.Id)

	page, err := c.controller.GetTimeSessions("", cdata.NewEmptyFilterParams(), cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)

	timerSession, err = c.controller.GetTimeSessionById("", c.TIMERS1.Id)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS1.Id)

	const status = dataV1.SessionStatusInUse
	timerSession.Status = status
	timerSession, err = c.controller.UpdateTimeSession("", timerSession)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Status, status)

	timerSession, err = c.controller.DeleteTimeSessionById("", c.TIMERS1.Id)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS1.Id)

	// Test Timers logic
	timerSession, err = c.controller.CreateTimeSession("", c.TIMERS1)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS1.Id)

	startDate, _ := time.Parse(time.RFC3339, "2020-01-12T15:01:01Z")
	stoppDate := startDate.Add(time.Second * 2)
	timerId1 := "timer1"
	timerId2 := "timer2"
	timer1 := dataV1.Timer{
		Id:        timerId1,
		StartedAt: startDate,
		StoppedAt: stoppDate,
		Status:    dataV1.TimersStatusOff,
	}
	timer2 := dataV1.Timer{
		Id:        timerId2,
		StartedAt: startDate,
		StoppedAt: stoppDate,
		Status:    dataV1.TimersStatusOff,
	}

	timerSession, err = c.controller.AddTimerToTimeSession("", c.TIMERS1.Id, timer1)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS1.Id)
	assert.NotNil(t, timerSession.Timers)
	assert.Len(t, timerSession.Timers, 1)
	assert.Equal(t, timerSession.Timers[0].Id, "timer1")
	assert.Equal(t, timerSession.Timers[0].StartedAt, startDate)
	assert.Equal(t, timerSession.Timers[0].StoppedAt, stoppDate)

	timerSession, err = c.controller.AddTimerToTimeSession("", c.TIMERS1.Id, timer2)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS1.Id)
	assert.NotNil(t, timerSession.Timers)
	assert.Len(t, timerSession.Timers, 2)

	timerSessionSum, err := c.controller.SumTimersForTimeSession("", c.TIMERS1.Id)
	assert.Nil(t, err)
	assert.NotNil(t, timerSessionSum)
	assert.NotZero(t, timerSessionSum.Sum)
	assert.Equal(t, timerSessionSum.Sum, time.Second.Seconds()*4)

	// Test update timer
	timerSession, err = c.controller.UpdateTimerToTimeSession("", c.TIMERS1.Id, dataV1.Timer{
		Id:        timerId1,
		StartedAt: startDate,
		StoppedAt: time.Time{},
		Status:    dataV1.TimersStatusOn,
	})
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS1.Id)
	assert.NotNil(t, timerSession.Timers)
	assert.Len(t, timerSession.Timers, 2)

	timerSessionSum, err = c.controller.SumTimersForTimeSession("", c.TIMERS1.Id)
	assert.Nil(t, err)
	assert.NotNil(t, timerSessionSum)
	assert.NotZero(t, timerSessionSum.Sum)
	assert.Equal(t, timerSessionSum.Sum, time.Second.Seconds()*2)

	timerSession, err = c.controller.DeleteTimerFromTimeSession("", c.TIMERS1.Id, timerId1)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS1.Id)
	assert.NotNil(t, timerSession.Timers)
	assert.Len(t, timerSession.Timers, 1)

	timerSession, err = c.controller.DeleteTimerFromTimeSession("", c.TIMERS1.Id, timerId2)
	assert.Nil(t, err)
	assert.NotNil(t, timerSession)
	assert.Equal(t, timerSession.Id, c.TIMERS1.Id)
	assert.Nil(t, timerSession.Timers)
}

func TestTimersController(t *testing.T) {
	c := newTimersControllerTest()

	c.setup(t)
	t.Run("CRUD Operations", c.testCrudOperations)
	c.teardown(t)
}
