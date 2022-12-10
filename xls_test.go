package pitaya

import (
	"fmt"
	"testing"
)

func TestXls(t *testing.T) {
	dt, err := ReadXls("testdata/error.xls")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt.sheets)

	// fmt.Println(dt.HasRow(strings.Split("市场代码,资产名称,资产代码,估值日期,单位净值,我方持仓份额,万份收益", ",")...))

	// dfs := dt.GetCell("基金单位净值：", "科目名称")
	// for _, df := range dfs {
	// 	fmt.Println(df)
	// }
	//
	// for _, df := range dt.Contains("基金单位净值") {
	// 	fmt.Println(df)
	// }

	// fmt.Println(dt.GetSheet("Sheet1", "市场代码"))
	// fmt.Println(dt.First("市场代码"))
	// fmt.Println(dt.SheetFirst("Sheet1", "市场代码"))
	// fmt.Println(dt.Last("汇安基金信银理财A类11期集合资产管理计划"))
	// fmt.Println(dt.SheetLast("Sheet1", "汇安基金信银理财A类11期集合资产管理计划"))
	// fmt.Println(dt.First("汇安基金信银理财A类11期集合资产管理计划"))
	// fmt.Println(dt.Contains("代码"))
	// fmt.Println(dt.SheetContains("Sheet2", "代码"))

	// for _, df := range dt.GetColByIndex(0) {
	// 	fmt.Println(df)
	// }
	// for _, df := range dt.GetSheetColByIndex("Sheet1", 1) {
	// 	fmt.Println(df)
	// }

	// for _, df := range dt.GetRowByIndex(0) {
	// 	fmt.Println(df)
	// }
	// for _, df := range dt.GetSheetRowByIndex("Sheet1", 1) {
	// 	fmt.Println(df)
	// }
}
