package ui

import "strings"

func ProgressBar(percent float64, width int) string {
    filled := int((percent / 100) * float64(width))
    if filled > width {
        filled = width
    }
    empty := width - filled
    return "[" + strings.Repeat("█", filled) + strings.Repeat("░", empty) + "]"
}


