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

	data := make(map[string][][]string)
	dfs := make(map[string][]*DataFrame)
	sheet := "Sheet1"
	rows := make([][]string, len(records))

	for rIdx, record := range records {
		cols := make([]string, len(record))
		for cIdx, str := range record {
			if option.TrimSpace {
				str = strings.TrimSpace(str)
			}
			cols[cIdx] = str
			dfs[sheet] = append(dfs[sheet], &DataFrame{
				ColIdx: cIdx,
				RowIdx: rIdx,
				Value:  str,
				Sheet:  sheet,
			})
		}
		rows[rIdx] = cols
	}
	data[sheet] = rows

	return &DataTable{
		data:   data,
		dfs:    dfs,
		sheets: []string{sheet},
		option: option,
	}, nil
}
