package convertnamingconvention

func Split(name string, opts *Options) []string {
	if opts == nil {
		opts = NewDefaultOptions()
	}
	nameParts := splitNameWithoutException(name, opts.commonInitialisms, opts.splitAlphaNum)
	if len(opts.exceptionRules) == 0 {
		return nameParts
	}
	exceptRules := toExceptRules(opts.exceptionRules, opts.commonInitialisms, opts.splitAlphaNum)
	return applyExceptRules(nameParts, exceptRules)
}
