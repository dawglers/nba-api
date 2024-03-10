package nba

import (
	"github.com/dawglers/nba-api/endpoint"
	"github.com/dawglers/nba-api/models/league"
)

type gamesBuilder struct {
	endpoint endpoint.Endpoint
	teamID   int
	gameIDs  []string
}

func GamesByTeam(teamID int) *gamesBuilder {
	gamesBuilder := &gamesBuilder{
		endpoint: endpoint.New(endpoint.CumeStatsTeamGames),
		teamID:   teamID,
	}
	gamesBuilder.SetDefaultParams()

	return gamesBuilder
}

func (g *gamesBuilder) LeagueID(leagueID league.ID) *gamesBuilder {
	g.endpoint.SetLeagueID(leagueID)
	return g
}

func (g *gamesBuilder) Season(season league.Season) *gamesBuilder {
	g.endpoint.SetSeason(season)
	return g
}

func (g *gamesBuilder) SeasonType(seasonType league.SeasonType) *gamesBuilder {
	g.endpoint.SetSeasonType(seasonType)
	return g
}

func (g *gamesBuilder) SetDefaultParams() {
	g.endpoint.SetLeagueID(league.NBA)
	g.endpoint.SetSeason(2024)
	g.endpoint.SetSeasonType(league.RegularSeason)
	g.endpoint.SetTeamID(g.teamID)
}

func (g *gamesBuilder) Execute() error {
	err := g.endpoint.Request()
	if err != nil {
		return err
	}

	for _, resultSet := range g.endpoint.GetResults().ResultSets {
		if resultSet.Name == "CumeStatsTeamGames" {
			fieldNames := resultSet.Headers

			for _, gameData := range resultSet.RowSet {
				for i, fieldName := range fieldNames {
					if fieldName == "GAME_ID" {
						g.gameIDs = append(g.gameIDs, gameData[i].(string))
					}
				}
			}
		}
	}

	return nil
}

func (g *gamesBuilder) GetGameIDs() []string {
	return g.gameIDs
}
