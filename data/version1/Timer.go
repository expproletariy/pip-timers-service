package version1

import "time"

type Timer struct {
	Id        string       `json:"id"`
	StartedAt time.Time    `json:"started_at"`
	StoppedAt time.Time    `json:"stopped_at"`
	Status    TimersStatus `json:"status"`
}
