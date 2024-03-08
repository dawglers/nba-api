package endpoint

import (
	"github.com/ndesai96/nba-api/models/league"
	"github.com/ndesai96/nba-api/models/player"
)

type Path string

const (
	UndefinedPath           = ""
	PlayerIndexPath    Path = "playerindex"
	PlayerGameLogsPath Path = "playergamelogs"
	ListPlayersPath    Path = "cumestatsplayerlist"
)

type Endpoint interface {
	Request() error
	Path() Path

	// Param setters
	SetPlayerID(playerID int)
	SetTeamID(teamID int)
	SetLeagueID(leagueID league.ID)
	SetOnlyAllStar(onlyAllStar bool)
	SetPlayerPosition(playerPosition player.Position)
	SetSeason(season league.Season)
	SetHistorical(includeHistorical bool)
}
