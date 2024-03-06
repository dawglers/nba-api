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