package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	namingconv "github.com/yinyin/go-convert-naming-convention"
)

type exceptRules struct {
	rules map[string]string
}

func (v exceptRules) String() string {
	var parts []string
	for toMatch, replaceWith := range v.rules {
		parts = append(parts, toMatch+"="+replaceWith)
	}
	return strings.Join(parts, ",")
}

func (v *exceptRules) Set(s string) error {
	if v.rules == nil {
		v.rules = make(map[string]string)
	}
	parts := strings.Split(s, ",")
	for _, part := range parts {
		p0 := strings.Split(part, "=")
		if len(p0) < 2 {
			return fmt.Errorf("invalid exception rule: [%s=]", part)
		}
		v.rules[p0[0]] = strings.Join(p0[1:], "-")
	}
	return nil
}

func (v *exceptRules) ApplyTo(opts *namingconv.Options) {
	if v.rules == nil {
		return
	}
	opts.SetExceptionRules(v.rules)
}

type commonInitialisms struct {
	initialisms map[string]bool
}

func (v commonInitialisms) String() string {
	var parts []string
	for k := range v.initialisms {
		parts = append(parts, k)
	}
	return strings.Join(parts, ",")
}

func (v *commonInitialisms) Set(s string) error {
	if v.initialisms == nil {
		v.initialisms = make(map[string]bool)
	}
	parts := strings.Split(s, ",")
	for _, part := range parts {
		part = strings.ToUpper(strings.TrimSpace(part))
		if part == "" {
			continue
		}
		v.initialisms[part] = true
	}
	return nil
}

func (v *commonInitialisms) ApplyTo(opts *namingconv.Options) {
	if v.initialisms == nil {
		return
	}
	initials := make([]string, 0, len(v.initialisms))
	for k := range v.initialisms {
		initials = append(initials, k)
	}
	opts.AddCommonInitialisms(initials...)
}

type Result struct {
	Splited []string

	LowerCamelCase string
	UpperCamelCase string
	SnakeCase      string
	KebebCase      string
}

func main() {
	var commonInitialisms commonInitialisms
	var exceptRules exceptRules
	var splitAlphaNum bool
	flag.Var(&commonInitialisms, "initial", "common initialisms")
	flag.Var(&exceptRules, "except", "exception rules")
	flag.BoolVar(&splitAlphaNum, "splitAlphaNum", false, "split alphabets and numbers")
	flag.Parse()
	opts := namingconv.NewDefaultOptions()
	commonInitialisms.ApplyTo(opts)
	exceptRules.ApplyTo(opts)
	if splitAlphaNum {
		opts.SplitAlphabetNumber()
	}
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "  ")
	for _, arg := range flag.Args() {
		r := Result{
			Splited:        namingconv.Split(arg, opts),
			LowerCamelCase: namingconv.ToLowerCamelCase(arg, opts),
			UpperCamelCase: namingconv.ToUpperCamelCase(arg, opts),
			SnakeCase:      namingconv.ToSnakeCase(arg, opts),
			KebebCase:      namingconv.ToKebebCase(arg, opts),
		}
		if err := w.Encode(r); nil != err {
			log.Fatalf("failed on encode result for [%s]: %v", arg, err)
			return
		}
	}
}
