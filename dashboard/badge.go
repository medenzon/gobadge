package dashboard

import (
	"github.com/medenzon/gobadge/dashboard/svg"
	geo "github.com/medenzon/gobadge/dashboard/svg/geometry"
)

type Badge struct {
	Label string `json:"label"`
	Tag   string `json:"tag"`
	Color string `json:"color"`
}

func (badge Badge) Draw(canvas *svg.Canvas) {
	canvas.Open()
	origin := geo.Coordinate{X: 0, Y: 0}
	badge.draw(canvas, origin)
	canvas.Close()
}
