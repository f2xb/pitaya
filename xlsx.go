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

	dfs := make(map[string][]*Row)

	for _, sheet := range f.GetSheetList() {

		rows, err := f.Rows(sheet)
		if err != nil {
			return nil, err
		}

		maxRow, maxCol := 0, 0

		for rows.Next() {
			maxRow++
			cols, err := rows.Columns()
			if err != nil {
				return nil, err
			}
			if len(cols) > maxCol {
				maxCol = len(cols)
			}
		}
		if err = rows.Close(); err != nil {
			return nil, err
		}

		list := make([]*Row, maxRow)
		for rIdx := 0; rIdx < maxRow; rIdx++ {
			var rowDfs []*DataFrame
			for cIdx := 0; cIdx < maxCol; cIdx++ {
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

				rowDfs = append(rowDfs, &DataFrame{
					Col:      cIdx + 1,
					Row:      rIdx + 1,
					Sheet:    sheet,
					Value:    val,
					RawValue: rawVal,
				})
			}
			list[rIdx] = newRow(rIdx+1, rowDfs)
		}
		dfs[sheet] = list
	}

	return &DataTable{
		dfs:    dfs,
		sheets: f.GetSheetList(),
		option: option,
		Ext:    Xlsx,
	}, nil
}

// xlsxByXlsx github.com/tealeg/xlsx
func xlsxByXlsx(filePath string, option *Options) (*DataTable, error) {
	wb, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	dfs := make(map[string][]*Row)
	sheets := make([]string, 0)

	for _, sh := range wb.Sheets {
		sheets = append(sheets, sh.Name)

		list := make([]*Row, 0)

		rIdx := -1
		if err = sh.ForEachRow(func(r *xlsx.Row) error {
			rIdx++

			var rowDfs []*DataFrame

			cIdx := -1
			if err = r.ForEachCell(func(c *xlsx.Cell) error {
				cIdx++
				if option.TrimSpace {
					c.Value = strings.TrimSpace(c.Value)
				}
				rowDfs = append(rowDfs, &DataFrame{
					Col:   cIdx + 1,
					Row:   rIdx + 1,
					Sheet: sh.Name,
					Value: c.Value,
				})
				return nil
			}); err != nil {
				return err
			}
			list = append(list, newRow(rIdx+1, rowDfs))
			return nil
		}); err != nil {
			return nil, err
		}
		dfs[sh.Name] = list
	}

	return &DataTable{
		dfs:    dfs,
		sheets: sheets,
		option: option,
		Ext:    Xlsx,
	}, nil
}
