package pitaya

import (
	"strings"
)

type DataFrame struct {
	ColIdx   int
	RowIdx   int
	Value    string
	RawValue string
	Sheet    string
}

type DataTable struct {
	data   map[string][][]string
	dfs    map[string][]*DataFrame
	sheets []string
	option *Options
}

func (dt *DataTable) getDataFrames(sheetName string) []*DataFrame {
	if val, ok := dt.dfs[sheetName]; ok {
		return val
	}
	return dt.dfs[dt.sheets[0]]
}

func (dt *DataTable) getSheets() []string {
	if dt.option.AllSheet {
		return dt.sheets
	}
	return []string{dt.sheets[0]}
}

func (dt *DataTable) GetSheets() []string {
	return dt.sheets
}

func (dt *DataTable) Get(str string) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.GetSheet(sheet, str)...)
	}
	return
}

func (dt *DataTable) GetSheet(sheet, str string) (dfs []*DataFrame) {
	for _, df := range dt.getDataFrames(sheet) {
		if df.Value == str {
			dfs = append(dfs, df)
		}
	}
	return dfs
}

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

func (dt *DataTable) SheetLast(sheet, str string) *DataFrame {
	dfs := dt.GetSheet(sheet, str)
	length := len(dfs)
	if length > 0 {
		return dfs[length-1]
	}
	return nil
}

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

func (dt *DataTable) SheetFirst(sheet, str string) *DataFrame {
	dfs := dt.GetSheet(sheet, str)
	if len(dfs) > 0 {
		return dfs[0]
	}
	return nil
}

func (dt *DataTable) Contains(str string) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.SheetContains(sheet, str)...)
	}
	return
}

func (dt *DataTable) SheetContains(sheet, str string) (dfs []*DataFrame) {
	for _, df := range dt.getDataFrames(sheet) {
		if strings.Index(df.Value, str) >= 0 {
			dfs = append(dfs, df)
		}
	}
	return
}

func (dt *DataTable) GetColByIndex(index int) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.GetSheetColByIndex(sheet, index)...)
	}
	return
}

func (dt *DataTable) GetSheetColByIndex(sheet string, index int) (dfs []*DataFrame) {
	for _, df := range dt.dfs[sheet] {
		if df.ColIdx == index {
			dfs = append(dfs, df)
		}
	}
	return
}

func (dt *DataTable) HasRow(strs ...string) bool {
	for _, sheet := range dt.getSheets() {
		if list, ok := dt.data[sheet]; ok {
			for _, rows := range list {
				if strings.Join(rows, ",") ==
					strings.Join(strs, ",") {
					return true
				}
			}
		}
	}
	return false
}

func (dt *DataTable) GetRowByIndex(index int) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		dfs = append(dfs, dt.GetSheetRowByIndex(sheet, index)...)
	}
	return
}

func (dt *DataTable) GetSheetRowByIndex(sheet string, index int) (dfs []*DataFrame) {
	for _, df := range dt.dfs[sheet] {
		if df.RowIdx == index {
			dfs = append(dfs, df)
		}
	}
	return
}

func (dt *DataTable) GetCellByIndex(rowIdx, colIdx int) (dfs []*DataFrame) {
	for _, sheet := range dt.getSheets() {
		df := dt.GetSheetCellByIndex(sheet, rowIdx, colIdx)
		if df == nil {
			continue
		}
		dfs = append(dfs, df)
	}
	return
}

func (dt *DataTable) GetSheetCellByIndex(sheet string, rowIdx, colIdx int) *DataFrame {
	for _, df := range dt.getDataFrames(sheet) {
		if df.RowIdx == rowIdx && df.ColIdx == colIdx {
			return df
		}
	}
	return nil
}

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

func (dt *DataTable) GetSheetCell(sheet, row, col string) *DataFrame {
	rowIdx := -1
	colIdx := -1
	for _, df := range dt.getDataFrames(sheet) {
		if rowIdx == -1 && df.Value == row {
			rowIdx = df.RowIdx
		}
		if colIdx == -1 && df.Value == col {
			colIdx = df.ColIdx
		}
		if df.RowIdx == rowIdx && df.ColIdx == colIdx {
			return df
		}
	}
	return nil
}
