package version1

import "time"

type TimeSession struct {
	Id        string        `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	User      string        `json:"user" bson:"user"`
	Tags      []string      `json:"tags" bson:"tags"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	Status    SessionStatus `json:"status" bson:"status"`
	Timers    []Timer       `json:"timers" bson:"timers"`
}
