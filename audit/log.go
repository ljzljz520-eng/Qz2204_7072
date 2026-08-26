package audit

import (
	"fmt"
	"time"
	"training41/domain"
	"training41/store"
)

type Logger struct{ DB *store.DB }

func New(db *store.DB) *Logger { return &Logger{DB: db} }
func (l *Logger) Write(record, action, actor, detail string) error {
	v := domain.Audit{ID: fmt.Sprintf("%d", time.Now().UnixNano()), RecordID: record, Action: action, Actor: actor, Detail: detail, At: time.Now().UTC()}
	return l.DB.SaveAudit(v)
}
func (l *Logger) Event(record, kind, actor string) error {
	v := domain.Event{ID: fmt.Sprintf("%d", time.Now().UnixNano()), RecordID: record, Kind: kind, Actor: actor, At: time.Now().UTC()}
	return l.DB.SaveEvent(v)
}
