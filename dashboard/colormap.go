package dashboard

func colormap(color string) string {

	var fill string

	red := "#FF2342"
	green := "#46BF50"
	amber := "#F68E00"
	black := "#2B2B2B"
	gray := "#737373"

	switch color {
	case "green":
		fill = green
	case "success":
		fill = green
	case "pass":
		fill = green
	case "red":
		fill = red
	case "amber":
		fill = amber
	case "warn":
		fill = amber
	case "fail":
		fill = red
	case "black":
		fill = black
	case "dark":
		fill = black
	case "gray":
		fill = gray
	default:
		fill = gray
	}

	return fill
}
