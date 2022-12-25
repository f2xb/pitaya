package pitaya

import (
	"encoding/csv"
	"os"
	"strings"
)

func ReadCsv(filePath string, opts ...Options) (*DataTable, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	option := parseOptions(opts...)
	option.AllSheet = false

	dfs := make(map[string][]*Row)
	sheet := "Sheet1"

	list := make([]*Row, len(records))
	for rIdx, record := range records {
		var rowDfs []*DataFrame
		for cIdx, str := range record {
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
		list[rIdx] = newRow(rIdx+1, rowDfs)
	}

	dfs[sheet] = list
	return &DataTable{
		dfs:    dfs,
		sheets: []string{sheet},
		option: option,
		Ext:    Csv,
	}, nil
}
