package convertnamingconvention

import (
	"strings"
)

func ToSnakeCase(name string, opts *Options) string {
	nameParts := Split(name, opts)
	resultBuf := make([]byte, 0, computeBufferSize(nameParts, false))
	for i, part := range nameParts {
		part = strings.ToLower(part)
		if i != 0 {
			resultBuf = append(resultBuf, '_')
		}
		resultBuf = append(resultBuf, []byte(part)...)
	}
	return string(resultBuf)
}
