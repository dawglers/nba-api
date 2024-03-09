package query

import (
	"github.com/dawglers/nba-api/endpoint"
	"github.com/dawglers/nba-api/models/league"
	"github.com/dawglers/nba-api/models/team"
)

type listTeamsBuilder struct {
	endpoint endpoint.Endpoint
	teams    []team.Team
}

func ListTeams() *listTeamsBuilder {
	listTeamsBuilder := &listTeamsBuilder{
		endpoint: endpoint.New(endpoint.ListTeams),
	}
	listTeamsBuilder.SetDefaultParams()

	return listTeamsBuilder
}

func (l *listTeamsBuilder) LeagueID(leagueID league.ID) *listTeamsBuilder {
	l.endpoint.SetLeagueID(leagueID)
	return l
}

func (l *listTeamsBuilder) Season(season league.Season) *listTeamsBuilder {
	l.endpoint.SetSeason(season)
	return l
}

func (l *listTeamsBuilder) SeasonType(seasonType league.SeasonType) *listTeamsBuilder {
	l.endpoint.SetSeasonType(seasonType)
	return l
}

func (l *listTeamsBuilder) SetDefaultParams() {
	l.endpoint.SetLeagueID(league.NBA)
	l.endpoint.SetSeason(2024)
	l.endpoint.SetSeasonType(league.RegularSeason)
}

func (l *listTeamsBuilder) Execute() error {
	err := l.endpoint.Request()
	if err != nil {
		return err
	}

	data := l.endpoint.GetResults().ResultSets[0]

	var name string
	var teamID int
	for _, teamData := range data.RowSet {
		t := &team.TeamBuilder{}
		if teamName, ok := teamData[0].(string); ok {
			name = teamName
		}
		if teammIDFloat, ok := teamData[1].(float64); ok {
			teamID = int(teammIDFloat)
		}

		team := t.
			ID(teamID).
			Name(name).
			Build()

		l.teams = append(l.teams, team)
	}

	return nil
}

func (l *listTeamsBuilder) GetTeams() []team.Team {
	return l.teams
}
