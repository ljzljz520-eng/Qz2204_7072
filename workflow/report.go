package workflow

import (
	"strings"
	"training41/domain"
)

func Summarize(r domain.Record) string {
	return strings.Join([]string{r.ID, r.TrainingID, r.UserID, r.Status}, "|")
}
func Eligible(r domain.Record) bool { return r.Status == "confirmed" || r.Status == "archived" }
