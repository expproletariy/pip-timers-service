package version1

import (
	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
	cvalid "github.com/pip-services3-go/pip-services3-commons-go/validate"
)

type TimerSchema struct {
	cvalid.ObjectSchema
}

func NewTimerSchema() *TimeSessionSchema {
	c := TimeSessionSchema{}
	c.ObjectSchema = *cvalid.NewObjectSchema()

	c.WithRequiredProperty("started_at", cconv.DateTime)
	c.WithRequiredProperty("stopped_at", cconv.DateTime)
	c.WithRequiredProperty("status", cconv.String)
	c.WithOptionalProperty("id", cconv.String)
	return &c
}
