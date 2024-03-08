package league

import "fmt"

type Season int

func (s Season) String() string {
	if s%100 < 10 {
		return fmt.Sprintf("%d-0%d", s-1, s%100)
	}
	return fmt.Sprintf("%d-%d", s-1, s%100)
}

type SeasonType string

const (
	RegularSeason SeasonType = "Regular Season"
	Preseason     SeasonType = "Pre Season"
	Playoffs      SeasonType = "Playoffs"
)
