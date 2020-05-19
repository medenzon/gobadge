package dashboard

import (
	"fmt"

	geo "gobadge/svg/geometry"
	"gobadge/svg/shapes"
	"gobadge/svg/style"
)

func (view View) DrawHeader() error {

	gradient := stockLinearGradient("header-texture", "#000000", 0.1, "#BBBBBB", 0.1)
	width := float64(((badgeW + 20) * (view.Columns - 1)) + badgeW)
	height := float64(headH)

	fillRect := shapes.Rect{
		ID:       "header-fill",
		Fill:     blockfill,
		FillRule: "nonzero",
		Origin:   geo.Coordinate{X: 0, Y: 0},
		Radius:   geo.Radius{X: 3, Y: 3},
		Frame:    geo.Frame{Width: width, Height: height},
	}

	gradientRect := shapes.Rect{
		ID:       "header-gradient",
		Fill:     "url(#header-texture)",
		FillRule: "nonzero",
		Origin:   geo.Coordinate{X: 0, Y: 0},
		Radius:   geo.Radius{X: 3, Y: 3},
		Frame:    geo.Frame{Width: width, Height: height},
	}

	titleText := shapes.Text{
		ID:    "title",
		Value: view.Header.Title,
		Font: shapes.Font{
			Family: "Verdana",
			Size:   10,
			Weight: "normal",
		},
		Origin: geo.Coordinate{X: 14, Y: 24},
	}

	filter := stockTextFilter()

	titleStyle := style.Text{Filter: filter, Text: titleText}

	output := []string{
		gradient.Vectorize(),
		fillRect.Vectorize(),
		gradientRect.Vectorize(),
		titleText.Vectorize(),
		titleStyle.Vectorize(),
	}

	_, err := fmt.Fprint(view.Canvas.Writer, output)

	return err
}
