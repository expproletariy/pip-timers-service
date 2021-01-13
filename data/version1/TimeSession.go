package version1

import "time"

type TimeSession struct {
	Id        string        `json:"id"`
	Name      string        `json:"name"`
	User      string        `json:"user"`
	Tags      []string      `json:"tags"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Status    SessionStatus `json:"status"`
	Timers    []Timer       `json:"timers"`
}
