package store

import (
	"path/filepath"
	"testing"
	"training41/domain"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "data.db")
	d, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := domain.NewRecord("r1", "training41", "u1")
	if e = d.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	d.Close()
	d, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer d.Close()
	if _, e = d.GetRecord("r1"); e != nil {
		t.Fatal(e)
	}
}
