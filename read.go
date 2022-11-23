package pitaya

import (
	"path/filepath"
	"strings"
)

func Read(filePath string, opts ...Options) (*DataTable, error) {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".xls":
		return ReadXls(filePath, opts...)
	case ".xlsx":
		return ReadXlsx(filePath, opts...)
	case "csv":
		return ReadCsv(filePath, opts...)
	case "txt":
		return ReadTxt(filePath, opts...)
	}
	return nil, ErrNotSupportFileType
}
