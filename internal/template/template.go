// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"io"
	"sort"

	ds "github.com/GoLangsam/do/strings" // Extract
)

const (
	tmplL = "{{"
	tmplR = "}}"
	commL = "/*"                // comment block beg
	commR = "*/"                // comment block end
	metaL = tmplL + commL + "-" // restrict comment to "{{/*-" for Meta-Comments
	metaR = "-" + commR + tmplR // restrict comment to "-*/}}" for Meta-Comments
)

type FuncMap map[string]interface{}

// Template represents the template used (html or text)
type Template interface {
	// // AddParseTree(name string, tree *parse.Tree) (Template, error)
	Clone() (Template, error)
	// // DefinedTemplates() string
	// Delims(left, right string) Template
	// Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Funcs(funcMap map[string]interface{}) Template
	// Lookup(name string) Template
	Name() string
	New(name string) Template
	// Option(opt ...string) Template
	Parse(text string) (Template, error)
	// ParseFiles(filenames ...string) (Template, error)
	// ParseGlob(pattern string) (Template, error)
	Templates() []Template
}

// Names returns the sorted names of the templates referenced by tmpl
func Names(tmpl Template) (names []string) {
	tS := tmpl.Templates()
	for i := range tS {
		names = append(names, tS[i].Name())
	}
	sort.Slice(names, func(i, j int) bool { return (names[i] < names[j]) })
	return names
}

// Meta returns the meta-text extraced from text
func Meta(text string) (string, error) {
	return ds.Extract(text, metaL, metaR) // extract meta-data
}
