package convertnamingconvention

import (
	"slices"
	"sort"
)

type exceptRule struct {
	toMatch     []string
	replaceWith []string
}

type byToMatch []exceptRule

func (a byToMatch) Len() int      { return len(a) }
func (a byToMatch) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byToMatch) Less(i, j int) bool {
	toMatchLenI := len(a[i].toMatch)
	toMatchLenJ := len(a[j].toMatch)
	if toMatchLenI > toMatchLenJ {
		return true
	}
	if toMatchLenI < toMatchLenJ {
		return false
	}
	for idx := 0; idx < toMatchLenI; idx++ {
		if len(a[i].toMatch[idx]) < len(a[j].toMatch[idx]) {
			return true
		}
	}
	for idx := 0; idx < toMatchLenI; idx++ {
		if a[i].toMatch[idx] < a[j].toMatch[idx] {
			return true
		}
	}
	return false
}

func toExceptRules(rawRule map[string]string, commonInitialisms map[string]bool, splitAlphaNum bool) (exceptRules []exceptRule) {
	for toMatch, replaceWith := range rawRule {
		exceptRules = append(exceptRules, exceptRule{
			toMatch:     splitNameWithoutException(toMatch, commonInitialisms, splitAlphaNum),
			replaceWith: splitNameWithDelimiter(replaceWith, commonInitialisms),
		})
	}
	sort.Sort(byToMatch(exceptRules))
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
