package query

import (
	"github.com/ndesai96/nba-api/endpoint"
	"github.com/ndesai96/nba-api/models/league"
	"github.com/ndesai96/nba-api/models/player"
)

type playerGameLogsBuilder struct {
	endpoint endpoint.Endpoint
	playerID int
	gameLogs []*player.GameLog
}

func PlayerGameLogs(playerID int) *playerGameLogsBuilder {
	playerGameLogsBuilder := &playerGameLogsBuilder{
		endpoint: endpoint.New(endpoint.PlayerGameLogs),
		playerID: playerID,
	}
	playerGameLogsBuilder.SetDefaultParams()

	return playerGameLogsBuilder
}

func (p *playerGameLogsBuilder) LastNGames(lastNGames int) *playerGameLogsBuilder {
	p.endpoint.SetLastNGames(lastNGames)
	return p
}

func (p *playerGameLogsBuilder) OpponentTeamID(opponentTeamID int) *playerGameLogsBuilder {
	p.endpoint.SetOpponentTeamID(opponentTeamID)
	return p
}

func (p *playerGameLogsBuilder) Season(season league.Season) *playerGameLogsBuilder {
	p.endpoint.SetSeason(season)
	return p
}

func (p *playerGameLogsBuilder) SetDefaultParams() {
	p.endpoint.SetPlayerID(p.playerID)
	p.endpoint.SetSeason(2024)
}

func (p *playerGameLogsBuilder) Execute() error {
	err := p.endpoint.Request()
	if err != nil {
		return err
	}

	data := p.endpoint.GetResults().ResultSets[0]
	fieldNames := data.Headers

	for _, gameLogData := range data.RowSet {
		gameLog := player.NewGameLog(fieldNames, gameLogData)
		p.gameLogs = append(p.gameLogs, gameLog)
	}

	return nil
}

func (p *playerGameLogsBuilder) GetGameLogs() []*player.GameLog {
	return p.gameLogs
}
