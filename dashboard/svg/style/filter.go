package style

import (
	"strings"
)

type Filter struct {
	ID string
}


func (filter Filter) Vectorize() string {
	return strings.Join([]string{`<filter x="-0.3%" y="-3.1%" width="101.4%" height="112.5%" filterUnits="objectBoundingBox" id="` + filter.ID + `">`,
	`<feOffset dx="0.75" dy="0.75" in="SourceAlpha" result="shadowOffsetOuter1"></feOffset>`,
	`<feColorMatrix values="0 0 0 0 0   0 0 0 0 0   0 0 0 0 0  0 0 0 0.5 0" type="matrix" in="shadowOffsetOuter1">`,
	`</feColorMatrix>`,
	`</filter>`}, "\n")
}