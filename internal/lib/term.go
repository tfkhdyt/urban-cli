package lib

import (
	"fmt"
	"os"
	"regexp"
	"slices"

	"github.com/fatih/color"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/term"
)

func PrintResult(result []Result, reverse bool) {
	if reverse {
		slices.Reverse(result)
	}

	for idx, r := range result {
		highlight := color.New(color.FgBlue, color.Bold, color.Underline).
			SprintFunc()

		for _, word := range r.Autolinks {
			regex := regexp.MustCompile(`\b` + regexp.QuoteMeta(word) + `\b`)
			r.Meaning = regex.ReplaceAllStringFunc(
				r.Meaning,
				func(matched string) string {
					return highlight(matched)
				},
			)

			r.Example = regex.ReplaceAllStringFunc(
				r.Example,
				func(matched string) string {
					return highlight(matched)
				},
			)
		}

		number := fmt.Sprintf("#%d", r.ID)

		// header
		color.New(color.Bold).Println(
			number,
			color.GreenString(r.Word),
			getHorizontalLine(
				getTermSize()-len(r.Contributor)-len(r.Word)-len(number)+28,
			),
			r.Contributor,
		)

		// content
		fmt.Println(
			wordwrap.String(
				fmt.Sprintf("%s\n", r.Meaning),
				getTermSize(),
			),
		)

		fmt.Println("Example:")
		fmt.Println(
			indent.String(wordwrap.String(r.Example, getTermSize()-4), 4),
		)
		if idx != len(result)-1 {
			fmt.Printf("\n\n\n")
		}
	}
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
