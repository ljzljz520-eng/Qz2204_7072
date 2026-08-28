package service

import (
	"fmt"
	"sync"
	"time"
	"training41/domain"
	"training41/store"
)

type Service struct {
	DB       *store.DB
	mu       sync.Mutex
	required int
}

func New(db *store.DB) *Service { return &Service{DB: db, required: 2} }
func (s *Service) Register(r domain.Record) error {
	if e := domain.ValidateRecord(r); e != nil {
		return e
	}
	r.Status = "pending"
	return s.DB.SaveRecord(r)
}
func (s *Service) Confirm(id, actor string) error {
	r, e := s.DB.GetRecord(id)
	if e != nil {
		return e
	}
	if r.ConfirmedBy(actor) {
		return nil
	}
	time.Sleep(15 * time.Millisecond)
	r.Confirmations = append(r.Confirmations, actor)
	r.UpdatedAt = time.Now().UTC()
	if r.IsComplete(s.required) {
		r.Status = "confirmed"
	}
	return s.DB.SaveRecord(r)
}
func (s *Service) Archive(id, actor string) error {
	r, e := s.DB.GetRecord(id)
	if e != nil {
		return e
	}
	if r.Status != "confirmed" {
		return fmt.Errorf("record not complete")
	}
	r.Status = "archived"
	r.UpdatedAt = time.Now().UTC()
	return s.DB.SaveRecord(r)
}
func (s *Service) Get(id string) (domain.Record, error) { return s.DB.GetRecord(id) }
func (s *Service) SetRequired(n int) {
	if n > 0 {
		s.required = n
	}
}
