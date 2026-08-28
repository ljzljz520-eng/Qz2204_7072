package store

import (
	"testing"
	"training41/domain"
)

func TestQueries(t *testing.T) {
	rs := []domain.Record{{ID: "1", Status: "confirmed", UserID: "u"}, {ID: "2", Status: "pending", UserID: "u"}}
	if len(FilterByStatus(rs, "confirmed")) != 1 {
		t.Fatal("filter")
	}
	if CountForTraining(rs, "") != 2 {
		t.Fatal("count")
	}
}
