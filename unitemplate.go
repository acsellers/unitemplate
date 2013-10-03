package unitemplate

import (
	htmlTemplate "html/template"
	textTemplate "text/template"
	"text/template/parse"
)

type Parser interface {
	ParseFile(name, content string) (map[string]*parse.Tree, err)
	RequiredHtmlFuncs() htmlTemplate.FuncMap
	RequiredTextFuncs() textTemplate.FuncMap
	ApplicableExtensions() []string
}

var registeredParsers = make(map[string]Parser)

func RegisterParser(p Parser) {
	for _, extension := range p.ApplicableExtensions() {
		if _, found := registeredParsers[extension]; !found {
			registeredParsers[extension] = p
		}
	}
}
