package endpoint

type Path string

const (
	PlayerIndex        Path = "playerindex"
	TeamInfoCommon     Path = "teaminfocommon"
	CumeStatsTeamList  Path = "cumestatsteamlist"
	CumeStatsTeamGames Path = "cumestatsteamgames"
	Scoreboard         Path = "scoreboardv2"

	PlayerGameLogs Path = "playergamelogs"
	ListPlayers    Path = "cumestatsplayerlist"
)
