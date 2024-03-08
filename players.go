package query

import (
	"github.com/ndesai96/nba-api/endpoint"
	"github.com/ndesai96/nba-api/models/league"
	"github.com/ndesai96/nba-api/models/player"
)

type playersBuilder struct {
	endpoint endpoint.Endpoint
	players  []*player.Player
}

func Players() *playersBuilder {
	playersBuilder := &playersBuilder{
		endpoint: endpoint.New(endpoint.PlayerIndex),
	}
	playersBuilder.SetDefaultParams()

	return playersBuilder
}

func (p *playersBuilder) TeamID(teamID int) *playersBuilder {
	p.endpoint.SetTeamID(teamID)
	return p
}

func (p *playersBuilder) LeagueID(leagueID league.ID) *playersBuilder {
	p.endpoint.SetLeagueID(leagueID)
	return p
}

func (p *playersBuilder) Season(season league.Season) *playersBuilder {
	p.endpoint.SetSeason(season)
	return p
}

func (p *playersBuilder) Position(position player.Position) *playersBuilder {
	p.endpoint.SetPosition(position)
	return p
}

func (p *playersBuilder) OnlyAllStar(onlyAllStar bool) *playersBuilder {
	p.endpoint.SetOnlyAllStar(onlyAllStar)
	return p
}

func (p *playersBuilder) OnlyActive(onlyActive bool) *playersBuilder {
	p.endpoint.SetOnlyActive(onlyActive)
	return p
}

func (p *playersBuilder) SetDefaultParams() {
	p.endpoint.SetHistorical(true)
	p.endpoint.SetLeagueID(league.NBA)
	p.endpoint.SetSeason(2024)
	p.endpoint.SetOnlyActive(true)
}

func (p *playersBuilder) Execute() error {
	err := p.endpoint.Request()
	if err != nil {
		return err
	}

	data := p.endpoint.GetResults().ResultSets[0]
	fieldNames := data.Headers

	for _, playerData := range data.RowSet {
		player := player.NewPlayer(fieldNames, playerData)
		p.players = append(p.players, player)
	}

	return nil
}

func (p *playersBuilder) GetPlayers() []*player.Player {
	return p.players
}
