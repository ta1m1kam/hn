package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/tealeg/xlsx"
)

func main() {
	app := tview.NewApplication()

	// Generate Sheet List Instance
	generateSheets := func() *tview.List {
		result := tview.NewList().ShowSecondaryText(false)
		result.SetBorder(true).SetTitle("Sheets")

		return result
	}
	sheets := generateSheets()

	// Generate Column Table Instance
	tables := tview.NewTable().SetBorders(true)
	tables.SetBorder(true).SetTitle("Contents")

	// Create Excel instance
	createExcelInstance := func(excelFileName string) *xlsx.File {
		xlFile, err := xlsx.OpenFile(excelFileName)
		if err != nil {
			panic(err)
		}

		return xlFile
	}
	excel := createExcelInstance("./files/test.xlsx")

	// Add Sheets to Sheet List
	addSheets := func(sheets *tview.List, excelInstance *xlsx.File) {
		for _, sheet := range excelInstance.Sheets {
			tmp := sheet

			sheets.AddItem(tmp.Name, "", 0, func() {
				// Clear the table
				tables.Clear()

				for i, row := range tmp.Rows {
					// Add row number:
					tables.SetCellSimple(i, 0, strconv.Itoa(i+1))

					for j, cell := range row.Cells {
						tables.SetCellSimple(i, j+1, cell.String())
					}
				}

				tables.ScrollToBeginning()
				app.SetFocus(tables)
			})

			sheets.SetDoneFunc(func() {
				app.Stop()
				os.Exit(0)
			})

			tables.SetDoneFunc(func(key tcell.Key) {
				switch key {
				case tcell.KeyEscape:
					tables.Clear()
					app.SetFocus(sheets)
				case tcell.KeyEnter:
					// Press Enter to select the rows
					tables.SetSelectable(true, false)
				}
			})

			tables.SetSelectedFunc(func(row int, column int) {
				for i := 0; i < tables.GetColumnCount(); i++ {
					tables.GetCell(row, i).SetTextColor(tcell.ColorRed)
				}
				tables.SetSelectable(false, false)
			})
		}

	}
	addSheets(sheets, excel)

	flex := tview.NewFlex().
		AddItem(sheets, 0, 1, true).
		AddItem(tables, 0, 5, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
	}
}
