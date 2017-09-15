// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"

	t "github.com/golangsam/dotgo/internal/texttmpl" // adapter to "text/template"
	//"github.com/golangsam/dotgo/internal/htmltmpl" // adapter to "html/template"
)

// Template represents the template used (html or text)
type Template struct {
	t.Template
}

// NewTemplate returns a new template
// with funcmap attached and delimiters set
func NewTemplate(name string) Template {
	return Template{t.New(name).Funcs(Funcs)}
}

// Beg implement Some

// S -
// return all Names
func (template Template) S() []string {
	return t.Names(template)
}

// Len -
// how many Names
func (template Template) Len() int {
	return len(template.S())
}

// Close -
// pretend to be a Closer (<=> an io.Closer)
func (template Template) Close() error {
	return nil
}

// Walker -
// traverse S() - the NameS
func (template Template) Walker(quit func() bool, out ...*Actor) func() {

	return func() {

		defer ActorsClose(out...)
		for _, item := range template.S() {
			if quit() {
				return // bail out
			}
			ActorsDo(item, out...)
		}
	}
}

// flagPrint prints the template, iff flag is true
func (template Template) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header, tab, cnt, template.Len(), tab, tab)

		if verbose {
			do := func(item string) { fmt.Println(tab, item, tab, tab) }
			template.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
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

// Meta returns the meta-text extraced from text
func Meta(text string) (string, error) {
	return t.Meta(text)
}
