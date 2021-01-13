package version1

import (
	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
	cvalid "github.com/pip-services3-go/pip-services3-commons-go/validate"
)

type TimeSessionSchema struct {
	cvalid.ObjectSchema
}

func NewTimeSessionSchema() *TimeSessionSchema {
	c := TimeSessionSchema{}
	c.ObjectSchema = *cvalid.NewObjectSchema()

	c.WithOptionalProperty("id", cconv.String)
	c.WithRequiredProperty("name", cconv.String)
	c.WithRequiredProperty("user", cconv.String)
	c.WithOptionalProperty("status", cconv.String)
	c.WithOptionalProperty("tags", cconv.Array)
	return &c
}
