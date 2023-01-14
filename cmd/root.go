package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/davidpalves/csviz/reader"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringP("filepath", "f", "", "Path to the CSV file to be read")
}

var rootCmd = &cobra.Command{
	Use:   "csviz",
	Short: "Reads CSV directly from terminal",
	Run: func(cmd *cobra.Command, args []string) {
		filePathFlag, _ := cmd.Flags().GetString("filepath")

		if strings.TrimSpace(filePathFlag) == "" && len(args) == 0 {
			cmd.Usage()
			return
		}

		if filePathFlag != "" {
			reader.RenderTable(filePathFlag)
		} else if len(args) > 0 {
			filePathArg := args[0]
			reader.RenderTable(filePathArg)
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
