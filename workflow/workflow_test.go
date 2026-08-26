package workflow

import (
	"path/filepath"
	"testing"
	"training41/domain"
	"training41/service"
	"training41/store"
)

func TestWorkflowOne(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	s := service.New(d)
	if e := Intake(s, domain.NewRecord("r", "t", "u")); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	s := service.New(d)
	Intake(s, domain.NewRecord("r", "t", "u"))
	s.SetRequired(1)
	if e := Review(s, "r", "a"); e != nil {
		t.Fatal(e)
	}
	if e := Archive(s, "r", "a"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	s := service.New(d)
	Intake(s, domain.NewRecord("r", "t", "u"))
	if e := Process(s, "r", "a"); e != nil {
		t.Fatal(e)
	}
}
