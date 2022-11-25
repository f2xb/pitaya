package pitaya

import (
	"strings"

	"github.com/tealeg/xlsx/v3"
	"github.com/xuri/excelize/v2"
)

func ReadXlsx(filePath string, opts ...Options) (*DataTable, error) {
	option := parseOptions(opts...)
	switch option.XlsxLib {
	case "xlsx":
		return xlsxByXlsx(filePath, option)
	default:
		return xlsxByExcelize(filePath, option)
	}
}

// xlsxByExcelize github.com/xuri/excelize
func xlsxByExcelize(filePath string, option *Options) (*DataTable, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := make(map[string][][]string)
	dfs := make(map[string][]*DataFrame)

	for _, sheet := range f.GetSheetList() {
		rows, err := f.GetRows(sheet)
		if err != nil {
			return nil, err
		}
		row := make([][]string, len(rows))
		for rIdx, cols := range rows {
			col := make([]string, len(cols))

			for cIdx := range rows[rIdx] {
				axis, err := excelize.CoordinatesToCellName(cIdx+1, rIdx+1)
				if err != nil {
					return nil, err
				}
				val, err := f.GetCellValue(sheet, axis)
				if err != nil {
					return nil, err
				}
				rawVal, err := f.GetCellValue(sheet, axis, excelize.Options{
					RawCellValue: true,
				})
				if err != nil {
					return nil, err
				}

				if option.TrimSpace {
					val = strings.TrimSpace(val)
					rawVal = strings.TrimSpace(rawVal)
				}

				col[cIdx] = val

				dfs[sheet] = append(dfs[sheet], &DataFrame{
					ColIdx:   cIdx,
					RowIdx:   rIdx,
					Sheet:    sheet,
					Value:    val,
					RawValue: rawVal,
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

// xlsxByXlsx github.com/tealeg/xlsx
func xlsxByXlsx(filePath string, option *Options) (*DataTable, error) {
	wb, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	data := make(map[string][][]string)
	dfs := make(map[string][]*DataFrame)
	sheets := make([]string, 0)

	for _, sh := range wb.Sheets {
		sheets = append(sheets, sh.Name)

		rows := make([][]string, 0)
		var rIdx, cIdx int

		if err = sh.ForEachRow(func(r *xlsx.Row) error {
			rIdx++
			cols := make([]string, 0)

			if err = r.ForEachCell(func(c *xlsx.Cell) error {
				cIdx++
				if option.TrimSpace {
					c.Value = strings.TrimSpace(c.Value)
				}
				cols = append(cols, c.Value)
				dfs[sh.Name] = append(dfs[sh.Name], &DataFrame{
					ColIdx: cIdx,
					RowIdx: rIdx,
					Sheet:  sh.Name,
					Value:  c.Value,
				})
				return nil
			}); err != nil {
				return err
			}
			rows = append(rows, cols)
			return nil
		}); err != nil {
			return nil, err
		}
		data[sh.Name] = rows
	}

	return &DataTable{
		data:   data,
		dfs:    dfs,
		sheets: sheets,
		option: option,
	}, nil
}
