package convertnamingconvention

import (
	"strings"
)

// ToSnakeCase converts the given name to snake_case.
// Leave opts as nil to use the default options.
func ToSnakeCase(name string, opts *Options) string {
	nameParts := Split(name, opts)
	resultBuf := make([]byte, 0, computeBufferSize(nameParts, true))
	for i, part := range nameParts {
		part = strings.ToLower(part)
		if i != 0 {
			resultBuf = append(resultBuf, '_')
		}
		resultBuf = append(resultBuf, []byte(part)...)
	}
	return string(resultBuf)
}
