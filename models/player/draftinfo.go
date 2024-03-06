package player

import "fmt"

type DraftInfo struct {
	Undrafted bool
	Year      int
	Round     int
	Pick      int
}

func (d DraftInfo) String() string {
	if d.Undrafted {
		return "Undrafted"
	}
	return fmt.Sprintf("%d: Rd %d, Pk %d", d.Year, d.Round, d.Pick)
}
