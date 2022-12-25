package pitaya

import "fmt"

type dfFunc func(df *DataFrame) bool

type DataFrame struct {
	Col, Row int
	Value    string
	RawValue string
	Sheet    string
}

// String format print
func (df *DataFrame) String() string {
	return fmt.Sprintf("{Sheet:%s,Row:%d,Col:%d,Value:%s,RawValue:%s}",
		df.Sheet, df.Row, df.Col, df.Value, df.RawValue)
}
