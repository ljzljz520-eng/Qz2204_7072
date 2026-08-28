package workflow

import (
	"fmt"
	"training41/domain"
	"training41/service"
)

func Intake(s *service.Service, r domain.Record) error {
	if e := s.Register(r); e != nil {
		return e
	}
	return nil
}
func Review(s *service.Service, id, actor string) error {
	if actor == "" {
		return fmt.Errorf("review actor required")
	}
	return s.Confirm(id, actor)
}
func Archive(s *service.Service, id, actor string) error         { return s.Archive(id, actor) }
func Query(s *service.Service, id string) (domain.Record, error) { return s.Get(id) }
