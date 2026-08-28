package domain

import "testing"

func TestValidation(t *testing.T) {
	if ValidateUser(User{}) == nil {
		t.Fatal("expected error")
	}
	if NormalizeStatus("x") != "pending" {
		t.Fatal("normalization")
	}
	if !CanTransition("pending", "confirmed") {
		t.Fatal("transition")
	}
}
