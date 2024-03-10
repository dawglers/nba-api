package nba

import (
	"time"

	"github.com/dawglers/nba-api/endpoint"
	"github.com/dawglers/nba-api/models/league"
)

type scoreboardBuilder struct {
	endpoint endpoint.Endpoint
	gameIDs  []string
}

func Scoreboard() *scoreboardBuilder {
	scoreboardBuilder := &scoreboardBuilder{
		endpoint: endpoint.New(endpoint.Scoreboard),
	}
	scoreboardBuilder.SetDefaultParams()

	return scoreboardBuilder
}

func (s *scoreboardBuilder) LeagueID(leagueID league.ID) *scoreboardBuilder {
	s.endpoint.SetLeagueID(leagueID)
	return s
}

func (s *scoreboardBuilder) Date(date time.Time) *scoreboardBuilder {
	s.endpoint.SetGameDate(date)
	return s
}

func (s *scoreboardBuilder) SetDefaultParams() {
	s.endpoint.SetLeagueID(league.NBA)
	s.endpoint.SetGameDate(time.Now())
	s.endpoint.SetDayOffset(0)
}

func (s *scoreboardBuilder) Execute() error {
	err := s.endpoint.Request()
	if err != nil {
		return err
	}

	for _, resultSet := range s.endpoint.GetResults().ResultSets {
		if resultSet.Name == "Available" {
			fieldNames := resultSet.Headers

			for _, gameData := range resultSet.RowSet {
				for i, fieldName := range fieldNames {
					if fieldName == "GAME_ID" {
						s.gameIDs = append(s.gameIDs, gameData[i].(string))
					}
				}
			}
		}
	}

	return nil
}

func (s *scoreboardBuilder) GetGameIDs() []string {
	return s.gameIDs
}
