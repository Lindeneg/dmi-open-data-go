package cli

func Init() {
	cf := newConfig()
	cf.Parse()
	runActions(cf)
}
