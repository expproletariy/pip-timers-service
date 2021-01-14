package logic

import (
	dataV1 "github.com/expproletariy/pip-timers-service/data/version1"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
)

type ITimeSessionController interface {
	GetTimeSessions(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *dataV1.TimeSessionDataPage, err error)

	GetTimeSessionById(correlationId string, timerSessionId string) (item *dataV1.TimeSession, err error)

	GetTimeSessionByUdi(correlationId string, timerSessionId string) (item *dataV1.TimeSession, err error)

	CreateTimeSession(correlationId string, beacon *dataV1.TimeSession) (item *dataV1.TimeSession, err error)

	UpdateTimeSession(correlationId string, beacon *dataV1.TimeSession) (item *dataV1.TimeSession, err error)

	DeleteTimeSessionById(correlationId string, timerSessionId string) (item *dataV1.TimeSession, err error)

	AddTimerToTimeSession(correlationId string, timerSessionId string, timer dataV1.Timer) (item *dataV1.TimeSession, err error)

	UpdateTimerToTimeSession(correlationId string, timerSessionId string, timer dataV1.Timer) (item *dataV1.TimeSession, err error)

	DeleteTimerFromTimeSession(correlationId, timerSessionId, timerId string) (item *dataV1.TimeSession, err error)

	SumTimersForTimeSession(correlationId, timerSessionId string) (timersSum *dataV1.TimersSum, err error)
}
