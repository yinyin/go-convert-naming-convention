package convertnamingconvention

func Split(name string, commonInitialisms map[string]bool, exceptionRules map[string]string) []string {
	nameParts := splitNameWithoutException(name, commonInitialisms)
	if len(exceptionRules) == 0 {
		return nameParts
	}
	exceptRules := toExceptRules(exceptionRules, commonInitialisms)
	return applyExceptRules(nameParts, exceptRules)
}
