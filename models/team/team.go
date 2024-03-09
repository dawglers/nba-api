package team

type Team struct {
	ID            int
	City          string
	Name          string
	Abbreviation  string
	Conference    string
	Division      string
	Code          string
	Slug          string
	InauguralYear int
	LatestYear    int
}

type TeamBuilder struct {
	team Team
}

func (t *TeamBuilder) Build() Team {
	return t.team
}

func (t *TeamBuilder) ID(id int) *TeamBuilder {
	t.team.ID = id
	return t
}

func (t *TeamBuilder) City(city string) *TeamBuilder {
	t.team.City = city
	return t
}

func (t *TeamBuilder) Name(name string) *TeamBuilder {
	t.team.Name = name
	return t
}

func (t *TeamBuilder) Abbreviation(abbreviation string) *TeamBuilder {
	t.team.Abbreviation = abbreviation
	return t
}

func (t *TeamBuilder) Conference(conference string) *TeamBuilder {
	t.team.Conference = conference
	return t
}

func (t *TeamBuilder) Division(division string) *TeamBuilder {
	t.team.Division = division
	return t
}

func (t *TeamBuilder) Code(code string) *TeamBuilder {
	t.team.Code = code
	return t
}

func (t *TeamBuilder) Slug(slug string) *TeamBuilder {
	t.team.Slug = slug
	return t
}

func (t *TeamBuilder) InauguralYear(inauguralYear int) *TeamBuilder {
	t.team.InauguralYear = inauguralYear
	return t
}

func (t *TeamBuilder) LatestYear(latestYear int) *TeamBuilder {
	t.team.LatestYear = latestYear
	return t
}
