package dashboard

import "github.com/medenzon/gobadge/dashboard/svg"

type View struct {
	Canvas  *svg.Canvas
	Header  Header  `json:"header"`
	Columns int     `json:"columns"`
	Badges  []Badge `json:"badge"`
}

type Header struct {
	Title  string `json:"title"`
	Detail string `json:"title"`
}

func (view *View) Draw() {
	view.Canvas.Open()
	view.DrawHeader()
	view.draw()
	view.Canvas.Close()
}
