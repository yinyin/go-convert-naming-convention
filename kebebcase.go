package convertnamingconvention

import (
	"strings"
)

// ToKebebCase converts the given name to kebeb-case.
// Leave opts as nil to use the default options.
func ToKebebCase(name string, opts *Options) string {
	nameParts := Split(name, opts)
	resultBuf := make([]byte, 0, computeBufferSize(nameParts, true))
	for i, part := range nameParts {
		part = strings.ToLower(part)
		if i != 0 {
			resultBuf = append(resultBuf, '-')
		}
		resultBuf = append(resultBuf, []byte(part)...)
	}
	return string(resultBuf)
}
