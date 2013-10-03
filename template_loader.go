package unitemplate

import (
	htmlTemplate "html/template"
	textTemplate "text/template"
)

type TemplateLoader struct {
	holderHtml  *htmlTemplate.Template
	holderText  *textTemplate.Template
	viewPaths   []string
	layoutPaths []string
}

func NewTemplateLoader(viewPaths, layoutPaths []string) (*TemplateLoader, error) {
	return nil, nil
}

func (tl *TemplateLoader) Lookup(name string) (Template, bool) {

}

func (tl *TemplateLoader) Refresh() error {

}

type Template struct {
	HtmlItem *htmlTemplate.Template
	TextItem *textTemplate.Template
	Parent   *TemplateLoader
}
