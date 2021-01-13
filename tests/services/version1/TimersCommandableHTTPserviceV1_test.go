package test_services1

import (
	"bytes"
	"encoding/json"
	dataV1 "github.com/expproletariy/pip-timers-service/data/version1"
	"github.com/expproletariy/pip-timers-service/logic"
	"github.com/expproletariy/pip-timers-service/persistence"
	serviceV1 "github.com/expproletariy/pip-timers-service/services/version1"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type timersCommandableHttpServiceV1Test struct {
	TIMERS1     *dataV1.TimeSession
	TIMERS2     *dataV1.TimeSession
	persistence *persistence.TimersMemoryPersistence
	controller  *logic.TimeSessionController
	service     *serviceV1.TimeSessionCommandableHttpService
}

func newTimersCommandableHttpServiceV1Test() *timersCommandableHttpServiceV1Test {
	TIMERS1 := &dataV1.TimeSession{
		Id:        "timer1",
		Name:      "timer1",
		User:      "user1",
		Tags:      []string{"first", "test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    dataV1.SessionStatusCreated,
		Timers:    nil,
	}

	TIMERS2 := &dataV1.TimeSession{
		Id:        "timer2",
		Name:      "timer2",
		User:      "user1",
		Tags:      []string{"second", "test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    dataV1.SessionStatusCreated,
		Timers:    nil,
	}

	persist := persistence.NewTimersMemoryPersistence()
	persist.Configure(cconf.NewEmptyConfigParams())

	controller := logic.NewTimeSessionController()
	controller.Configure(cconf.NewEmptyConfigParams())

	service := serviceV1.NewTimeSessionCommandableHttpService()
	service.Configure(cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3005",
		"connection.host", "localhost",
	))

	references := refer.NewReferencesFromTuples(
		refer.NewDescriptor("pip-timers-service", "persistence", "memory", "default", "1.0"), persist,
		refer.NewDescriptor("pip-timers-service", "controller", "default", "default", "1.0"), controller,
		refer.NewDescriptor("pip-timers-service", "service", "http", "default", "1.0"), service,
	)
	controller.SetReferences(references)
	service.SetReferences(references)
	return &timersCommandableHttpServiceV1Test{
		TIMERS1:     TIMERS1,
		TIMERS2:     TIMERS2,
		persistence: persist,
		controller:  controller,
		service:     service,
	}
}

func (c *timersCommandableHttpServiceV1Test) setup(t *testing.T) {
	err := c.persistence.Open("")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.service.Open("")
	if err != nil {
		t.Error("Failed to open service", err)
	}

	err = c.persistence.Clear("")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *timersCommandableHttpServiceV1Test) teardown(t *testing.T) {
	err := c.service.Close("")
	if err != nil {
		t.Error("Failed to close service", err)
	}

	err = c.persistence.Close("")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func (c *timersCommandableHttpServiceV1Test) invoke(
	route string, body *cdata.AnyValueMap, result interface{}) error {
	var url string = "http://localhost:3005" + route

	var bodyReader *bytes.Reader = nil
	if body != nil {
		jsonBody, _ := json.Marshal(body.Value())
		bodyReader = bytes.NewReader(jsonBody)
	}

	postResponse, postErr := http.Post(url, "application/json", bodyReader)

	if postErr != nil {
		return postErr
	}

	if postResponse.StatusCode == 204 {
		return nil
	}

	resBody, bodyErr := ioutil.ReadAll(postResponse.Body)
	if bodyErr != nil {
		return bodyErr
	}

	if postResponse.StatusCode >= 400 {
		appErr := cerr.ApplicationError{}
		json.Unmarshal(resBody, &appErr)
		return &appErr
	}

	if result == nil {
		return nil
	}

	jsonErr := json.Unmarshal(resBody, result)
	return jsonErr
}

func (c *timersCommandableHttpServiceV1Test) testCrudOperations(t *testing.T) {
	var timeSession dataV1.TimeSession
	var page dataV1.TimeSessionDataPage
	// Create 1st time session
	body := cdata.NewAnyValueMapFromTuples(
		"time_session", c.TIMERS1,
	)
	err := c.invoke("/v1/time_sessions/create_time_session", body, &timeSession)
	assert.Nil(t, err)
	assert.Equal(t, timeSession.Id, c.TIMERS1.Id)
	assert.Equal(t, timeSession.Status, c.TIMERS1.Status)
	assert.Len(t, timeSession.Tags, 2)

	// Create 2nd time session
	body = cdata.NewAnyValueMapFromTuples(
		"time_session", c.TIMERS2,
	)
	err = c.invoke("/v1/time_sessions/create_time_session", body, &timeSession)
	assert.Nil(t, err)
	assert.Equal(t, timeSession.Id, c.TIMERS2.Id)
	assert.Equal(t, timeSession.Status, c.TIMERS2.Status)
	assert.Len(t, timeSession.Tags, 2)

	// Get time session
	body = cdata.NewAnyValueMapFromTuples(
		"time_session_id", c.TIMERS1.Id,
	)
	err = c.invoke("/v1/time_sessions/get_time_session_by_id", body, &timeSession)
	assert.Nil(t, err)
	assert.Equal(t, timeSession.Id, c.TIMERS1.Id)
	assert.Equal(t, timeSession.Status, c.TIMERS1.Status)
	assert.Len(t, timeSession.Tags, 2)

	// Get time sessions by filters
	body = cdata.NewAnyValueMapFromTuples(
		"filter", cdata.NewEmptyFilterParams(),
		"paging", cdata.NewEmptyFilterParams(),
	)
	err = c.invoke("/v1/time_sessions/get_time_sessions", body, &page)
	assert.Nil(t, err)
	assert.NotNil(t, page.Data)
	assert.Len(t, page.Data, 2)

	// Add timer to time session
	startDate, _ := time.Parse(time.RFC3339, "2020-01-12T15:01:01Z")
	stoppDate := startDate.Add(time.Second * 2)
	timer := dataV1.Timer{
		StartedAt: startDate,
		StoppedAt: stoppDate,
		Status:    dataV1.TimersStatusOff,
	}
	body = cdata.NewAnyValueMapFromTuples(
		"time_session_id", c.TIMERS1.Id,
		"timer", timer,
	)
	err = c.invoke("/v1/time_sessions/add_timer_to_time_session", body, &timeSession)
	assert.Nil(t, err)
	assert.NotNil(t, timeSession)
	assert.Equal(t, timeSession.Id, c.TIMERS1.Id)
	assert.NotNil(t, timeSession.Timers)
	assert.Len(t, timeSession.Timers, 1)

	// Sum timers for time session
	body = cdata.NewAnyValueMapFromTuples(
		"time_session_id", c.TIMERS1.Id,
	)
	var timersSum dataV1.TimersSum
	err = c.invoke("/v1/time_sessions/sum_timers_for_time_session", body, &timersSum)
	assert.Nil(t, err)
	assert.NotNil(t, timersSum)
	assert.NotZero(t, timersSum.Sum)
	assert.Equal(t, timersSum.Sum, time.Second.Seconds()*2)

	// Update timer to time session
	body = cdata.NewAnyValueMapFromTuples(
		"time_session_id", c.TIMERS1.Id,
		"timer", dataV1.Timer{
			Id:        timeSession.Timers[0].Id,
			StartedAt: startDate,
			StoppedAt: startDate.Add(time.Minute),
			Status:    dataV1.TimersStatusOn,
		},
	)
	err = c.invoke("/v1/time_sessions/update_timer_to_time_session", body, &timeSession)
	assert.Nil(t, err)
	assert.NotNil(t, timeSession)
	assert.Equal(t, timeSession.Id, c.TIMERS1.Id)
	assert.NotNil(t, timeSession.Timers)
	assert.Len(t, timeSession.Timers, 1)
	assert.Equal(t, timeSession.Timers[0].Status, dataV1.TimersStatusOn)

	body = cdata.NewAnyValueMapFromTuples(
		"time_session_id", c.TIMERS1.Id,
	)
	err = c.invoke("/v1/time_sessions/sum_timers_for_time_session", body, &timersSum)
	assert.Nil(t, err)
	assert.NotNil(t, timersSum)
	assert.Zero(t, timersSum.Sum)

	// Delete timer from time session
	body = cdata.NewAnyValueMapFromTuples(
		"time_session_id", c.TIMERS1.Id,
		"timer_id", timeSession.Timers[0].Id,
	)
	err = c.invoke("/v1/time_sessions/delete_timer_from_time_session", body, &timeSession)
	assert.Nil(t, err)
	assert.NotNil(t, timeSession)
	assert.Nil(t, timeSession.Timers)

	body = cdata.NewAnyValueMapFromTuples(
		"time_session_id", c.TIMERS1.Id,
	)
	err = c.invoke("/v1/time_sessions/sum_timers_for_time_session", body, &timersSum)
	assert.Nil(t, err)
	assert.NotNil(t, timersSum)
	assert.Zero(t, timersSum.Sum)
}

func TestTimersCommmandableHttpServiceV1(t *testing.T) {
	c := newTimersCommandableHttpServiceV1Test()

	c.setup(t)
	t.Run("CRUD Operations", c.testCrudOperations)
	c.teardown(t)

}
