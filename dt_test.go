package pitaya

import (
	"fmt"
	"testing"
)

func TestDataTable_Foreach(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	dt.Foreach(func(df *DataFrame) bool {
		fmt.Println(df.Value)
		return false
	})

	dt.Foreach(func(df *DataFrame) bool {
		fmt.Println(df.Value)
		return false
	}, "Sheet1")
}

func TestDataTable_ForeachRow(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	dt.ForeachRow(func(idx int, row *Row) bool {
		fmt.Println(idx, row)
		return false
	})

	dt.ForeachRow(func(idx int, row *Row) bool {
		fmt.Println(idx, row)
		return false
	}, "Sheet1")
}

func TestDataTable_MaxRow(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.MaxRow())
	fmt.Println(dt.MaxRow("Sheet1"))
}

func TestDataTable_GetSheets(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetSheets())
}

func TestDataTable_GetSheetByIndex(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetSheetByIndex(1))
	fmt.Println(dt.GetSheetByIndex(3))
}

func TestDataTable_First(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.First("D5"))
	fmt.Println(dt.First("D5", "Sheet1"))
}

func TestDataTable_Last(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.Last("D5"))
	fmt.Println(dt.Last("D5", "Sheet1"))
}

func TestDataTable_Rows(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range dt.Rows() {
		fmt.Println(row)
	}

	for _, row := range dt.Rows("Sheet1") {
		fmt.Println(row)
	}
}

func TestDataTable_Contains(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	for _, df := range dt.Contains("A") {
		fmt.Println(df)
	}
	fmt.Println()
	for _, df := range dt.Contains("B", "Sheet1") {
		fmt.Println(df)
	}
}

func TestDataTable_Get(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	for _, df := range dt.Get("D5") {
		fmt.Println(df)
	}
	fmt.Println()
	for _, df := range dt.Get("D5", "Sheet1") {
		fmt.Println(df)
	}
}

func TestDataTable_GetCol(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetCol(1))
	fmt.Println(dt.GetCol(2, "Sheet1"))
}

func TestDataTable_GetRow(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetRow(1))
	fmt.Println(dt.GetRow(3, "Sheet1"))
}

func TestDataTable_GetCell(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetCell(1, 1))
	fmt.Println(dt.GetCell(2, 3, "Sheet1"))
}

func TestDataTable_GetCellByVal(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetCellByVal("A1", "B1"))
	fmt.Println(dt.GetCellByVal("A1", "D1", "Sheet1"))
}
