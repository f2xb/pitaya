package pitaya

import (
	"strings"
)

type DataFrame struct {
	Col, Row int
	Value    string
	RawValue string
	Sheet    string
}

type DataTable struct {
	data   map[string][][]string
	dfs    map[string][]*DataFrame
	sheets []string
	option *Options
	Ext    string
}

func (dt *DataTable) getDataFrames(sheetName string) []*DataFrame {
	if val, ok := dt.dfs[sheetName]; ok {
		return val
	}
	return dt.dfs[dt.sheets[0]]
}

func (dt *DataTable) getData(sheetName string) [][]string {
	if val, ok := dt.data[sheetName]; ok {
		return val
	}
	return dt.data[dt.sheets[0]]
}

func (dt *DataTable) getSheets() []string {
	if dt.option.AllSheet {
		return dt.sheets
	}
	return []string{dt.sheets[0]}
}

// MaxRow get max row, if option AllSheet is true, get all sheet row
// default sheet[0]
// Example: MaxRow()
func (dt *DataTable) MaxRow() (total int) {
	for _, sheet := range dt.getSheets() {
		total += dt.SheetMaxRow(sheet)
	}
	return
}

// SheetMaxRow get sheet max row
// Example: SheetMaxRow("Sheet1")
func (dt *DataTable) SheetMaxRow(sheet string) (total int) {
	return len(dt.getData(sheet))
}

// GetSheets get all sheet name
// Example: GetSheets()
func (dt *DataTable) GetSheets() []string {
	return dt.sheets
}

// GetSheetByIndex get sheet name by index
// Example: GetSheetByIndex(1)
func (dt *DataTable) GetSheetByIndex(idx int) string {
	for i, sheet := range dt.sheets {
		if i == idx-1 {
			return sheet
		}
	}
	return ""
}

// Get get df, str eq df.value
// default sheet[0]
// Example: Get("AA")
func (dt *DataTable) Get(str string) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.GetSheet(sheet, str)...)
	}
	return
}

// GetSheet get sheet df, str eq df.value
// Example: GetSheet("Sheet1", "AA")
func (dt *DataTable) GetSheet(sheet, str string) (dfs []*DataFrame) {
	for _, df := range dt.getDataFrames(sheet) {
		if df.Value == str {
			dfs = append(dfs, df)
		}
	}
	return dfs
}

// Last get last df, str eq df.value
// default sheet[0]
// Example: Last("AA")
func (dt *DataTable) Last(str string) *DataFrame {
	var dfs []*DataFrame
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.GetSheet(sheet, str)...)
	}
	length := len(dfs)
	if length > 0 {
		return dfs[length-1]
	}
	return nil
}

// SheetLast get sheet last df, str eq df.value
// Example: SheetLast("Sheet1", "AA")
func (dt *DataTable) SheetLast(sheet, str string) *DataFrame {
	dfs := dt.GetSheet(sheet, str)
	length := len(dfs)
	if length > 0 {
		return dfs[length-1]
	}
	return nil
}

// First get first df, str eq df.value
// default sheet[0]
// Example: First("AA")
func (dt *DataTable) First(str string) *DataFrame {
	for _, sheet := range dt.getSheets() {
		for _, df := range dt.GetSheet(sheet, str) {
			if df.Value == str {
				return df
			}
		}
	}
	return nil
}

// SheetFirst get sheet first df, str eq df.value
// Example: SheetFirst("Sheet1", "AA")
func (dt *DataTable) SheetFirst(sheet, str string) *DataFrame {
	dfs := dt.GetSheet(sheet, str)
	if len(dfs) > 0 {
		return dfs[0]
	}
	return nil
}

// Contains get contains str df
// default sheet[0]
// Example: Contains("AA")
func (dt *DataTable) Contains(str string) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.SheetContains(sheet, str)...)
	}
	return
}

// SheetContains get sheet contains str df
// Example: SheetContains("Sheet1", "AA")
func (dt *DataTable) SheetContains(sheet, str string) (dfs []*DataFrame) {
	for _, df := range dt.getDataFrames(sheet) {
		if strings.Index(df.Value, str) >= 0 {
			dfs = append(dfs, df)
		}
	}
	return
}

// GetCol get col
// default sheet[0]
// Example: GetCol(1, 2)
func (dt *DataTable) GetCol(cols ...int) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.GetSheetCol(sheet, cols...)...)
	}
	return
}

// GetSheetCol get sheet col
// default all col
// Example: GetSheetCol("Sheet1", 1, 2)
func (dt *DataTable) GetSheetCol(sheet string, cols ...int) (dfs []*DataFrame) {
	if len(cols) == 0 {
		return dt.dfs[sheet]
	}
	for _, df := range dt.dfs[sheet] {
		for _, col := range cols {
			if df.Col == col {
				dfs = append(dfs, df)
			}
		}
	}
	return
}

// GetRow get row
// default sheet[0]
// Example: GetRow(1, 2)
func (dt *DataTable) GetRow(rows ...int) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.GetSheetRow(sheet, rows...)...)
	}
	return
}

// GetSheetRow get sheet row
// default all row
// Example: GetSheetRow("Sheet1", 1, 2)
func (dt *DataTable) GetSheetRow(sheet string, rows ...int) (dfs []*DataFrame) {
	if len(rows) == 0 {
		return dt.dfs[sheet]
	}
	for _, df := range dt.dfs[sheet] {
		for _, row := range rows {
			if df.Row == row {
				dfs = append(dfs, df)
			}
		}
	}
	return
}

// RowContains get contains str row
// default sheet[0]
// Example: RowContains("AAA")
func (dt *DataTable) RowContains(strs ...string) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.SheetRowContains(sheet, strs...)...)
	}
	return
}

// SheetRowContains get sheet contains str row
// Example:
// 	SheetRowContains("Sheet1", "AAA")
// 	SheetRowContains("Sheet1", "AAA", "BBB")
func (dt *DataTable) SheetRowContains(sheet string, strs ...string) (dfs []*DataFrame) {
	if len(strs) == 0 {
		return
	}
	row := -1
	for _, df := range dt.getDataFrames(sheet) {
		for _, str := range strs {
			if str == "" {
				continue
			}
			if row != -1 && df.Row != row {
				continue
			}
			if strings.Contains(df.Value, str) {
				if row == -1 {
					row = df.Row
				}
				dfs = append(dfs, df)
			}

		}
	}
	return
}

// RowContainsByRow get contains str row by row
// default sheet[0]
// Example:
// 	RowContainsByRow(1)
// 	RowContainsByRow(1, "AAA")
func (dt *DataTable) RowContainsByRow(row int, strs ...string) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.SheetRowContainsByRow(sheet, row, strs...)...)
	}
	return
}

// SheetRowContainsByRow get sheet contains str row by row
// Example:
// 	SheetRowContainsByRow("Sheet1", 1)
// 	SheetRowContainsByRow("Sheet1", 1, "AA")
func (dt *DataTable) SheetRowContainsByRow(sheet string, row int, strs ...string) (dfs []*DataFrame) {
	list := dt.getDataFrames(sheet)
	if len(strs) == 0 {
		for _, df := range list {
			if df.Row != row {
				continue
			}
			dfs = append(dfs, df)
		}
		return
	}
	for _, df := range list {
		for _, str := range strs {
			if str == "" {
				continue
			}
			if df.Row != row {
				continue
			}
			if strings.Index(df.Value, str) > 0 {
				dfs = append(dfs, df)
			}
		}
	}
	return
}

// GetRowByRow get row by row or cols
// default: sheet[0]
// Example:
// 	GetRowByRow(1)
// 	GetRowByRow(1, 1, 2, 3)
func (dt *DataTable) GetRowByRow(index int, cols ...int) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.GetSheetRowByRow(sheet, index, cols...)...)
	}
	return
}

// GetSheetRowByRow get sheet row by row or cols
// Example:
// 	GetSheetRowByRow("Sheet1", 1)
// 	GetSheetRowByRow("Sheet1", 1, 1, 2, 3)
func (dt *DataTable) GetSheetRowByRow(sheet string, index int, cols ...int) (dfs []*DataFrame) {
	index -= 1
	col := len(cols) > 0
	for _, df := range dt.dfs[sheet] {
		if df.Row == index {
			if col {
				for _, idx := range cols {
					if df.Col == idx {
						dfs = append(dfs, df)
					}
				}
			} else {
				dfs = append(dfs, df)
			}
		}
	}
	return
}

// GetCellByIndex get cell by row, col
// default sheet[0]
// Example: GetCellByIndex(1, 1)
func (dt *DataTable) GetCellByIndex(row, col int) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		df := dt.GetSheetCellByIndex(sheet, row, col)
		if df == nil {
			continue
		}
		dfs = append(dfs, df)
	}
	return
}

// GetSheetCellByIndex get sheet cell by row, col
// Example: GetSheetCellByIndex("Sheet1", 1, 1)
func (dt *DataTable) GetSheetCellByIndex(sheet string, row, col int) *DataFrame {
	for _, df := range dt.getDataFrames(sheet) {
		if df.Row == row && df.Col == col {
			return df
		}
	}
	return nil
}

// GetCell get cell by row str, col str
// default sheet[0]
// Example: GetCell("A001", "001")
func (dt *DataTable) GetCell(row, col string) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		df := dt.GetSheetCell(sheet, row, col)
		if df == nil {
			continue
		}
		dfs = append(dfs, df)
	}
	return
}

// GetSheetCell get sheet cell by row str, col str
// Example: GetSheetCell("Sheet1", "A001", "001")
func (dt *DataTable) GetSheetCell(sheet, row, col string) *DataFrame {
	rowIdx := -1
	colIdx := -1
	for _, df := range dt.getDataFrames(sheet) {
		if rowIdx == -1 && df.Value == row {
			rowIdx = df.Row
		}
		if colIdx == -1 && df.Value == col {
			colIdx = df.Col
		}
		if df.Row == rowIdx && df.Col == colIdx {
			return df
		}
	}
	return nil
}
