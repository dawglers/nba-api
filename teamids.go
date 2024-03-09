package nba

import (
	"github.com/dawglers/nba-api/endpoint"
	"github.com/dawglers/nba-api/models/league"
)

type teamsBuilder struct {
	endpoint endpoint.Endpoint
	teamIDs  []int
}

func Teams() *teamsBuilder {
	teamsBuilder := &teamsBuilder{
		endpoint: endpoint.New(endpoint.CumeStatsTeamList),
	}
	teamsBuilder.SetDefaultParams()

	return teamsBuilder
}

func (t *teamsBuilder) LeagueID(leagueID league.ID) *teamsBuilder {
	t.endpoint.SetLeagueID(leagueID)
	return t
}

func (t *teamsBuilder) Season(season league.Season) *teamsBuilder {
	t.endpoint.SetSeason(season)
	return t
}

func (t *teamsBuilder) SeasonType(seasonType league.SeasonType) *teamsBuilder {
	t.endpoint.SetSeasonType(seasonType)
	return t
}

func (t *teamsBuilder) SetDefaultParams() {
	t.endpoint.SetLeagueID(league.NBA)
	t.endpoint.SetSeason(2024)
	t.endpoint.SetSeasonType(league.RegularSeason)
}

func (t *teamsBuilder) Execute() error {
	err := t.endpoint.Request()
	if err != nil {
		return err
	}

	for _, resultSet := range t.endpoint.GetResults().ResultSets {
		if resultSet.Name == "CumeStatsTeamList" {
			fieldNames := resultSet.Headers

			for _, teamData := range resultSet.RowSet {
				for i, fieldName := range fieldNames {
					if fieldName == "TEAM_ID" {
						t.teamIDs = append(t.teamIDs, int(teamData[i].(float64)))
					}
				}
			}
		}
	}

	return nil
}

func (t *teamsBuilder) GetTeamIDs() []int {
	return t.teamIDs
}
