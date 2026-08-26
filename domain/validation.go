package domain

import "fmt"

func ValidateRecord(r Record) error {
	if r.ID == "" {
		return fmt.Errorf("record id required")
	}
	if r.TrainingID == "" {
		return fmt.Errorf("training id required")
	}
	if r.UserID == "" {
		return fmt.Errorf("user id required")
	}
	return nil
}
func ValidateUser(u User) error {
	if u.ID == "" || u.Name == "" {
		return fmt.Errorf("user identity required")
	}
	if u.Role == "" {
		return fmt.Errorf("role required")
	}
	return nil
}
func ValidateTraining(t Training) error {
	if t.ID == "" || t.Title == "" {
		return fmt.Errorf("training metadata required")
	}
	return nil
}
func NormalizeStatus(status string) string {
	switch status {
	case "pending", "confirmed", "archived":
		return status
	default:
		return "pending"
	}
}
