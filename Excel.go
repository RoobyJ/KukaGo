package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"time"
)

func (data *dataTree) createExcel(path string, fileName string) {
	wb := excelize.NewFile()
	data.startTime = time.Now().UnixMilli()
	// Create a new sheet for each type
	for _, dataType := range data.types {
		currentSheet := wb.NewSheet(dataType.name)
		wb.SetActiveSheet(currentSheet)

		wb = data.configureSheet(dataType, wb)
		//// Set value of a cell.
		wb = data.passData(dataType, wb)
		//checkError(wb.SetCellValue(dataType.name, "B2", 100))
		// Set active sheet of the workbook.

		// Save spreadsheet by the given path.
		wb.DeleteSheet("Sheet1")
		data.endTime = time.Now().UnixMilli()

		if err := wb.SaveAs(fmt.Sprintf("%s\\config_%s.xlsx", path, fileName)); err != nil {
			fmt.Println(err)
		}
	}

}

func (data *dataTree) passData(dataType datatype, wb *excelize.File) *excelize.File {
	for _, typ := range data.types {
		if dataType.name == typ.name {
			for i := 0; i <= 127; i++ { // 127 because each data type has 128 rows and the for loop starts from 0
				for ii, num := range typ.typeValues[i] {
					switch ii {
					case 0:
						str := fmt.Sprintf("%s%d", string(rune(ii+65)), i+2)
						checkError(wb.SetCellValue(typ.name, str, i))
						str = fmt.Sprintf("%s%d", string(rune(ii+67)), i+2)
						checkError(wb.SetCellValue(typ.name, str, num))
					case 1:
						str := fmt.Sprintf("%s%d", string(rune(ii+65)), i+2)
						checkError(wb.SetCellValue(typ.name, str, typ.typeNames[i]))
						str = fmt.Sprintf("%s%d", string(rune(ii+67)), i+2)
						checkError(wb.SetCellValue(typ.name, str, num))
					default:
						str := fmt.Sprintf("%s%d", string(rune(ii+67)), i+2)
						checkError(wb.SetCellValue(typ.name, str, num))
					}

				}
			}
		}
	}
	return wb
}

// Set width of columns in specified sheet and pass column names
func (data *dataTree) configureSheet(currentType datatype, wb *excelize.File) *excelize.File {
	titles := map[string][]string{
		"bases": {"base_data", "base_name", "X", "Y", "Z", "A", "B", "C"},
		"tools": {"tool_data", "tool_name", "X", "Y", "Z", "A", "B", "C"},
		"loads": {"load_data", "load_name", "M", "X", "Y", "Z", "A", "B", "C", "Ix", "Iy", "Iz"},
	}
	checkError(wb.SetColWidth(currentType.name, "A", "B", 10))
	checkError(wb.SetColWidth(currentType.name, "B", "L", 20))
	for dType, dataList := range titles {
		if dType == currentType.name {
			for i, title := range dataList {
				str := []string{string(rune(i + 65)), strconv.Itoa(1)}
				checkError(wb.SetCellValue(currentType.name, strings.Join(str, ""), title))
			}
		}
	}
	return wb
}
