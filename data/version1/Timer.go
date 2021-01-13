package version1

import "time"

type Timer struct {
	Id        string       `json:"id" bson:"_id"`
	StartedAt time.Time    `json:"started_at" bson:"started_at"`
	StoppedAt time.Time    `json:"stopped_at" bson:"stopped_at"`
	Status    TimersStatus `json:"status" bson:"status"`
}
