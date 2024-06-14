package convertnamingconvention

func computeBufferSize(nameParts []string, withDelimiter bool) (bufferSize int) {
	for _, part := range nameParts {
		bufferSize += len(part)
	}
	if withDelimiter {
		bufferSize += (len(nameParts) - 1)
	}
	return
}
