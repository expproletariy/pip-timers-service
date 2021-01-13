package logic

import (
	"encoding/json"
	dataV1 "github.com/expproletariy/pip-timers-service/data/version1"
	ccmd "github.com/pip-services3-go/pip-services3-commons-go/commands"
	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	crun "github.com/pip-services3-go/pip-services3-commons-go/run"
	cvalid "github.com/pip-services3-go/pip-services3-commons-go/validate"
)

type TimeSessionCommandSet struct {
	ccmd.CommandSet
	controller IBeaconsController
}

func NewTimeSessionCommandSet(controller IBeaconsController) *TimeSessionCommandSet {
	c := &TimeSessionCommandSet{
		CommandSet: *ccmd.NewCommandSet(),
		controller: controller,
	}
	c.AddCommand(c.makeGetTimeSessionsCommand())
	c.AddCommand(c.makeGetTimeSessionByIdCommand())
	c.AddCommand(c.makeCreateTimeSessionCommand())
	c.AddCommand(c.makeUpdateTimeSessionCommand())
	c.AddCommand(c.makeDeleteTimeSessionByIdCommand())
	c.AddCommand(c.makeAddTimerToTimeSessionCommand())
	c.AddCommand(c.makeUpdateTimerToTimeSessionCommand())
	c.AddCommand(c.makeDeleteTimerFromTimeSessionCommand())
	c.AddCommand(c.makeSumTimersForTimeSessionCommand())

	return c
}

func (c *TimeSessionCommandSet) makeGetTimeSessionsCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_time_sessions",
		cvalid.NewObjectSchema().
			WithOptionalProperty("filter", cvalid.NewFilterParamsSchema()).
			WithOptionalProperty("paging", cvalid.NewPagingParamsSchema()),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			filter := cdata.NewFilterParamsFromValue(args.Get("filter"))
			paging := cdata.NewPagingParamsFromValue(args.Get("paging"))
			return c.controller.GetTimeSessions(correlationId, filter, paging)
		})
}

func (c *TimeSessionCommandSet) makeGetTimeSessionByIdCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_time_session_by_id",
		cvalid.NewObjectSchema().
			WithRequiredProperty("time_session_id", cconv.String),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			timeSessionId := args.GetAsString("time_session_id")
			return c.controller.GetTimeSessionById(correlationId, timeSessionId)
		})
}

func (c *TimeSessionCommandSet) makeCreateTimeSessionCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"create_time_session",
		cvalid.NewObjectSchema().
			WithRequiredProperty("time_session", dataV1.NewTimeSessionSchema()),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {

			val, errJ := json.Marshal(args.Get("time_session"))
			var timeSession dataV1.TimeSession
			errJ = json.Unmarshal(val, &timeSession)
			if errJ != nil {
				return nil, errJ
			}
			return c.controller.CreateTimeSession(correlationId, &timeSession)
		})
}

func (c *TimeSessionCommandSet) makeUpdateTimeSessionCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"update_time_session",
		cvalid.NewObjectSchema().
			WithRequiredProperty("time_session", dataV1.NewTimeSessionSchema()),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			val, _ := json.Marshal(args.Get("time_session"))
			var timeSession dataV1.TimeSession
			_ = json.Unmarshal(val, &timeSession)
			return c.controller.UpdateTimeSession(correlationId, &timeSession)
		})
}

func (c *TimeSessionCommandSet) makeDeleteTimeSessionByIdCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"delete_time_session_by_id",
		cvalid.NewObjectSchema().
			WithRequiredProperty("time_session_id", cconv.String),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			timeSessionId := args.GetAsString("time_session_id")
			return c.controller.DeleteTimeSessionById(correlationId, timeSessionId)
		})
}

func (c *TimeSessionCommandSet) makeAddTimerToTimeSessionCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"add_timer_to_time_session",
		cvalid.NewObjectSchema().
			WithRequiredProperty("time_session_id", cconv.String).
			WithRequiredProperty("timer", dataV1.NewTimerSchema()),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			timeSessionId := args.GetAsString("time_session_id")
			var timer dataV1.Timer
			bufferMap, err := json.Marshal(args.GetAsObject("timer"))
			if err != nil {
				return timer, cerr.NewError("can not marshal timer object")
			}
			err = json.Unmarshal(bufferMap, &timer)
			if err != nil {
				return timer, cerr.NewError("can not unmarshal timer object")
			}
			return c.controller.AddTimerToTimeSession(correlationId, timeSessionId, timer)
		})
}

func (c *TimeSessionCommandSet) makeUpdateTimerToTimeSessionCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"update_timer_to_time_session",
		cvalid.NewObjectSchema().
			WithRequiredProperty("time_session_id", cconv.String).
			WithRequiredProperty("timer", dataV1.NewTimerSchema()),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			timeSessionId := args.GetAsString("time_session_id")
			var timer dataV1.Timer
			bufferMap, err := json.Marshal(args.GetAsObject("timer"))
			if err != nil {
				return timer, cerr.NewError("can not marshal timer object")
			}
			err = json.Unmarshal(bufferMap, &timer)
			if err != nil {
				return timer, cerr.NewError("can not unmarshal timer object")
			}
			return c.controller.UpdateTimerToTimeSession(correlationId, timeSessionId, timer)
		})
}

func (c *TimeSessionCommandSet) makeDeleteTimerFromTimeSessionCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"delete_timer_from_time_session",
		cvalid.NewObjectSchema().
			WithRequiredProperty("time_session_id", cconv.String).
			WithRequiredProperty("timer_id", cconv.String),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			timeSessionId := args.GetAsString("time_session_id")
			timerId := args.GetAsString("timer_id")
			return c.controller.DeleteTimerFromTimeSession(correlationId, timeSessionId, timerId)
		})
}

func (c *TimeSessionCommandSet) makeSumTimersForTimeSessionCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"sum_timers_for_time_session",
		cvalid.NewObjectSchema().
			WithRequiredProperty("time_session_id", cconv.String),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			timeSessionId := args.GetAsString("time_session_id")
			return c.controller.SumTimersForTimeSession(correlationId, timeSessionId)
		})
}
