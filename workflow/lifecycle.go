package workflow

import (
	"training41/domain"
	"training41/service"
)

func Process(s *service.Service, id, actor string) error { return Review(s, id, actor) }
func Close(s *service.Service, id, actor string) error   { return Archive(s, id, actor) }
func Reopen(r domain.Record) domain.Record               { r.Status = "pending"; return r }
