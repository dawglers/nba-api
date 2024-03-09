package endpoint

import (
	"strconv"

	"github.com/dawglers/nba-api/models/league"
	"github.com/dawglers/nba-api/models/player"
)

type ParamSetter interface {
	SetPlayerID(playerID int)
	SetTeamID(teamID int)
	SetLeagueID(leagueID league.ID)
	SetOnlyAllStar(onlyAllStar bool)
	SetPosition(position player.Position)
	SetSeason(season league.Season)
	SetSeasonType(seasonType league.SeasonType)
	SetHistorical(includeHistorical bool)
	SetOnlyActive(onlyActive bool)
	SetLastNGames(lastNGames int)
	SetOpponentTeamID(opponentTeamID int)
}

func (b *base) SetPlayerID(playerID int) {
	b.params.Set("PlayerID", strconv.Itoa(playerID))
}

func (b *base) SetTeamID(teamID int) {
	b.params.Set("TeamID", strconv.Itoa(teamID))
}

func (b *base) SetLeagueID(leagueID league.ID) {
	b.params.Set("LeagueID", string(leagueID))
}

func (b *base) SetOnlyAllStar(onlyAllStar bool) {
	b.params.Set("AllStar", strconv.Itoa(boolToInt(onlyAllStar)))
}

func (b *base) SetPosition(position player.Position) {
	b.params.Set("PlayerPosition", string(position))
}

func (b *base) SetSeason(season league.Season) {
	b.params.Set("Season", season.String())
}

func (b *base) SetSeasonType(seasonType league.SeasonType) {
	b.params.Set("SeasonType", string(seasonType))
}

func (b *base) SetHistorical(includeHistorical bool) {
	b.params.Set("Historical", strconv.Itoa(boolToInt(includeHistorical)))
}

func (b *base) SetOnlyActive(onlyActive bool) {
	b.params.Set("Active", strconv.Itoa(boolToInt(onlyActive)))
}

func (b *base) SetLastNGames(lastNGames int) {
	b.params.Set("LastNGames", strconv.Itoa(lastNGames))
}

func (b *base) SetOpponentTeamID(opponentTeamID int) {
	b.params.Set("OpponentTeamID", strconv.Itoa(opponentTeamID))
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
