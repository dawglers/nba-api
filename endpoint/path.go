package endpoint

type Path string

const (
	PlayerIndex    Path = "playerindex"
	TeamInfoCommon Path = "teaminfocommon"

	PlayerGameLogs Path = "playergamelogs"
	ListTeams      Path = "cumestatsteamlist"
	ListPlayers    Path = "cumestatsplayerlist"
)
