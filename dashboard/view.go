package dashboard

import (
	"gobadge/svg"
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
	view.Canvas.Open()
	view.DrawHeader()
	view.Draw()
	view.Canvas.Close()
}
