package pitaya

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	filename := "testdata/01.txt"
	dt, err := Read(filename)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(dt)
}
