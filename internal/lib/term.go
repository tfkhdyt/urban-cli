package lib

import (
	"fmt"
	"os"
	"slices"

	"golang.org/x/term"
)

func PrintResult(result []Result, reverse bool) {
	if reverse {
		slices.Reverse(result)
	}

	for _, r := range result {
		number := fmt.Sprintf("#%d", r.ID)

		fmt.Println(
			number,
			r.Word,
			getHorizontalLine(
				getTermSize()-len(r.Contributor)-len(r.Word)-len(number)-3,
			),
			r.Contributor,
		)
		fmt.Printf("%s\n\n", r.Meaning)
		fmt.Printf("%s\n\n", r.Example)
	}
	fmt.Println(getHorizontalLine(getTermSize()))
}

func getTermSize() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		os.Exit(1)
	}

	return width
}

func getHorizontalLine(width int) string {
	line := ""
	for i := 0; i < width; i++ {
		line += "─"
	}
	return line
}
