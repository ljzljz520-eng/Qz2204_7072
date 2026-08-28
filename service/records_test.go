package service

import (
	"path/filepath"
	"testing"
	"training41/domain"
	"training41/store"
)

func TestRecordService(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	s := New(d)
	if e := s.Register(domain.NewRecord("r", "t", "u")); e != nil {
		t.Fatal(e)
	}
	if e := s.Confirm("r", "a"); e != nil {
		t.Fatal(e)
	}
	r, _ := s.Get("r")
	if len(r.Confirmations) != 1 {
		t.Fatal("confirmation")
	}
}
