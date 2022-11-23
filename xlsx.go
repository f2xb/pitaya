package pitaya

import (
	"strings"

	"github.com/xuri/excelize/v2"
)

func ReadXlsx(filePath string, opts ...Options) (*DataTable, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	option := parseOptions(opts...)
	opt := excelize.Options{}
	opt.RawCellValue = true
	if !option.RawValue {
		opt.RawCellValue = false
	}
	data := make(map[string][][]string)
	dfs := make(map[string][]*DataFrame)

	for _, sheet := range f.GetSheetList() {
		rows, err := f.GetRows(sheet, opt)
		if err != nil {
			return nil, err
		}
		row := make([][]string, len(rows))
		for rIdx, cols := range rows {
			col := make([]string, len(cols))
			for cIdx, cell := range cols {
				if option.TrimSpace {
					cell = strings.TrimSpace(cell)
				}
				col[cIdx] = cell
				dfs[sheet] = append(dfs[sheet], &DataFrame{
					ColIdx: cIdx,
					RowIdx: rIdx,
					Value:  cell,
					Sheet:  sheet,
				})
			}
			row[rIdx] = col
		}
		data[sheet] = row
	}

	return &DataTable{
		data:   data,
		dfs:    dfs,
		sheets: f.GetSheetList(),
		option: option,
	}, nil
}
