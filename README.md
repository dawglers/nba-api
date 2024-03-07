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

#### Example 2: Fetch all NBA Centers who were All Stars in the 1995-96 season

```
p := query.Players().
  OnlyAllStar(true).
  Position(player.Center).
  Season(1996)

_ := p.Execute()
players := p.GetPlayers()
```

Output
```
&{121 Patrick Ewing patrick-ewing 1610612752 C 7-0 255 Georgetown Jamaica 1985: Rd 1, Pk 1 33}
&{297 Alonzo Mourning alonzo-mourning 1610612748 C 6-10 261 Georgetown USA 1992: Rd 1, Pk 2 33}
&{87 Dikembe Mutombo dikembe-mutombo 1610612743 C 7-2 260 Georgetown Congo 1991: Rd 1, Pk 4 55}
&{406 Shaquille O'Neal shaquille-oneal 1610612753 C 7-1 325 Louisiana State USA 1992: Rd 1, Pk 1 32}
&{165 Hakeem Olajuwon hakeem-olajuwon 1610612745 C 7-0 255 Houston Nigeria 1984: Rd 1, Pk 1 34}
&{764 David Robinson david-robinson 1610612759 C 7-1 250 Navy USA 1987: Rd 1, Pk 1 50}
```

### Player Game Logs
Query game logs for a player

#### Example 1: Fetch all game logs for Alperen Sengun against the San Antonio Spurs in the 2023 season

```
q := query.PlayerGameLogs(1630578).
  OpponentTeamID(1610612759).
  Season(2023)


_ := q.Execute()
gameLogs := q.GetGameLogs()
```

Output
```
&{1630578 0022200967 HOU vs. SAS 2023-03-05 00:00:00 +0000 UTC 27 6 9 0 0 3 0 3 11 3 1 3 1 2 15 22 47}
&{1630578 0022200959 HOU @ SAS 2023-03-04 00:00:00 +0000 UTC 28 6 9 1 2 3 0 5 5 4 4 1 0 4 16 10 33}
&{1630578 0022200455 HOU vs. SAS 2022-12-19 00:00:00 +0000 UTC 24 8 10 0 0 6 0 1 3 3 0 1 1 3 22 -16 37}
&{1630578 0022200375 HOU @ SAS 2022-12-08 00:00:00 +0000 UTC 29 8 17 0 0 0 0 7 4 2 3 0 0 3 16 -12 29}
```

#### Example 2: Fetch last 10 game logs for Alperen Sengun

```
q := query.PlayerGameLogs(1630578).
  LastNGames(10)

_ := q.Execute()
gameLogs := q.GetGameLogs()
```

Output
```
&{1630578 0022300897 HOU vs. LAC 2024-03-06 00:00:00 +0000 UTC 42 10 19 0 0 3 0 6 13 14 3 2 0 4 23 2 69}
&{1630578 0022300891 HOU vs. SAS 2024-03-05 00:00:00 +0000 UTC 37 19 32 2 3 5 0 6 10 3 2 5 1 3 45 18 84}
&{1630578 0022300870 HOU @ PHX 2024-03-02 00:00:00 +0000 UTC 33 9 21 0 2 3 0 5 5 2 3 2 1 4 21 16 42}
&{1630578 0022300854 HOU @ PHX 2024-02-29 00:00:00 +0000 UTC 23 3 7 0 1 2 0 1 2 4 1 0 0 6 8 -17 16}
&{1630578 0022300841 HOU @ OKC 2024-02-27 00:00:00 +0000 UTC 36 8 13 0 1 7 0 3 8 6 5 2 2 2 23 -11 52}
&{1630578 0022300822 HOU vs. OKC 2024-02-25 00:00:00 +0000 UTC 29 6 12 1 3 6 0 5 7 2 6 1 1 4 19 -4 36}
&{1630578 0022300806 HOU vs. PHX 2024-02-23 00:00:00 +0000 UTC 35 6 15 0 1 5 0 6 6 4 2 0 1 4 17 6 38}
&{1630578 0022300798 HOU @ NOP 2024-02-22 00:00:00 +0000 UTC 27 6 9 1 1 7 0 2 7 2 3 1 0 3 20 -15 33}
&{1630578 0022300783 HOU @ MEM 2024-02-14 00:00:00 +0000 UTC 34 8 16 1 2 2 0 3 2 6 1 1 1 6 19 -6 39}
&{1630578 0022300765 HOU vs. NYK 2024-02-12 00:00:00 +0000 UTC 36 6 10 0 0 6 0 0 3 6 3 1 3 3 18 -6 39}
```