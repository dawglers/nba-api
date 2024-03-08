package player

import (
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/ndesai96/nba-api/models/game"
)

const dateLayout = "2006-01-02T15:04:05"

type GameLog struct {
	PlayerID              int
	GameID                int
	TeamID                int
	Matchup               string
	GameDate              time.Time
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
			gameID, err := strconv.Atoi(data[i].(string))
			if err == nil {
				gameLog.GameID = gameID
			}
		case "TEAM_ID":
			gameLog.TeamID = int(data[i].(float64))
		case "GAME_DATE":
			dateString := data[i].(string)
			gameDate, err := time.Parse(dateLayout, dateString)
			if err == nil {
				gameLog.GameDate = gameDate
			}
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
