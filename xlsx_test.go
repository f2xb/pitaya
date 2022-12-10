package pitaya

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestReadXlsx(t *testing.T) {
	dt, err := ReadXlsx("testdata/004.xlsx", Options{
		XlsxLib: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	// for _, frame := range dt.RowContains("名称,代码,日期") {
	//	fmt.Println(frame)
	// }
	// for _, frame := range dt.SheetRowContains("Sheet1", "名称,代码,日期") {
	//	fmt.Println(frame)
	// }
	// for _, frame := range dt.RowContainsByRowIdx(1, "名称,代码,日期") {
	//	fmt.Println(frame)
	// }
	// for _, frame := range dt.SheetRowContainsByRowIdx(1, "Sheet1", "名称,代码,日期") {
	//	fmt.Println(frame)
	// }

	fmt.Println(dt.GetRowByRowIdx(2))

	// fmt.Println(dt.Contains())
}

func TestName(t *testing.T) {
	d, err := excelize.ExcelDateToTime(44888, false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(d)
}
