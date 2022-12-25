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

	dfs := make(map[string][]*Row)
	sheet := "Sheet1"

	list := make([]*Row, 0)
	for rIdx, row := range strings.Split(string(file), option.RowSep) {
		var rowDfs []*DataFrame
		for cIdx, str := range strings.Split(row, option.ColSep) {
			if option.TrimSpace {
				str = strings.TrimSpace(str)
			}
			rowDfs = append(rowDfs, &DataFrame{
				Col:   cIdx + 1,
				Row:   rIdx + 1,
				Value: str,
				Sheet: sheet,
			})
		}
		list = append(list, newRow(rIdx+1, rowDfs))
	}

	dfs[sheet] = list
	return &DataTable{
		dfs:    dfs,
		sheets: []string{sheet},
		option: option,
	}, nil
}
