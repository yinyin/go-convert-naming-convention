package convertnamingconvention

import (
	"maps"
	"strings"
)

func normalizeCommonInitialism(c string) string {
	return strings.ToUpper(strings.TrimSpace(c))
}

// Options can be used to customize the behavior of the naming convention conversion.
type Options struct {
	commonInitialisms             map[string]bool
	usingDefaultCommonInitialisms bool

	exceptionRules map[string]string

	splitAlphaNum bool
}

// NewDefaultOptions creates a new Options with default settings.
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

// AddCommonInitialisms adds given initialism into common initialisms list.
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

// SetCommonInitialisms sets the common initialisms list.
func (opts *Options) SetCommonInitialisms(commonInitialisms ...string) {
	opts.commonInitialisms = nil
	opts.AddCommonInitialisms(commonInitialisms...)
}

// RemoveCommonInitialisms removes given initialism from common initialisms list.
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

// SetExceptionRules sets the exception rules to override convention of matched words.
// The key of exceptionRules is the words to match.
// The value of exceptionRules is the words to replace with.
// Words to replace must concatnate with `-`, `_` or ` ` as delimiter.
func (opts *Options) SetExceptionRules(exceptionRules map[string]string) {
	opts.exceptionRules = maps.Clone(exceptionRules)
}

// SplitAlphabetNumber enables splitting alphabet and number.
func (opts *Options) SplitAlphabetNumber() {
	opts.splitAlphaNum = true
}
