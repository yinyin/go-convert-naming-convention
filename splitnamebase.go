// Copyright (c) 2013 The Go Authors. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd.

// The code in this file heavily based on:
// https://github.com/golang/lint/blob/master/lint.go

package convertnamingconvention

import (
	"strings"
	"unicode"
)

func isDelimiter(r rune) bool {
	return r == '_' || r == '-' || unicode.IsSpace(r) || unicode.IsControl(r)
}

func appendNamePart(nameParts []string, runes []rune, w, i int, commonInitialisms map[string]bool) []string {
	word := string(runes[w:i])
	if u := strings.ToUpper(word); commonInitialisms[u] {
		nameParts = append(nameParts, u)
	} else {
		if strings.ToLower(word) == word {
			runes[w] = unicode.ToUpper(runes[w])
			word = string(runes[w:i])
		}
		nameParts = append(nameParts, word)
	}
	return nameParts
}

func splitNameWithDelimiter(name string, commonInitialisms map[string]bool) (nameParts []string) {
	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return []string{name}
	}
	runes := []rune(name)
	if isDelimiter(runes[0]) {
		runes[0] = '_'
	}
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word
		if i+1 == len(runes) {
			eow = true
		} else if isDelimiter(runes[i+1]) {
			eow = true
		}
		i++
		if !eow {
			continue
		}
		//
		// [w,i) is a word.
		nameParts = appendNamePart(nameParts, runes, w, i, commonInitialisms)
		for (i+1 <= len(runes)) && isDelimiter(runes[i]) {
			i++
		}
		w = i
	}
	return

}

func splitNameWithoutException(name string, commonInitialisms map[string]bool) (nameParts []string) {
	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return []string{name}
	}
	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	if allLower {
		return []string{name}
	}
	//
	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	if isDelimiter(runes[0]) {
		runes[0] = '_'
	}
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word
		if i+1 == len(runes) {
			eow = true
		} else if isDelimiter(runes[i+1]) {
			eow = true
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}
		//
		// [w,i) is a word.
		nameParts = appendNamePart(nameParts, runes, w, i, commonInitialisms)
		for (i+1 <= len(runes)) && isDelimiter(runes[i]) {
			i++
		}
		w = i
	}
	return
}
