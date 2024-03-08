package player

import (
	"strconv"
)

type Player struct {
	ID        int
	FirstName string
	LastName  string
	Slug      string
	TeamID    int
	Position  Position
	Height    string
	Weight    int
	College   string
	Country   string
	DraftInfo DraftInfo
	Jersey    string
}

type PlayerBuilder struct {
	player Player
}

func (p *PlayerBuilder) Build() Player {
	return p.player
}

func (p *PlayerBuilder) FirstName(firstName string) *PlayerBuilder {
	p.player.FirstName = firstName
	return p
}

func (p *PlayerBuilder) LastName(lastName string) *PlayerBuilder {
	p.player.LastName = lastName
	return p
}

func (p *PlayerBuilder) ID(id int) *PlayerBuilder {
	p.player.ID = id
	return p
}

func NewPlayer(fieldNames []string, data []any) *Player {
	player := &Player{}

	for i, fieldName := range fieldNames {
		switch fieldName {
		case "PERSON_ID":
			player.ID = int(data[i].(float64))
		case "PLAYER_LAST_NAME":
			player.LastName = data[i].(string)
		case "PLAYER_FIRST_NAME":
			player.FirstName = data[i].(string)
		case "PLAYER_SLUG":
			player.Slug = data[i].(string)
		case "TEAM_ID":
			player.TeamID = int(data[i].(float64))
		case "JERSEY_NUMBER":
			if jersey, ok := data[i].(string); ok {
				player.Jersey = jersey
			}
		case "POSITION":
			player.Position = StringToPosition(data[i].(string))
		case "HEIGHT":
			player.Height = data[i].(string)
		case "WEIGHT":
			weight, err := strconv.Atoi(data[i].(string))
			if err == nil {
				player.Weight = weight
			}
		case "COUNTRY":
			player.Country = data[i].(string)
		case "COLLEGE":
			player.College = data[i].(string)
		case "DRAFT_YEAR":
			if draftYear, ok := data[i].(float64); ok {
				player.DraftInfo.Year = int(draftYear)
			} else {
				player.DraftInfo.Undrafted = true
			}
		case "DRAFT_ROUND":
			if draftRound, ok := data[i].(float64); ok {
				player.DraftInfo.Round = int(draftRound)
			} else {
				player.DraftInfo.Undrafted = true
			}
		case "DRAFT_NUMBER":
			if draftNumber, ok := data[i].(float64); ok {
				player.DraftInfo.Pick = int(draftNumber)
			} else {
				player.DraftInfo.Undrafted = true
			}
		}

	}

	return player
}
