package pitaya

import (
	"fmt"
	"strings"
)

type Col struct {
	idx int
	dfs []*DataFrame
}

func newCol(idx int, dfs []*DataFrame) *Col {
	return &Col{idx, dfs}
}

func (c *Col) String() string {
	return fmt.Sprintf("Col[%d], Rows: %d", c.idx, len(c.dfs))
}

func (c *Col) MaxRow() int {
	return len(c.dfs)
}

func (c *Col) Foreach(fn dfFunc) {
	for _, df := range c.dfs {
		if ok := fn(df); ok {
			return
		}
	}
}

func (c *Col) Last(str string) *DataFrame {
	var obj *DataFrame
	c.Foreach(func(df *DataFrame) bool {
		if df.Value == str {
			obj = df
		}
		return false
	})
	return obj
}

func (c *Col) First(str string) *DataFrame {
	var obj *DataFrame
	c.Foreach(func(df *DataFrame) bool {
		if df.Value == str {
			obj = df
			return true
		}
		return false
	})
	return obj
}

func (c *Col) Get(strs ...string) (list []*DataFrame) {
	if len(strs) == 0 {
		return c.dfs
	}
	c.Foreach(func(df *DataFrame) bool {
		for _, str := range strs {
			if df.Value == str {
				list = append(list, df)
			}
		}
		return false
	})
	return
}

func (c *Col) Contains(strs ...string) (list []*DataFrame) {
	c.Foreach(func(df *DataFrame) bool {
		for _, str := range strs {
			if strings.Index(df.Value, str) >= 0 {
				list = append(list, df)
			}
		}
		return false
	})
	return
}
