package service

import (
	"path/filepath"
	"testing"
	"training41/domain"
	"training41/store"
)

func TestUsersAndTraining(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	s := New(d)
	if e := s.AddUser(domain.User{ID: "u", Name: "A", Role: "learner", Active: true}); e != nil {
		t.Fatal(e)
	}
	if e := s.AddTraining(domain.Training{ID: "t", Title: "Training 41"}); e != nil {
		t.Fatal(e)
	}
}
