package domain

import "time"

type Record struct {
	ID, TrainingID, UserID, Status string
	Confirmations                  []string
	CreatedAt, UpdatedAt           time.Time
}
type User struct {
	ID, Name, Role string
	Active         bool
}
type Event struct {
	ID, RecordID, Kind, Actor string
	At                        time.Time
	Payload                   string
}
type Audit struct {
	ID, RecordID, Action, Actor, Detail string
	At                                  time.Time
}
type Training struct {
	ID, Title, Version string
	Required           bool
}

func NewRecord(id, training, user string) Record {
	now := time.Now().UTC()
	return Record{ID: id, TrainingID: training, UserID: user, Status: "pending", CreatedAt: now, UpdatedAt: now}
}
func (r Record) ConfirmedBy(user string) bool {
	for _, v := range r.Confirmations {
		if v == user {
			return true
		}
	}
	return false
}
func (r Record) IsComplete(required int) bool { return len(r.Confirmations) >= required }
