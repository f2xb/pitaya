package pitaya

import (
	"fmt"
	"testing"
)

func TestRow_MaxCol(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetRow(1).MaxCol())
}

func TestRow_Foreach(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	dt.GetRow(1).Foreach(func(df *DataFrame) bool {
		fmt.Println(df)
		return false
	})
}

func TestRow_Last(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetRow(1).Last("C1"))
}

func TestRow_First(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetRow(1).First("D1"))
}

func TestRow_Get(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetRow(1).Get("C1"))
}

func TestRow_Contains(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.GetRow(1).Contains("B"))
}
