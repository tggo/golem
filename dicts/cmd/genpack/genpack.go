package main

import (
	"flag"
	"os"
	"text/template"
)

type data struct {
	Locale string
}

// This is a code generator for language pack constructors. Check dicts/en/pack.go for an example
func main() {
	var d data
	flag.StringVar(&d.Locale, "locale", "", "The locale abbreviation this language pack is generated for")
	flag.Parse()

	t := template.Must(template.New("pack").Parse(packTemplate))
	t.Execute(os.Stdout, d)
}

var packTemplate = ` // Code generated by golem/dicts/cmd/generate_pack.go DO NOT EDIT
package {{.Locale}}

import "github.com/aaaton/golem/dicts"

const locale = "{{.Locale}}"

// LanguagePack is an implementation of the generic golem.LanguagePack interface for {{.Locale}}
type LanguagePack struct {
}

// NewPackage creates a language pack
func New() dicts.LanguagePack {
	return &LanguagePack{}
}

// GetResource returns the dictionary of lemmatized words
func (l *LanguagePack) GetResource() ([]byte, error) {
	return Asset("data/" + locale + ".gob")
}

// GetLocale returns the language name
func (l *LanguagePack) GetLocale() string {
	return locale
}

`