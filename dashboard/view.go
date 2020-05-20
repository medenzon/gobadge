package dashboard

import (
	"gobadge/svg"
	"math"
)

type View struct {
	Canvas  *svg.Canvas
	Header  Header  `json:"header"`
	Columns int     `json:"columns"`
	Badges  []Badge `json:"badge"`
}

type Header struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (view *View) Draw() {

	count := len(view.Badges)
	split := float64(count) / float64(view.Columns)
	rows, cols := int(math.Ceil(split)), view.Columns

	width := float64(((badgeW + 20) * (view.Columns - 1)) + badgeW)
	height := float64(((badgeH + vpad) * rows) - vpad + voff)

	view.Canvas.Open(width, height)
	view.DrawHeader()
	view.draw(count, rows, cols)
	view.Canvas.Close()
}
