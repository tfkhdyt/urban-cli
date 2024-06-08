package lib

import (
	"fmt"
	"strings"
)

func formatUsenameAndDate(input string) string {
	// Split the input string to extract the date part
	parts := strings.SplitN(input, " ", -1)
	dateStart := findDateStartIndex(parts)
	username := strings.TrimSpace(strings.Join(parts[:dateStart], " "))
	date := strings.Join(parts[dateStart:], " ")

	// Format the output string
	formattedOutput := fmt.Sprintf("%s (%s)", username, date)

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
