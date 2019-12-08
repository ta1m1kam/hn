package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/tealeg/xlsx"
	"os"
	"strconv"
)

func generateSheets() *tview.List {
	result := tview.NewList().ShowSecondaryText(false)
	result.SetBorder(true).SetTitle("Sheets")
	return result
}

func generateTables() *tview.Table {
	tables := tview.NewTable().SetBorders(true)
	tables.SetBorder(true).SetTitle("Content")
	return tables
}

func createExcelInstance(excelFileName string) *xlsx.File {
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	return xlFile
}

func addSheets(sheets *tview.List, tables *tview.Table, excelInstance *xlsx.File, app *tview.Application) {
	for _, sheet := range excelInstance.Sheets {
		tmp := sheet
		sheets.AddItem(tmp.Name, "", 0, func() {
			tables.Clear()

			for i, row := range tmp.Rows {
				tables.SetCellSimple(i, 0, strconv.Itoa(i+1))

				for j, cell := range row.Cells {
					tables.SetCellSimple(i, j+1, cell.String())
				}
			}

			tables.ScrollToBeginning()
			app.SetFocus(tables)
		})
	}
}

func selectSheet(sheets *tview.List, tables *tview.Table, app *tview.Application) {
	tables.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			tables.Clear()
			app.SetFocus(sheets)
		case tcell.KeyEnter:
			tables.SetSelectable(true, false)
		}
	})

	sheets.SetDoneFunc(func() {
		app.Stop()
		os.Exit(0)
	})

	tables.SetSelectedFunc(func(row int, column int) {
		for i := 0; i < tables.GetColumnCount(); i++ {
			tables.GetCell(row, i).SetTextColor(tcell.ColorRed)
		}
		tables.SetSelectable(false, false)
	})
}

func main() {
	app := tview.NewApplication()
	sheets := generateSheets()
	tables := generateTables()
	excel := createExcelInstance("./files/test.xlsx")
	addSheets(sheets, tables, excel, app)
	selectSheet(sheets, tables, app)

	flex := tview.NewFlex().
		AddItem(sheets, 0, 1, true).
		AddItem(tables, 0, 5, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		fmt.Printf("Error running appplication: %s\n", err)
	}
}
