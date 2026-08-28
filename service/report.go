package service

import (
	"training41/domain"
	"training41/store"
)

func (s *Service) Records(status string) ([]domain.Record, error) {
	rs, e := s.DB.ListRecords()
	if e != nil {
		return nil, e
	}
	if status != "" {
		rs = store.FilterByStatus(rs, status)
	}
	return store.SortRecords(rs), nil
}
func (s *Service) Completion(id string) (bool, error) {
	r, e := s.Get(id)
	if e != nil {
		return false, e
	}
	return r.IsComplete(s.required), nil
}
