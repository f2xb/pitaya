package pitaya

import (
	"fmt"
	"strings"
)

type rowFunc func(int, *Row) bool

type Row struct {
	idx int
	dfs []*DataFrame
}

func newRow(idx int, dfs []*DataFrame) *Row {
	return &Row{idx, dfs}
}

func (r *Row) String() string {
	return fmt.Sprintf("Row[%d], Cols: %d", r.idx, len(r.dfs))
}

func (r *Row) MaxCol() int {
	return len(r.dfs)
}

func (r *Row) Foreach(fn dfFunc) {
	for _, df := range r.dfs {
		if ok := fn(df); ok {
			return
		}
	}
}

func (r *Row) Last(str string) *DataFrame {
	var obj *DataFrame
	r.Foreach(func(df *DataFrame) bool {
		if df.Value == str {
			obj = df
		}
		return false
	})
	return obj
}

func (r *Row) First(str string) *DataFrame {
	var obj *DataFrame
	r.Foreach(func(df *DataFrame) bool {
		if df.Value == str {
			obj = df
			return true
		}
		return false
	})
	return obj
}

func (r *Row) Get(strs ...string) (list []*DataFrame) {
	if len(strs) == 0 {
		return r.dfs
	}
	r.Foreach(func(df *DataFrame) bool {
		for _, str := range strs {
			if df.Value == str {
				list = append(list, df)
			}
		}
		return false
	})
	return
}

func (r *Row) Contains(strs ...string) (list []*DataFrame) {
	r.Foreach(func(df *DataFrame) bool {
		for _, str := range strs {
			if strings.Index(df.Value, str) >= 0 {
				list = append(list, df)
			}
		}
		return false
	})
	return
}
