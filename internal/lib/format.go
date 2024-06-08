package lib

import (
	"strings"

	"github.com/fatih/color"
)

func formatUsernameAndDate(input string) string {
	// Split the input string to extract the date part
	parts := strings.Split(input, " ")
	dateStart := findDateStartIndex(parts)
	username := color.BlueString(
		strings.TrimSpace(
			strings.TrimPrefix(strings.Join(parts[:dateStart], " "), "by "),
		),
	)
	date := strings.Join(parts[dateStart:], " ")

	// Format the output string
	formattedOutput := color.New(color.Reset, color.Bold).
		Sprintf("by %s (%s)", username, color.New(color.Bold).Sprint(date))

	return formattedOutput
}

func findDateStartIndex(parts []string) int {
	months := map[string]struct{}{
		"January": {}, "February": {}, "March": {}, "April": {}, "May": {}, "June": {},
		"July": {}, "August": {}, "September": {}, "October": {}, "November": {}, "December": {},
	}
	for i, part := range parts {
		if _, exists := months[part]; exists {
			return i
		}
	}
	return len(
		parts,
	)
}
