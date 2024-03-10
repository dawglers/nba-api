package nba

import (
	"strconv"

	"github.com/dawglers/nba-api/endpoint"
	"github.com/dawglers/nba-api/models/league"
	"github.com/dawglers/nba-api/models/player"
)

type playersBuilder struct {
	endpoint endpoint.Endpoint
	players  []player.Player
	teamID   int
}

func Players() *playersBuilder {
	playersBuilder := &playersBuilder{
		endpoint: endpoint.New(endpoint.PlayerIndex),
	}
	playersBuilder.SetDefaultParams()

	return playersBuilder
}

func (p *playersBuilder) TeamID(teamID int) *playersBuilder {
	p.teamID = teamID
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
	if p.teamID != 0 {
		p.endpoint.SetHistorical(false)
	}

	err := p.endpoint.Request()
	if err != nil {
		return err
	}

	data := p.endpoint.GetResults().ResultSets[0]
	fieldNames := data.Headers

	for _, playerData := range data.RowSet {
		playerBuilder := &player.PlayerBuilder{}
		draftInfoBuidler := &player.DraftInfoBuilder{}

		for i, fieldName := range fieldNames {
			switch fieldName {
			case "PERSON_ID":
				playerBuilder.ID(int(playerData[i].(float64)))
			case "PLAYER_LAST_NAME":
				playerBuilder.LastName(playerData[i].(string))
			case "PLAYER_FIRST_NAME":
				playerBuilder.LastName(playerData[i].(string))
			case "PLAYER_SLUG":
				playerBuilder.LastName(playerData[i].(string))
			case "JERSEY_NUMBER":
				if jersey, ok := playerData[i].(string); ok {
					playerBuilder.Jersey(jersey)
				}
			case "POSITION":
				if positionString, ok := playerData[i].(string); ok {
					position := player.ToPosition(positionString)
					playerBuilder.Position(position)
				}
			case "HEIGHT":
				if height, ok := playerData[i].(string); ok {
					playerBuilder.Height(height)
				}
			case "WEIGHT":
				if weightString, ok := playerData[i].(string); ok {
					weight, err := strconv.Atoi(weightString)
					if err == nil {
						playerBuilder.Weight(weight)
					}
				}
			case "COUNTRY":
				if country, ok := playerData[i].(string); ok {
					playerBuilder.Country(country)
				}
			case "COLLEGE":
				if college, ok := playerData[i].(string); ok {
					playerBuilder.College(college)
				}
			case "DRAFT_YEAR":
				if draftYear, ok := playerData[i].(float64); ok {
					draftInfoBuidler.Year(int(draftYear))
				}
			case "DRAFT_ROUND":
				if draftRound, ok := playerData[i].(float64); ok {
					draftInfoBuidler.Round(int(draftRound))
				}
			case "DRAFT_NUMBER":
				if draftPick, ok := playerData[i].(float64); ok {
					draftInfoBuidler.Pick(int(draftPick))
				}
			}
		}

		playerBuilder.DraftInfo(draftInfoBuidler.Build())
		p.players = append(p.players, playerBuilder.Build())
	}

	return nil
}

func (p *playersBuilder) GetPlayers() []player.Player {
	return p.players
}
