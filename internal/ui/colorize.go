package ui

import "github.com/fatih/color"


func Colorize(text string, percent float64) string {
	switch{
	case percent < 50:
		return color.GreenString(text)
	case percent < 80:
		return color.YellowString(text)
	default:
		return color.RedString(text)
	}
}