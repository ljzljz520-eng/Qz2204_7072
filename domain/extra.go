package domain

func StatusRank(s string) int {
	switch s {
	case "pending":
		return 1
	case "confirmed":
		return 2
	case "archived":
		return 3
	default:
		return 0
	}
}
func CanTransition(from, to string) bool {
	if from == to {
		return true
	}
	return StatusRank(to) == StatusRank(from)+1
}
