package convertnamingconvention

import (
	"strings"
)

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
