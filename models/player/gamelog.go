package player

import (
	"math"
	"regexp"

	"github.com/ndesai96/nba-api/models/game"
)

type GameLog struct {
	PlayerID              int
	GameID                string
	Matchup               string
	Minutes               int
	FieldGoalMade         int
	FieldGoalAttempted    int
	ThreePointerMade      int
	ThreePointerAttempted int
	FreeThrowMade         int
	FreeThrowAttempted    int
	OffensiveRebound      int
	DefensiveRebound      int
	Assists               int
	Turnovers             int
	Steals                int
	Blocks                int
	Fouls                 int
	Points                int
	PlusMinus             int
	NBAFantasyPoints      int
}

func (g GameLog) Rebounds() int {
	return g.OffensiveRebound + g.DefensiveRebound
}

func (g GameLog) Location() game.Location {
	pattern := regexp.MustCompile(`@`)

	if pattern.MatchString(g.Matchup) {
		return game.Away
	}
	return game.Home
}

func NewGameLog(fieldNames []string, data []any) *GameLog {
	gameLog := &GameLog{}

	for i, fieldName := range fieldNames {
		switch fieldName {
		case "PLAYER_ID":
			gameLog.PlayerID = int(data[i].(float64))
		case "GAME_ID":
			gameLog.GameID = data[i].(string)
		case "MATCHUP":
			gameLog.Matchup = data[i].(string)
		case "MIN":
			gameLog.Minutes = int(math.Round(data[i].(float64)))
		case "FGM":
			gameLog.FieldGoalMade = int(data[i].(float64))
		case "FGA":
			gameLog.FieldGoalAttempted = int(data[i].(float64))
		case "FG3M":
			gameLog.ThreePointerMade = int(data[i].(float64))
		case "FG3A":
			gameLog.ThreePointerAttempted = int(data[i].(float64))
		case "FTM":
			gameLog.FreeThrowMade = int(data[i].(float64))
		case "OREB":
			gameLog.OffensiveRebound = int(data[i].(float64))
		case "DREB":
			gameLog.DefensiveRebound = int(data[i].(float64))
		case "AST":
			gameLog.Assists = int(data[i].(float64))
		case "TOV":
			gameLog.Turnovers = int(data[i].(float64))
		case "STL":
			gameLog.Steals = int(data[i].(float64))
		case "BLK":
			gameLog.Blocks = int(data[i].(float64))
		case "PF":
			gameLog.Fouls = int(data[i].(float64))
		case "PTS":
			gameLog.Points = int(data[i].(float64))
		case "PLUS_MINUS":
			gameLog.PlusMinus = int(data[i].(float64))
		case "NBA_FANTASY_PTS":
			gameLog.NBAFantasyPoints = int(data[i].(float64))
		}
	}

	return gameLog
}
