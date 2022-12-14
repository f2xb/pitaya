package pitaya

type Options struct {
	Password  string
	TrimSpace bool
	AllSheet  bool
	ColSep    string // default \t
	RowSep    string // default \n
	XlsxLib   string // default 'excelize' or 'xlsx'
}

func parseOptions(opts ...Options) *Options {
	opt := &Options{}
	for _, o := range opts {
		opt = &o
	}
	if opt.RowSep == "" {
		opt.RowSep = "\n"
	}
	if opt.ColSep == "" {
		opt.ColSep = "\t"
	}
	return opt
}
