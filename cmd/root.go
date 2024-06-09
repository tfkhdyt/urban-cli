/*
Copyright © 2024 Taufik Hidayat <tfkhdyt@proton.me>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/tfkhdyt/urban-cli/internal/lib"
)

var (
	maxResults int
	reverse    bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "urban-cli [keyword]",
	Short:   "Blazingly fast command line interface for Urban Dictionary",
	Long:    "Blazingly fast command line interface for Urban Dictionary",
	Version: "0.2.3",
	Args:    cobra.ExactArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		keyword := args[0]

		result, err := lib.Scrape(keyword, maxResults)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		num := maxResults
		if num > len(result) {
			num = len(result)
		}

		lib.PrintResult(result[:num], reverse)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.urban-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().
		IntVarP(&maxResults, "max", "m", 7, "maximum number of results to show")

	rootCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "reverse results")
}
