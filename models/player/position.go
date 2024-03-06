package player

type Position string

const (
	AllPositions Position = ""
	Forward      Position = "F"
	Center       Position = "C"
	Guard        Position = "G"
)

func StringToPosition(s string) Position {
	switch s {
	case "F":
		return Forward
	case "C":
		return Center
	case "G":
		return Guard
	default:
		return AllPositions
	}
}
