package pitaya

func ReadTxt(filePath string, opts ...Options) (dt *DataTable, err error) {
	dt, err = read(filePath, opts...)
	if err != nil {
		return nil, err
	}
	dt.Ext = Txt
	return
}
