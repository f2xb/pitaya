package pitaya

import (
	"fmt"
	"testing"
)

func TestReadCsv(t *testing.T) {
	file := "testdata/1.csv"
	dt, err := Read(file)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.MaxRow())
}

func TestReadDat(t *testing.T) {
	file := "testdata/1.dat"
	// dt, err := Read(file)
	dt, err := Read(file, Options{
		RowSep: "",
		ColSep: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.MaxRow())
}

func TestReadTxt(t *testing.T) {
	file := "testdata/1.txt"
	// dt, err := Read(file)
	dt, err := Read(file, Options{
		RowSep: "",
		ColSep: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.MaxRow())
}

func TestReadXls(t *testing.T) {
	file := "testdata/1.xls"
	dt, err := Read(file)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.MaxRow())
	dt.ForeachRow(func(i int, row *Row) bool {
		fmt.Println(row)
		return false
	})
}

func TestReadXlsx(t *testing.T) {
	file := "testdata/4.xlsx"
	// dt, err := Read(file)
	dt, err := Read(file, Options{
		XlsxLib: "xlsx",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.MaxRow())
}
