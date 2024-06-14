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
		for k := range namingconv.DefaultCommonInitialisms {
			v.initialisms[k] = true
		}
	}
	parts := strings.Split(s, ",")
	for _, part := range parts {
		part = strings.ToUpper(part)
		v.initialisms[part] = true
	}
	return nil
}

func (v *commonInitialisms) Get() map[string]bool {
	if len(v.initialisms) == 0 {
		return namingconv.DefaultCommonInitialisms
	}
	return v.initialisms
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
	flag.Var(&commonInitialisms, "initial", "common initialisms")
	flag.Var(&exceptRules, "except", "exception rules")
	flag.Parse()
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "  ")
	for _, arg := range flag.Args() {
		r := Result{
			Splited:        namingconv.Split(arg, commonInitialisms.Get(), exceptRules.rules),
			LowerCamelCase: namingconv.ToLowerCamelCase(arg, commonInitialisms.Get(), exceptRules.rules),
			UpperCamelCase: namingconv.ToUpperCamelCase(arg, commonInitialisms.Get(), exceptRules.rules),
			SnakeCase:      namingconv.ToSnakeCase(arg, commonInitialisms.Get(), exceptRules.rules),
			KebebCase:      namingconv.ToKebebCase(arg, commonInitialisms.Get(), exceptRules.rules),
		}
		if err := w.Encode(r); nil != err {
			log.Fatalf("failed on encode result for [%s]: %v", arg, err)
			return
		}
	}
}
