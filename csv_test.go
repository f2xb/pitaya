package pitaya

import (
	"fmt"
	"testing"
)

func TestCsv(t *testing.T) {
	dt, err := ReadCsv("testdata/01.csv")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dt)
}
