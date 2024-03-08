package team

type Team struct {
	ID   int
	Name string
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

func (t *TeamBuilder) Name(name string) *TeamBuilder {
	t.team.Name = name
	return t
}
