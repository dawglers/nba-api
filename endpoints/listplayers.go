package endpoint

import (
	"net/url"

	"github.com/ndesai96/nba-api/models/league"
)

type ListPlayers struct {
	Base
}

func NewListPlayers() *ListPlayers {
	listPlayers := &ListPlayers{}
	listPlayers.SetDefaultParams()
	return listPlayers
}

func (l *ListPlayers) SetDefaultParams() {
	l.Base.params = url.Values{}
	l.Base.SetSeason(2024)
	l.Base.SetLeagueID(league.NBA)
	l.Base.SetSeasonType(league.RegularSeason)
}

func (p *ListPlayers) Request() error {
	return p.Base.Request(p.Path())
}

func (p ListPlayers) Path() Path {
	return ListPlayersPath
}

func (p ListPlayers) GetResults() *Results {
	return p.Base.results
}
