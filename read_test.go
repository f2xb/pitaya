package pitaya

import (
	"fmt"
	"strings"
	"testing"
)

func TestSheetRowContainsByRow(t *testing.T) {
	dt, err := Read("testdata/4.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	dfs := dt.RowContainsByRow(1, "账户")
	fmt.Println(dfs)
}

func TestName(t *testing.T) {
	fmt.Println(strings.Index("账户", "账户"))

}
