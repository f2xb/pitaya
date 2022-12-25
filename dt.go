package pitaya

import (
	"strings"
)

type DataTable struct {
	dfs    map[string][]*Row
	sheets []string
	option *Options

	Ext string
}

func (dt *DataTable) getDataFrames(sheets ...string) []*Row {
	sheet := dt.sheets[0]
	for _, s := range sheets {
		sheet = s
	}
	return dt.dfs[sheet]
}

// Foreach for range all df
// default sheet[0]
func (dt *DataTable) Foreach(fn dfFunc, sheets ...string) {
	for _, row := range dt.getDataFrames(sheets...) {
		row.Foreach(fn)
	}
}

// ForeachRow for range all row
// default sheet[0]
func (dt *DataTable) ForeachRow(fn rowFunc, sheets ...string) {
	for idx, row := range dt.getDataFrames(sheets...) {
		if ok := fn(idx+1, row); ok {
			return
		}
	}
}

// MaxRow Get the maximum number of rows
// default sheet[0]
// Example:
// MaxRow() // return: 4
// MaxRow("Sheet2") // return: 1
func (dt *DataTable) MaxRow(sheets ...string) (total int) {
	return len(dt.getDataFrames(sheets...))
}

// GetSheets Get all sheet name
// Example:
// GetSheets() // return: [Sheet1]
func (dt *DataTable) GetSheets() []string {
	return dt.sheets
}

// GetSheetByIndex Get sheet name by index
// Example:
// GetSheetByIndex(1) // return: Sheet1
func (dt *DataTable) GetSheetByIndex(idx int) string {
	for i, sheet := range dt.sheets {
		if i == idx-1 {
			return sheet
		}
	}
	return ""
}

// Last get df, str eq df.value
// default sheet[0]
// Example:
// -----------------
// A1	B1	C1
// A2	aa	bb
// A3	aa	bb
// -----------------
// Last("aa") // return: {Sheet:Sheet1,Row:3,Col:2,Value:aa,RawValue:aa}
// Last("aa", "Sheet1")
func (dt *DataTable) Last(str string, sheets ...string) *DataFrame {
	var obj *DataFrame
	dt.Foreach(func(df *DataFrame) bool {
		if df.Value == str {
			obj = df
		}
		return false
	}, sheets...)
	return obj
}

// First get df, str eq df.value
// default sheet[0]
// Example:
// -----------------
// A1	B1	C1
// A2	aa	bb
// A3	aa	bb
// -----------------
// First("aa") // return: {Sheet:Sheet1,Row:2,Col:2,Value:aa,RawValue:aa}
// First("aa", "Sheet1")
func (dt *DataTable) First(str string, sheets ...string) *DataFrame {
	var obj *DataFrame
	dt.ForeachRow(func(idx int, row *Row) bool {
		obj = row.First(str)
		if obj != nil {
			return true
		}
		return false
	}, sheets...)
	return obj
}

// Rows Get rows
// default sheet[0]
// Example:
// -----------------
// A1	B1	C1
// A2	B2	C2
// -----------------
// Rows()
// Rows("Sheet1")
func (dt *DataTable) Rows(sheets ...string) (list []*Row) {
	return dt.getDataFrames(sheets...)
}

// Get Get df.Value eq str
// default sheet[0]
// Example:
// -----------------
// A1	B1	C1
// A2	aa	bb
// A3	aa	bb
// -----------------
// Get("A1") // return: [A1]
// Get("A1", "Sheet1") // return: [A1]
func (dt *DataTable) Get(str string, sheets ...string) (list []*DataFrame) {
	dt.Foreach(func(df *DataFrame) bool {
		if df.Value == str {
			list = append(list, df)
		}
		return false
	}, sheets...)
	return
}

// Contains Get contains str df
// default sheet[0]
// Example:
// -----------------
// A1	B1	C1
// A2	aa	bb
// A3	aa	bb
// -----------------
// Contains("A") // return: [A1, A2, A3]
// Contains("A", "Sheet1") // return: [A1, A2, A3]
func (dt *DataTable) Contains(str string, sheets ...string) (list []*DataFrame) {
	dt.Foreach(func(df *DataFrame) bool {
		if strings.Index(df.Value, str) >= 0 {
			list = append(list, df)
		}
		return false
	}, sheets...)
	return
}

// GetCol get col
// default sheet[0]
// Example:
// -----------------
// A1	B1	C1
// A2	aa	bb
// A3	aa	bb
// -----------------
// GetCol(1) // return: A1, A2, A3
// GetCol(2, "Sheet1") // return: B1, aa, aa
func (dt *DataTable) GetCol(col int, sheets ...string) (c *Col) {
	list := make([]*DataFrame, 0)
	dt.Foreach(func(df *DataFrame) bool {
		if df.Col == col {
			list = append(list, df)
		}
		return false
	}, sheets...)
	return newCol(col, list)
}

// GetRow get row
// default sheet[0]
// Example:
// -----------------
// A1	B1	C1
// A2	aa	bb
// A3	aa	bb
// -----------------
// GetRow(1) // return: A1, B1, C1
// GetRow(2, "Sheet1") // return: A2, aa, bb
func (dt *DataTable) GetRow(row int, sheets ...string) (r *Row) {
	dt.ForeachRow(func(idx int, rows *Row) bool {
		if idx == row {
			r = rows
			return true
		}
		return false
	}, sheets...)
	return
}

// GetCell get cell by row, col
// default sheet[0]
// Example:
// -----------------
// A1	B1	C1
// A2	aaa	bb
// A3	aaa	bb
// A4	aaa	bb
// -----------------
// GetCell(1, 1) // return: A1
// GetCell(1, 3, "Sheet1") // return: B1
func (dt *DataTable) GetCell(row, col int, sheets ...string) *DataFrame {
	obj := new(DataFrame)
	dt.Foreach(func(df *DataFrame) bool {
		if df.Row == row && df.Col == col {
			obj = df
			return true
		}
		return false
	}, sheets...)
	return obj
}

// GetCellByVal get cell by row str, col str
// default sheet[0]
// Example:
// -----------------
// A1	B1	C1
// A2	aaa	bb
// A3	aaa	bb
// A4	aaa	bb
// -----------------
// GetCellByVal("A4", "C1") // return: bb
// GetCellByVal("A4", "B1", "Sheet1") // return: aaa
func (dt *DataTable) GetCellByVal(row, col string, sheets ...string) *DataFrame {
	rowIdx := -1
	colIdx := -1
	obj := new(DataFrame)
	dt.Foreach(func(df *DataFrame) bool {
		if rowIdx == -1 && df.Value == row {
			rowIdx = df.Row
		}
		if colIdx == -1 && df.Value == col {
			colIdx = df.Col
		}
		if df.Row == rowIdx && df.Col == colIdx {
			obj = df
			return true
		}
		return false
	}, sheets...)
	return obj
}
