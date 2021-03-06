// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"html/template"
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

// Template defines what is used from "html/template"
type Template interface {
	New(name string) *template.Template
	Clone() (*template.Template, error)
	Parse(text string) (*template.Template, error)
	Name() string
	Templates() []*template.Template
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

// New returns a new template with funcmap attached and delimiters set
func New(name string) *template.Template {
	return template.New(name).Funcs(Funcs).Delims(tmplL, tmplR)
}

// Names returns the sorted names of the templates referenced by tmpl
func Names(tmpl Template) (names []string) {
	tS := tmpl.Templates()
	sort.Slice(tS, func(i, j int) bool { return (tS[i].Name() < tS[j].Name()) })
	for i := range tS {
		names = append(names, tS[i].Name())
	}
	return
}

// Meta returns the meta-text extraced from text
func Meta(text string) (string, error) {
	return ds.Extract(text, metaL, metaR) // extract meta-data
}
