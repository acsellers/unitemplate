package bham

import (
	"strings"
	"text/template/parse"
)

var (
	// Strict determines whether only tabs will be considered
	// as indentation operators (Strict == true) or whether
	// two spaces can be counted as an indentation operator
	// (Strict == false), this is included for haml
	// semi-comapibility
	Strict bool

	// To add multiple id declarations, the outputter puts them together
	// with a join string, by default this is an underscore
	IdJoin = "_"

	// Like the template library, you need to be able to set code delimeters
	LeftDelim  = "{{"
	RightDelim = "}}"
)

var Doctypes = map[string]string{
	"":             `<!DOCTYPE html>`,
	"Transitional": `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">`,
	"Strict":       `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`,
	"Frameset":     `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">`,
	"5":            `<!DOCTYPE html>`,
	"1.1":          `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`,
	"Basic":        `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.1//EN" "http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd">`,
	"Mobile":       `<!DOCTYPE html PUBLIC "-//WAPFORUM//DTD XHTML Mobile 1.2//EN" "http://www.openmobilealliance.org/tech/DTD/xhtml-mobile12.dtd">`,
	"RDFa":         `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML+RDFa 1.0//EN" "http://www.w3.org/MarkUp/DTD/xhtml-rdfa-1.dtd">`,
}

// parse will return a parse tree containing a single
func Parse(name, text string) (map[string]*parse.Tree, error) {
	pt := &protoTree{source: text, name: name}
	pt.lex()
	pt.analyze()
	pt.compile()
	i := strings.Index(name, ".bham")

	return map[string]*parse.Tree{
		name[:i] + name[i+5:]: pt.outputTree,
	}, pt.err
}

type protoTree struct {
	name       string
	source     string
	lineList   []templateLine
	nodes      []protoNode
	currNodes  []protoNode
	outputTree *parse.Tree
	err        error
}

type protoNode struct {
	level      int
	identifier int
	content    string
	filter     FilterHandler
	list       []protoNode
	elseList   []protoNode
}

func (pn protoNode) needsRuntimeData() bool {
	return strings.Contains(pn.content, LeftDelim) &&
		strings.Contains(pn.content, RightDelim)
}
