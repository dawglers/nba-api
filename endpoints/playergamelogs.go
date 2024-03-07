package endpoint

import (
	"net/url"
)

type PlayerGameLogs struct {
	Base
	playerID int
}

func NewPlayerGameLogs(playerID int) *PlayerGameLogs {
	playerGameLogs := &PlayerGameLogs{
		playerID: playerID,
	}
	playerGameLogs.SetDefaultParams()
	return playerGameLogs
}

func (p *PlayerGameLogs) SetDefaultParams() {
	p.Base.params = url.Values{}
	p.Base.SetPlayerID(p.playerID)
	p.Base.SetSeason(2024)
}

func (p *PlayerGameLogs) Request() error {
	return p.Base.Request(p.Path())
}

func (p PlayerGameLogs) Path() Path {
	return PlayerGameLogsPath
}

func (p PlayerGameLogs) GetResults() *Results {
	return p.Base.results
}
