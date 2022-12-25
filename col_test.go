package pitaya

import (
	"fmt"
	"testing"
)

func TestCol_MaxRow(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	col := dt.GetCol(1)
	fmt.Println(col.MaxRow())
}

func TestCol_Foreach(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	dt.GetCol(1).Foreach(func(df *DataFrame) bool {
		fmt.Println(df)
		return false
	})
}

func TestCol_Last(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetCol(2).Last("D5"))
}

func TestCol_First(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetCol(2).First("D5"))
}

func TestCol_Contains(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetCol(2).Contains("D5"))
}

func TestCol_Get(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetCol(2).Get("D5"))
}
