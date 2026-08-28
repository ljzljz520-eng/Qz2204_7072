package store

import (
	"sort"
	"training41/domain"
)

func SortRecords(rs []domain.Record) []domain.Record {
	sort.Slice(rs, func(i, j int) bool { return rs[i].UpdatedAt.Before(rs[j].UpdatedAt) })
	return rs
}
func FilterByStatus(rs []domain.Record, status string) []domain.Record {
	out := []domain.Record{}
	for _, r := range rs {
		if r.Status == status {
			out = append(out, r)
		}
	}
	return out
}
func FindByUser(rs []domain.Record, user string) []domain.Record {
	out := []domain.Record{}
	for _, r := range rs {
		if r.UserID == user {
			out = append(out, r)
		}
	}
	return out
}
