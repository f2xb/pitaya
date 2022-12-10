package pitaya

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	Xlsx = ".xlsx"
	Xls  = ".xls"
	Csv  = ".csv"
	Txt  = ".txt"
	Dat  = ".dat"
)

func Read(filePath string, opts ...Options) (*DataTable, error) {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case Xls:
		return ReadXls(filePath, opts...)
	case Xlsx:
		return ReadXlsx(filePath, opts...)
	case Csv:
		return ReadCsv(filePath, opts...)
	case Txt:
		return ReadTxt(filePath, opts...)
	case Dat:
		return ReadDat(filePath, opts...)
	}
	return nil, ErrNotSupportFileType
}

func read(filePath string, opts ...Options) (*DataTable, error) {
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
				Col:   cIdx + 1,
				Row:   rIdx + 1,
				Value: str,
				Sheet: sheet,
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
