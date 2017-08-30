// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	t "github.com/golangsam/dotgo/internal/texttmpl" // adapter to "text/template"
	//"github.com/golangsam/dotgo/internal/htmltmpl" // adapter to "html/template"
)

// Template represents the template used (html or text)
type Template struct {
	t.Template
}

func flagPrintTemplate(flag bool, tmpl Template, prefix string) {
	t.IfPrintTemplate(flag, tmpl, prefix)
}

// NewTemplate returns a new template
// with funcmap attached and delimiters set
func NewTemplate(name string) Template {
	return Template{t.New(name)}
}

// Make returns a fresh Template made from parsing body
func (template Template) Make(name, body string) (new Template, err error) {
	tmpl, err := t.New(name).Parse(body)
	new = Template{tmpl}
	return new, err
}

// Meta returns the meta-text extraced from text
func Meta(text string) (string, error) {
	return t.Meta(text)
}

func (template Template) S() []string {
	return t.Names(template)
}

// Close - pretend to be a Closer (<=> an io.Closer)
func (template Template) Close() error {
	return nil
}
