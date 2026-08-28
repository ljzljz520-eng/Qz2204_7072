package service

import "training41/domain"

func (s *Service) AddUser(u domain.User) error {
	if e := domain.ValidateUser(u); e != nil {
		return e
	}
	return s.DB.SaveUser(u)
}
func (s *Service) User(id string) (domain.User, error) { return s.DB.GetUser(id) }
func (s *Service) AddTraining(t domain.Training) error {
	if e := domain.ValidateTraining(t); e != nil {
		return e
	}
	return s.DB.SaveTraining(t)
}
