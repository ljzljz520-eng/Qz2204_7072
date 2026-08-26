package store

import "training41/domain"

func CountComplete(rs []domain.Record) int {
	n := 0
	for _, r := range rs {
		if r.Status == "confirmed" || r.Status == "archived" {
			n++
		}
	}
	return n
}
func CountForTraining(rs []domain.Record, id string) int {
	n := 0
	for _, r := range rs {
		if r.TrainingID == id {
			n++
		}
	}
	return n
}
