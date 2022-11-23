package pitaya

type Options struct {
	Password  string
	RawValue  bool
	TrimSpace bool
	AllSheet  bool
	ColSep    string
	RowSep    string
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
