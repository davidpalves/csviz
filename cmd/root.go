package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/davidpalves/csviz/reader"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().StringP("filepath", "f", "", "Path to the CSV file to be read")
}

var rootCmd = &cobra.Command{
	Use:   "csviz",
	Short: "Reads CSV directly from terminal",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("filepath")

		if strings.TrimSpace(filePath) == "" {
			cmd.Usage()
			return
		}

		if filePath != "" {
			reader.RenderTable(filePath)
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
