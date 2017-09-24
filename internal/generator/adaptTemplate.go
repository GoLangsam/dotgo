// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"io"
	"sort"

	t "github.com/GoLangsam/template"
)

const (
	tmplL = "{{"
	tmplR = "}}"
	commL = "/*" // comment block beg
	commR = "*/" // comment block end
)

// template defines what is *really* used from t.Template
type template interface {
	// // AddParseTree(name string, tree *parse.Tree) (Template, error)
	Clone() (t.Template, error)
	// // DefinedTemplates() string
	// Delims(left, right string) Template
	// Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	// Funcs(funcMap map[string]interface{}) t.Template
	// Lookup(name string) Template
	Name() string
	New(name string) t.Template
	// Option(opt ...string) t.Template
	Parse(text string) (t.Template, error)
	// ParseFiles(filenames ...string) (Template, error)
	// ParseGlob(pattern string) (Template, error)
	Templates() []t.Template
}

// Template represents the template used (html or text)
type Template struct {
	template // t.Template
}

// NewTemplate returns a new template
// with funcmap attached and delimiters set
func NewTextTemplate(name string) Template {
	return Template{t.Text(name).Funcs(Funcs)}
}

// Beg implement Some

// Len -
// how many Names
func (template Template) Len() int {
	return len(template.Templates())
}

// Close -
// pretend to be a Closer (<=> an io.Closer)
func (template Template) Close() error {
	return nil
}

// Walker -
// traverse S() - the NameS
func (template Template) Walker(quit func() bool, out ...Actor) func() {

	return func() {

		defer ActorsClose(out...)
		names := Names(template)
		for i := range names {
			if quit() {
				return // bail out
			}
			ActorsDo(names[i], out...)
		}
	}
}

// End implement Some

// ParseName is slightly similar to ParseFiles
func (template Template) ParseName(name, body string) (tmpl Template, err error) {

	if name == template.Name() {
		tmpl = template
	} else {
		tmpl = Template{template.New(name)}
	}

	parsed, err := tmpl.Parse(body) // Parse the data
	return Template{parsed}, err
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
