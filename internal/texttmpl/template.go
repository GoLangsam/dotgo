// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"io"
	"sort"
	"text/template"

	ds "github.com/golangsam/do/strings" // Extract
)

const (
	tmplL = "{{"
	tmplR = "}}"
	commL = "/*"                // comment block beg
	commR = "*/"                // comment block end
	metaL = tmplL + commL + "-" // restrict comment to "{{/*-" for Meta-Comments
	metaR = "-" + commR + tmplR // restrict comment to "-*/}}" for Meta-Comments
)

// Template defines what is used from "text/template"
type Template interface {
	New(name string) *template.Template
	Clone() (*template.Template, error)
	Parse(text string) (*template.Template, error)
	Name() string
	Templates() []*template.Template
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

// New returns a new template with delimiters set
func New(name string) *template.Template {
	return template.New(name).Delims(tmplL, tmplR)
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
