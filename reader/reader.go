package reader

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/davidpalves/csviz/pager"
	"github.com/pterm/pterm"
)

func FormatTable(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	headerStyle := pterm.NewStyle(pterm.FgRed, pterm.Bold)
	separatorStyle := pterm.NewStyle(pterm.FgGray, pterm.Bold)

	return pterm.DefaultTable.
		WithHeaderStyle(headerStyle).
		WithSeparatorStyle(separatorStyle).
		WithRowSeparatorStyle(separatorStyle).
		WithCSVReader(csvReader).
		WithHasHeader().
		WithBoxed().
		WithHeaderRowSeparator("=").
		WithRowSeparator("-").
		Srender()
}

func RenderTable(filepath string) {
	table, err := FormatTable(filepath)
	if err != nil {
		log.Fatal("Error processing the CSV file.")
	}

	pager.Pager(table)
}
