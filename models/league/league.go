package league

var idToName = map[ID]string{
	NBA:     "NBA",
	ABA:     "ABA",
	WNBA:    "WNBA",
	GLeague: "GLeague",
}

type ID string

const (
	NBA     ID = "00"
	ABA     ID = "01"
	WNBA    ID = "10"
	GLeague ID = "20"
	Default ID = NBA
)

func (l ID) Name() string {
	return idToName[l]
}
