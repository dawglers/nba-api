package endpoint

import (
	"net/url"

	"github.com/ndesai96/nba-api/models/league"
)

type ListTeams struct {
	Base
}

func NewListTeams() *ListTeams {
	listTeams := &ListTeams{}
	listTeams.SetDefaultParams()
	return listTeams
}

func (l *ListTeams) SetDefaultParams() {
	l.Base.params = url.Values{}
	l.Base.SetSeason(2024)
	l.Base.SetLeagueID(league.NBA)
	l.Base.SetSeasonType(league.RegularSeason)
}

func (p *ListTeams) Request() error {
	return p.Base.Request(p.Path())
}

func (p ListTeams) Path() Path {
	return ListTeamsPath
}

func (p ListTeams) GetResults() *Results {
	return p.Base.results
}
