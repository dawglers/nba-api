package nba

import (
	"strconv"

	"github.com/dawglers/nba-api/endpoint"
	"github.com/dawglers/nba-api/models/league"
	"github.com/dawglers/nba-api/models/team"
)

type teamsBuilder struct {
	endpoint endpoint.Endpoint
	teamID   int
	team     team.Team
}

func TeamInfo(teamID int) *teamsBuilder {
	teamsBuilder := &teamsBuilder{
		endpoint: endpoint.New(endpoint.TeamInfoCommon),
		teamID:   teamID,
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
	t.endpoint.SetTeamID(t.teamID)
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

func (t *teamsBuilder) GetTeam() team.Team {
	return t.team
}
