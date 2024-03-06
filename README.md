# nba-api

## Queries

### Players
Query all players

#### Example 1: Fetch all WNBA Guards playing for the Las Vegas Aces

```
p := query.Players().
  LeagueID(league.WNBA).
  TeamID(1611661319).
  Position(player.Guard)

_ := p.Execute()
players := p.GetPlayers()
```

Output
```
&{202641 Sydney Colson sydney-colson 1611661319 G 5-8 140 Texas A&M USA 2011: Rd 2, Pk 16 51}
&{203833 Chelsea Gray chelsea-gray 1611661319 G 5-11 170 Duke USA 2014: Rd 1, Pk 11 12}
&{1628276 Kelsey Plum kelsey-plum 1611661319 G 5-8 145 Washington USA 2017: Rd 1, Pk 1 10}
&{203029 Riquna Williams riquna-williams 1611661319 G 5-7 165 Miami USA 2012: Rd 2, Pk 17 2}
&{1629498 Jackie Young jackie-young 1611661319 G 6-0 165 Notre Dame USA 2019: Rd 1, Pk 1 0}
```

#### Example 2: Fetch all NBA All Stars playing for the Boston Celtics

```
p := query.Players().
	TeamID(1610612738).
	OnlyAllStar(true)

_ := p.Execute()
players := p.GetPlayers()
```

Output
```
&{1627759 Jaylen Brown jaylen-brown 1610612738  6-6 223 California USA 2016: Rd 1, Pk 3 7}
&{1628369 Jayson Tatum jayson-tatum 1610612738  6-8 210 Duke USA 2017: Rd 1, Pk 3 0}
```