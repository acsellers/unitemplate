package bham

import (
	htmlTemplate "html/template"
	"strings"
	textTemplate "text/template"
	"text/template/parse"
)

type BhamParser string

func (bp BhamParser) ParseFile(name, content string) (map[string]*parse.Tree, error) {
	return Parse(name, content)
}

func (bp BhamParser) RequiredHtmlFuncs() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap{
		"bhamFilterConcat": func(filterName string, s ...string) htmlTemplate.HTML {
			for _, filter := range Filters {
				if filter.Trigger == filterName {
					return htmlTemplate.HTML(
						filter.Open + filter.Handler(strings.Join(s, "")) + filter.Close,
					)
				}
			}
			return htmlTemplate.HTML(strings.Join(s, ""))
		},
	}
}

func (bp BhamParser) RequiredTextFuncs() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"bhamFilterConcat": func(filterName string, s ...string) string {
			for _, filter := range Filters {
				if filter.Trigger == filterName {
					return filter.Open + filter.Handler(strings.Join(s, "")) + filter.Close
				}
			}
			return strings.Join(s, "")
		},
	}
}

func (bp BhamParser) ApplicableExtensions() []string {
	return []string{"bham"}
}
