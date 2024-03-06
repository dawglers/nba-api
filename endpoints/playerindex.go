package endpoint

import (
	"net/url"

	"github.com/ndesai96/nba-api/models/league"
)

type PlayerIndex struct {
	Base
}

func NewPlayerIndex() *PlayerIndex {
	playerIndex := &PlayerIndex{}
	playerIndex.SetDefaultParams()
	return playerIndex
}

func (p *PlayerIndex) SetDefaultParams() {
	p.Base.params = url.Values{}
	p.Base.SetHistorical(true)
	p.Base.SetLeagueID(league.NBA)
	p.Base.SetSeason(2024)
	p.Base.SetOnlyActive(true)
}

func (p *PlayerIndex) Request() error {
	return p.Base.Request(p.Path())
}

func (p PlayerIndex) Path() Path {
	return PlayerIndexPath
}

func (p PlayerIndex) GetResults() *Results {
	return p.Base.results
}
