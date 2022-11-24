package pitaya

import (
	"strings"

	"github.com/f2xb/xls/xls"
)

func ReadXls(filePath string, opts ...Options) (*DataTable, error) {
	wb, err := xls.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	option := parseOptions(opts...)
	data := make(map[string][][]string)
	sheets := make([]string, wb.GetNumberSheets())
	dfs := make(map[string][]*DataFrame)

	for i, sheet := range wb.GetSheets() {
		shName := sheet.GetName()
		sheets[i] = shName
		rows := make([][]string, 0)
		for rIdx, row := range sheet.GetRows() {
			cols := make([]string, 0)
			for cIdx, cell := range row.GetCols() {
				val := cell.GetString()
				if option.TrimSpace {
					val = strings.TrimSpace(val)
				}
				cols = append(cols, val)
				dfs[shName] = append(dfs[shName], &DataFrame{
					ColIdx: cIdx,
					RowIdx: rIdx,
					Value:  val,
					Sheet:  shName,
				})
			}
			rows = append(rows, cols)
		}
		data[shName] = rows
	}

	return &DataTable{
		data:   data,
		dfs:    dfs,
		sheets: sheets,
		option: option,
	}, nil
}
