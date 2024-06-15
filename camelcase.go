package convertnamingconvention

import (
	"strings"
)

// ToLowerCamelCase converts the given name to lowerCamelCase.
// Leave opts as nil to use the default options.
func ToLowerCamelCase(name string, opts *Options) string {
	nameParts := Split(name, opts)
	resultBuf := make([]byte, 0, computeBufferSize(nameParts, false))
	for i, part := range nameParts {
		if i == 0 {
			part = strings.ToLower(part)
		}
		resultBuf = append(resultBuf, []byte(part)...)
	}
	return string(resultBuf)
}

// ToUpperCamelCase converts the given name to UpperCamelCase.
// Leave opts as nil to use the default options.
func ToUpperCamelCase(name string, opts *Options) string {
	nameParts := Split(name, opts)
	resultBuf := make([]byte, 0, computeBufferSize(nameParts, false))
	for _, part := range nameParts {
		resultBuf = append(resultBuf, []byte(part)...)
	}
	return string(resultBuf)
}
