unitemplate
===========

Unitemplate is a library for executing templates written in multiple formats 
cooperatively. For instance, a mustache template might call a template written
in a haml-ish languages, and the result would get inserted into a block from a
django style template.

The reason this works is that each format is supported by creating a parser for
that format that ouputs a number of Tree structs from Go's text/template/parse
library. Each Tree is then combined into into a html/template and text/template
template, so when you execute the Engine instance, you can have either an html
escaped form or a plain text form.

Current Formats:
* None

Formats in Progress:
* Haml-ish
* Mustache
* Django-ish
