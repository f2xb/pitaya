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

	sheets := make([]string, wb.GetNumberSheets())

	dfs := make(map[string][]*Row)

	for i, sheet := range wb.GetSheets() {
		shName := sheet.GetName()
		sheets[i] = shName
		list := make([]*Row, sheet.GetNumberRows())
		for rIdx, row := range sheet.GetRows() {
			var rowDfs []*DataFrame
			for cIdx, cell := range row.GetCols() {
				val := cell.GetString()
				if option.TrimSpace {
					val = strings.TrimSpace(val)
				}
				rowDfs = append(rowDfs, &DataFrame{
					Col:   cIdx + 1,
					Row:   rIdx + 1,
					Value: val,
					Sheet: shName,
				})
			}
			list = append(list, newRow(rIdx+1, rowDfs))
		}
		dfs[shName] = list
	}

	return &DataTable{
		dfs:    dfs,
		sheets: sheets,
		option: option,
		Ext:    Xls,
	}, nil
}
