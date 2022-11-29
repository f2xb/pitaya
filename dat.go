package pitaya

func ReadDat(filePath string, opts ...Options) (dt *DataTable, err error) {
	dt, err = read(filePath, opts...)
	if err != nil {
		return nil, err
	}
	dt.Ext = Dat
	return
}
