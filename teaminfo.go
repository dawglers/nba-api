package nba

import (
	"strconv"

	"github.com/dawglers/nba-api/endpoint"
	"github.com/dawglers/nba-api/models/league"
	"github.com/dawglers/nba-api/models/team"
)

type teamInfoBuilder struct {
	endpoint endpoint.Endpoint
	teamID   int
	team     team.Team
}

func TeamInfo(teamID int) *teamInfoBuilder {
	teamInfoBuilder := &teamInfoBuilder{
		endpoint: endpoint.New(endpoint.TeamInfoCommon),
		teamID:   teamID,
	}
	teamInfoBuilder.SetDefaultParams()

	return teamInfoBuilder
}

func (t *teamInfoBuilder) LeagueID(leagueID league.ID) *teamInfoBuilder {
	t.endpoint.SetLeagueID(leagueID)
	return t
}

func (t *teamInfoBuilder) Season(season league.Season) *teamInfoBuilder {
	t.endpoint.SetSeason(season)
	return t
}

func (t *teamInfoBuilder) SeasonType(seasonType league.SeasonType) *teamInfoBuilder {
	t.endpoint.SetSeasonType(seasonType)
	return t
}

func (t *teamInfoBuilder) SetDefaultParams() {
	t.endpoint.SetTeamID(t.teamID)
	t.endpoint.SetLeagueID(league.NBA)
	t.endpoint.SetSeason(2024)
	t.endpoint.SetSeasonType(league.RegularSeason)
}

func (t *teamInfoBuilder) Execute() error {
	err := t.endpoint.Request()
	if err != nil {
		return err
	}

	for _, resultSet := range t.endpoint.GetResults().ResultSets {
		if resultSet.Name == "TeamInfoCommon" && len(resultSet.RowSet[0]) > 0 {
			teamBuilder := &team.TeamBuilder{}

			fieldNames := resultSet.Headers
			teamData := resultSet.RowSet[0]

			for i, fieldName := range fieldNames {
				switch fieldName {
				case "TEAM_ID":
					teamBuilder.ID(int(teamData[i].(float64)))
				case "TEAM_CITY":
					teamBuilder.City(teamData[i].(string))
				case "TEAM_NAME":
					teamBuilder.Name(teamData[i].(string))
				case "TEAM_ABBREVIATION":
					teamBuilder.Abbreviation(teamData[i].(string))
				case "TEAM_CONFERENCE":
					teamBuilder.Conference(teamData[i].(string))
				case "TEAM_DIVISION":
					teamBuilder.Division(teamData[i].(string))
				case "TEAM_CODE":
					teamBuilder.Code(teamData[i].(string))
				case "MIN_YEAR":
					inauguralYear, err := strconv.Atoi(teamData[i].(string))
					if err == nil {
						teamBuilder.InauguralYear(inauguralYear)
					}
				case "MAX_YEAR":
					latestYear, err := strconv.Atoi(teamData[i].(string))
					if err == nil {
						teamBuilder.LatestYear(latestYear)
					}
				}
			}

			t.team = teamBuilder.Build()
		}
	}

	return nil
}

func (t *teamInfoBuilder) GetTeam() team.Team {
	return t.team
}
