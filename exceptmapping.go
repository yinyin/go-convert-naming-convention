package convertnamingconvention

import (
	"slices"
)

type exceptRule struct {
	toMatch     []string
	replaceWith []string
}

func toExceptRules(rawRule map[string]string, commonInitialisms map[string]bool, splitAlphaNum bool) (exceptRules []exceptRule) {
	for toMatch, replaceWith := range rawRule {
		exceptRules = append(exceptRules, exceptRule{
			toMatch:     splitNameWithoutException(toMatch, commonInitialisms, splitAlphaNum),
			replaceWith: splitNameWithDelimiter(replaceWith, commonInitialisms),
		})
	}
	return
}

func applyExceptRules(nameParts []string, exceptRules []exceptRule) []string {
	lenNameParts := len(nameParts)
	for _, rule := range exceptRules {
		lenToMatch := len(rule.toMatch)
		for i := 0; i <= (lenNameParts - lenToMatch); i++ {
			if slices.Equal(nameParts[i:i+lenToMatch], rule.toMatch) {
				appliedNameParts := make([]string, 0, lenNameParts-lenToMatch+len(rule.replaceWith))
				appliedNameParts = append(appliedNameParts, nameParts[:i]...)
				appliedNameParts = append(appliedNameParts, rule.replaceWith...)
				appliedNameParts = append(appliedNameParts, nameParts[i+lenToMatch:]...)
				return appliedNameParts
			}
		}
	}
	return nameParts
}
