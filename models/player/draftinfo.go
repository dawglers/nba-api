package player

import "fmt"

type DraftInfo struct {
	Year  int
	Round int
	Pick  int
}

type DraftInfoBuilder struct {
	draftInfo DraftInfo
}

func (d *DraftInfoBuilder) Build() DraftInfo {
	return d.draftInfo
}

func (d *DraftInfoBuilder) Year(draftYear int) *DraftInfoBuilder {
	d.draftInfo.Year = draftYear
	return d
}

func (d *DraftInfoBuilder) Round(draftRound int) *DraftInfoBuilder {
	d.draftInfo.Round = draftRound
	return d
}

func (d *DraftInfoBuilder) Pick(draftPick int) *DraftInfoBuilder {
	d.draftInfo.Pick = draftPick
	return d
}

func (d DraftInfo) Undrafted() bool {
	if d.Year == 0 || d.Round == 0 || d.Pick == 0 {
		return true
	}
	return false
}

func (d DraftInfo) String() string {
	if d.Year == 0 || d.Round == 0 || d.Pick == 0 {
		return "Undrafted"
	}
	return fmt.Sprintf("%d: Rd %d, Pk %d", d.Year, d.Round, d.Pick)
}
