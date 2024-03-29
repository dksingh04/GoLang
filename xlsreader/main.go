package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

var shopMap = make(map[string][]string)

func main() {
	fileName := flag.String("fname", "rc ac new.xlsx", "Passing file Name")
	cellNum := flag.Int("cell", 4, "Pass cell number you want to read, index start from 0")
	outputFNamePath := flag.String("outputDirFname", "output.xlsx", "Pass output File Name with complete path")

	flag.Parse()
	excelFileName := *fileName
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to parse spreadSheet")
	}
	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			cell := sheet.Cell(rowIndex, *cellNum)
			//cells := row.Cells
			style := cell.GetStyle()
			text := strings.TrimLeft(cell.String(), " ")
			chCell := row.AddCell()
			chCell.SetStyle(style)
			//chCell.Cell(rowIndex, len(cells))
			chCell.SetString(text[:1])
			shopMap[text[:1]] = append(shopMap[text[:1]], text)
		}
	}
	xlFile.Save("updated.xlsx")
	// Creating new File with character starting with
	file := xlsx.NewFile()

	newSheet, err := file.AddSheet("output")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in creating output %s", err.Error())
	}

	for key, values := range shopMap {
		row := newSheet.AddRow()
		cell1 := row.AddCell()
		cell1.Value = key
		cell2 := row.AddCell()
		cell2.Value = strings.Join(values, ",")

	}

	err = file.Save(*outputFNamePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create output file %s", err.Error())
	}

}
