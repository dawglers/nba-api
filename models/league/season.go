package league

import "fmt"

type Season int

func (s Season) String() string {
	return fmt.Sprintf("%d-%d", s-1, s%100)
}

type SeasonType string

const (
	RegularSeason SeasonType = "Regular Season"
	Preseason     SeasonType = "Pre Season"
	Playoffs      SeasonType = "Playoffs"
)
