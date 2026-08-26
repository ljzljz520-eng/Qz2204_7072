package domain

import "testing"

func TestRecordRules(t *testing.T) {
	r := NewRecord("r", "t", "u")
	if e := ValidateRecord(r); e != nil {
		t.Fatal(e)
	}
	if r.ConfirmedBy("x") {
		t.Fatal("unexpected")
	}
}
