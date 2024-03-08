package player

type Player struct {
	ID        int
	FirstName string
	LastName  string
	Slug      string
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

func (p *PlayerBuilder) ID(id int) *PlayerBuilder {
	p.player.ID = id
	return p
}

func (p *PlayerBuilder) FirstName(firstName string) *PlayerBuilder {
	p.player.FirstName = firstName
	return p
}

func (p *PlayerBuilder) LastName(lastName string) *PlayerBuilder {
	p.player.LastName = lastName
	return p
}

func (p *PlayerBuilder) Slug(slug string) *PlayerBuilder {
	p.player.Slug = slug
	return p
}

func (p *PlayerBuilder) Jersey(jersey string) *PlayerBuilder {
	p.player.Jersey = jersey
	return p
}

func (p *PlayerBuilder) Position(position Position) *PlayerBuilder {
	p.player.Position = position
	return p
}

func (p *PlayerBuilder) Height(height string) *PlayerBuilder {
	p.player.Height = height
	return p
}

func (p *PlayerBuilder) Weight(weight int) *PlayerBuilder {
	p.player.Weight = weight
	return p
}

func (p *PlayerBuilder) Country(country string) *PlayerBuilder {
	p.player.Country = country
	return p
}

func (p *PlayerBuilder) College(college string) *PlayerBuilder {
	p.player.College = college
	return p
}

func (p *PlayerBuilder) DraftInfo(draftInfo DraftInfo) *PlayerBuilder {
	p.player.DraftInfo = draftInfo
	return p
}
