package pitaya

import (
	"fmt"
	"testing"
)

func TestTxt(t *testing.T) {
	dt, err := ReadTxt("testdata/01.txt")
	if err != nil {
		t.Fatal(err)
	}

	dfs := dt.GetCellByIndex(0, 1)
	for _, df := range dfs {
		fmt.Println(df)
	}

	dfs = dt.GetCell("汇安基金信银理财A类11期集合资产管理计划", "资产代码")
	for _, df := range dfs {
		fmt.Println(df)
	}
}
