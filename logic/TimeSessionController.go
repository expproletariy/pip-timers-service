package logic

import (
	dataV1 "github.com/expproletariy/pip-timers-service/data/version1"
	persist "github.com/expproletariy/pip-timers-service/persistence"
	ccmd "github.com/pip-services3-go/pip-services3-commons-go/commands"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

type TimeSessionController struct {
	persistence persist.ITimeSessionPersistence
	commandSet  *TimeSessionCommandSet
}

func NewTimeSessionController() *TimeSessionController {
	return &TimeSessionController{}
}

func (c *TimeSessionController) Configure(config *cconf.ConfigParams) {
	// Todo: Read configuration parameters here...
}

func (c *TimeSessionController) SetReferences(references cref.IReferences) {
	p, err := references.GetOneRequired(cref.NewDescriptor("pip-timers-service", "persistence", "*", "*", "1.0"))
	if p != nil && err == nil {
		if persistence, ok := p.(persist.ITimeSessionPersistence); ok {
			c.persistence = persistence
		}
	}
}

func (c *TimeSessionController) GetCommandSet() *ccmd.CommandSet {
	if c.commandSet == nil {
		c.commandSet = NewTimeSessionCommandSet(c)
	}
	return &c.commandSet.CommandSet
}

func (c *TimeSessionController) GetTimeSessions(
	correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (*dataV1.TimeSessionDataPage, error) {
	return c.persistence.GetPageByFilter(correlationId, filter, paging)
}

func (c *TimeSessionController) GetTimeSessionById(
	correlationId string, timeSessionId string) (*dataV1.TimeSession, error) {
	return c.persistence.GetOneById(correlationId, timeSessionId)
}

func (c *TimeSessionController) GetTimeSessionByUdi(
	correlationId string, timeSessionId string) (*dataV1.TimeSession, error) {
	//return c.persistence.GetOneByUdi(correlationId, timeSessionId)
	return nil, nil
}

func (c *TimeSessionController) CreateTimeSession(
	correlationId string, timeSession *dataV1.TimeSession) (*dataV1.TimeSession, error) {
	if timeSession.Id == "" {
		timeSession.Id = cdata.IdGenerator.NextLong()
	}

	timeSession.Status = dataV1.SessionStatusCreated

	return c.persistence.Create(correlationId, timeSession)
}

func (c *TimeSessionController) UpdateTimeSession(
	correlationId string, timeSession *dataV1.TimeSession) (*dataV1.TimeSession, error) {
	return c.persistence.Update(correlationId, timeSession)
}

func (c *TimeSessionController) DeleteTimeSessionById(
	correlationId string, timeSessionId string) (*dataV1.TimeSession, error) {
	return c.persistence.DeleteById(correlationId, timeSessionId)
}

func (c *TimeSessionController) AddTimerToTimeSession(
	correlationId string, timerSessionId string, timer dataV1.Timer) (item *dataV1.TimeSession, err error) {
	session, err := c.persistence.GetOneById(correlationId, timerSessionId)
	if err != nil {
		return session, err
	}

	if timer.Id == "" {
		timer.Id = cdata.IdGenerator.NextLong()
	}

	if session.Timers == nil {
		session.Timers = []dataV1.Timer{
			timer,
		}
	} else {
		session.Timers = append(session.Timers, timer)
	}

	return c.persistence.Update(correlationId, session)
}

func (c *TimeSessionController) UpdateTimerToTimeSession(
	correlationId string, timerSessionId string, timer dataV1.Timer) (item *dataV1.TimeSession, err error) {
	session, err := c.persistence.GetOneById(correlationId, timerSessionId)
	if err != nil {
		return session, err
	}

	if timer.Id == "" {
		return nil, cerr.NewError("Empty time session timers id ")
	}

	if session.Timers == nil {
		return nil, cerr.NewError("Empty time session timers ")
	}

	for i := range session.Timers {
		if session.Timers[i].Id == timer.Id {
			session.Timers[i] = timer
			break
		}
	}

	return c.persistence.Update(correlationId, session)
}

func (c *TimeSessionController) DeleteTimerFromTimeSession(correlationId, timerSessionId, timerId string) (item *dataV1.TimeSession, err error) {
	session, err := c.persistence.GetOneById(correlationId, timerSessionId)
	if err != nil {
		return session, err
	}

	var newTimers []dataV1.Timer
	if session.Timers != nil && len(session.Timers) != 0 {
		if len(session.Timers) == 1 && session.Timers[0].Id == timerId {
			newTimers = nil
		} else {
			for i := range session.Timers {
				if session.Timers[i].Id == timerId {
					newTimers = append(session.Timers[:i], session.Timers[i+1:]...)
					break
				}
			}
		}
		session.Timers = newTimers
	}

	return c.persistence.Update(correlationId, session)
}

func (c *TimeSessionController) SumTimersForTimeSession(correlationId, timerSessionId string) (*dataV1.TimersSum, error) {
	session, err := c.persistence.GetOneById(correlationId, timerSessionId)
	if err != nil {
		return &dataV1.TimersSum{
			Sum: 0,
		}, err
	}

	var sum float64
	for _, timer := range session.Timers {
		if timer.Status == dataV1.TimersStatusOff {
			sum += timer.StoppedAt.Sub(timer.StartedAt).Seconds()
		}
	}

	return &dataV1.TimersSum{
		Sum:           sum,
		TimeSessionId: session.Id,
	}, nil
}
