package convertnamingconvention

import (
	"strings"
)

func ToLowerCamelCase(name string, commonInitialisms map[string]bool, exceptionRules map[string]string) string {
	nameParts := Split(name, commonInitialisms, exceptionRules)
	resultBuf := make([]byte, 0, computeBufferSize(nameParts, false))
	for i, part := range nameParts {
		if i == 0 {
			part = strings.ToLower(part)
		}
		resultBuf = append(resultBuf, []byte(part)...)
	}
	return string(resultBuf)
}

func ToUpperCamelCase(name string, commonInitialisms map[string]bool, exceptionRules map[string]string) string {
	nameParts := Split(name, commonInitialisms, exceptionRules)
	resultBuf := make([]byte, 0, computeBufferSize(nameParts, false))
	for _, part := range nameParts {
		resultBuf = append(resultBuf, []byte(part)...)
	}
	return string(resultBuf)
}
