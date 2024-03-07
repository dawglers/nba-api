package query

import (
	endpoint "github.com/ndesai96/nba-api/endpoints"
	"github.com/ndesai96/nba-api/models/league"
	"github.com/ndesai96/nba-api/models/player"
)

type playerGameLogsBuilder struct {
	endpoint *endpoint.PlayerGameLogs
	gameLogs []*player.GameLog
}

func PlayerGameLogs(playerID int) *playerGameLogsBuilder {
	return &playerGameLogsBuilder{
		endpoint: endpoint.NewPlayerGameLogs(playerID),
	}
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
