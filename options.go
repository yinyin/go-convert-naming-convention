package convertnamingconvention

import (
	"maps"
	"strings"
)

func normalizeCommonInitialism(c string) string {
	return strings.ToUpper(strings.TrimSpace(c))
}

type Options struct {
	commonInitialisms             map[string]bool
	usingDefaultCommonInitialisms bool

	exceptionRules map[string]string

	splitAlphaNum bool
}

func NewDefaultOptions() *Options {
	return &Options{
		commonInitialisms:             defaultCommonInitialisms,
		usingDefaultCommonInitialisms: true,
	}
}

func (opts *Options) prepareCommonInitialisms() {
	if opts.usingDefaultCommonInitialisms {
		opts.commonInitialisms = maps.Clone(defaultCommonInitialisms)
		opts.usingDefaultCommonInitialisms = false
	} else if opts.commonInitialisms == nil {
		opts.commonInitialisms = make(map[string]bool)
	}
}

func (opts *Options) AddCommonInitialisms(commonInitialisms ...string) {
	opts.prepareCommonInitialisms()
	for _, commonInitial := range commonInitialisms {
		commonInitial = normalizeCommonInitialism(commonInitial)
		if commonInitial == "" {
			continue
		}
		opts.commonInitialisms[commonInitial] = true
	}
}

func (opts *Options) SetCommonInitialisms(commonInitialisms ...string) {
	opts.commonInitialisms = nil
	opts.AddCommonInitialisms(commonInitialisms...)
}

func (opts *Options) RemoveCommonInitialisms(commonInitialisms ...string) {
	if len(opts.commonInitialisms) == 0 {
		return
	}
	opts.prepareCommonInitialisms()
	for _, commonInitial := range commonInitialisms {
		commonInitial = normalizeCommonInitialism(commonInitial)
		delete(opts.commonInitialisms, commonInitial)
	}
}

func (opts *Options) SetExceptionRules(exceptionRules map[string]string) {
	opts.exceptionRules = maps.Clone(exceptionRules)
}

func (opts *Options) SplitAlphabetNumber() {
	opts.splitAlphaNum = true
}
