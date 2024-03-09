package nba

import (
	"strings"

	"github.com/dawglers/nba-api/endpoint"
	"github.com/dawglers/nba-api/models/league"
	"github.com/dawglers/nba-api/models/player"
)

type listPlayersBuilder struct {
	endpoint endpoint.Endpoint
	players  []player.Player
}

func ListPlayers() *listPlayersBuilder {
	listPlayersBuilder := &listPlayersBuilder{
		endpoint: endpoint.New(endpoint.ListPlayers),
	}
	listPlayersBuilder.SetDefaultParams()

	return listPlayersBuilder
}

func (l *listPlayersBuilder) LeagueID(leagueID league.ID) *listPlayersBuilder {
	l.endpoint.SetLeagueID(leagueID)
	return l
}

func (l *listPlayersBuilder) Season(season league.Season) *listPlayersBuilder {
	l.endpoint.SetSeason(season)
	return l
}

func (l *listPlayersBuilder) SeasonType(seasonType league.SeasonType) *listPlayersBuilder {
	l.endpoint.SetSeasonType(seasonType)
	return l
}

func (l *listPlayersBuilder) SetDefaultParams() {
	l.endpoint.SetLeagueID(league.NBA)
	l.endpoint.SetSeason(2024)
	l.endpoint.SetSeasonType(league.RegularSeason)
}

func (l *listPlayersBuilder) Execute() error {
	err := l.endpoint.Request()
	if err != nil {
		return err
	}

	data := l.endpoint.GetResults().ResultSets[0]

	var firstName, lastName string
	var playerID int
	for _, playerData := range data.RowSet {
		p := &player.PlayerBuilder{}
		if playerName, ok := playerData[0].(string); ok {
			parts := strings.Split(playerName, ", ")
			if len(parts) == 2 {
				firstName = parts[1]
				lastName = parts[0]
			}
		}
		if playerIDFloat, ok := playerData[1].(float64); ok {
			playerID = int(playerIDFloat)
		}

		player := p.
			ID(playerID).
			FirstName(firstName).
			LastName(lastName).
			Build()

		l.players = append(l.players, player)
	}

	return nil
}

func (l *listPlayersBuilder) GetPlayers() []player.Player {
	return l.players
}
