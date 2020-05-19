package dashboard

import (
	"fmt"
	geo "gobadge/svg/geometry"
	"gobadge/svg/shapes"
	"gobadge/svg/style"
	"strings"
)

func stockBadgeLabel(text string) shapes.Text {
	id := makeTextID(text, "label")
	origin := geo.Coordinate{X: 8, Y: 14}
	return stockText(text, id, origin)
}

func stockBadgeTag(text string) shapes.Text {
	id := makeTextID(text, "tag")
	origin := geo.Coordinate{X: 133, Y: 14}
	return stockText(text, id, origin)
}

func stockBadgeBlock() shapes.Path {
	return shapes.Path{
		ID:       "block",
		Fill:     blockfill,
		FillRule: "nonzero",
		D:        "M3,0 L125,0 L125,0 L125,20 L3,20 C1.34314575,20 2.02906125e-16,18.6568542 0,17 L0,3 C-2.02906125e-16,1.34314575 1.34314575,3.04359188e-16 3,0 Z",
	}
}

func stockBadgeColor(fill string) shapes.Path {
	return shapes.Path{
		ID:       "color",
		Fill:     fill,
		FillRule: "nonzero",
		D:        "M125,0 L227,0 C228.656854,-3.04359188e-16 230,1.34314575 230,3 L230,17 C230,18.6568542 228.656854,20 227,20 L125,20 L125,20 L125,0 Z",
	}
}

func stockBadgeGradientRect() shapes.Rect {
	return shapes.Rect{
		ID:       "gradient",
		Fill:     "url(#texture)",
		FillRule: "nonzero",
		Origin:   geo.Coordinate{X: 0, Y: 0},
		Radius:   geo.Radius{X: 3, Y: 3},
		Frame:    geo.Frame{Width: 230, Height: 20},
	}
}

func stockBadgeGradient(topColor string, topOpacity float64, bottomColor string, bottomOpacity float64) style.LinearGradient {
	return stockLinearGradient("texture", topColor, topOpacity, bottomColor, bottomOpacity)
}

func stockLinearGradient(id string, topColor string, topOpacity float64, bottomColor string, bottomOpacity float64) style.LinearGradient {
	return style.LinearGradient{
		ID:     id,
		Top:    style.Point{X: 50, Y: 0},
		Bottom: style.Point{X: 50, Y: 100},
		Stops: []style.Stop{
			style.Stop{
				Color:   bottomColor,
				Opacity: bottomOpacity,
				Offset:  0,
			},
			style.Stop{
				Color:   topColor,
				Opacity: topOpacity,
				Offset:  100,
			},
		},
	}
}

func stockTextFilter() style.Filter {
	return style.Filter{ID: "shadow"}
}

func stockText(text string, id string, origin geo.Coordinate) shapes.Text {
	return shapes.Text{
		ID:    id,
		Value: text,
		Font: shapes.Font{
			Family: "Verdana",
			Size:   10,
			Weight: "normal",
		},
		Origin: origin,
	}
}

func makeTextID(text string, element string) string {
	return fmt.Sprintf("%s-%s-text", strings.Replace(text, " ", "-", 1), element)
}
