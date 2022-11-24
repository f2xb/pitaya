package pitaya

import (
	"io/ioutil"
	"strings"
)

func ReadTxt(filePath string, opts ...Options) (*DataTable, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	option := parseOptions(opts...)
	option.AllSheet = false

	data := make(map[string][][]string)
	dfs := make(map[string][]*DataFrame)
	sheet := "Sheet1"
	rows := make([][]string, 0)

	for rIdx, row := range strings.Split(string(file), option.RowSep) {
		cols := make([]string, 0)
		for cIdx, str := range strings.Split(row, option.ColSep) {
			if option.TrimSpace {
				str = strings.TrimSpace(str)
			}
			cols = append(cols, str)
			dfs[sheet] = append(dfs[sheet], &DataFrame{
				ColIdx: cIdx,
				RowIdx: rIdx,
				Value:  str,
				Sheet:  sheet,
			})
		}
		rows = append(rows, cols)
	}
	data[sheet] = rows

	return &DataTable{
		data:   data,
		dfs:    dfs,
		sheets: []string{sheet},
		option: option,
	}, nil
}
