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

// NewTemplate returns a new template
// with funcmap attached and delimiters set
func NewTemplate(name string) Template {
	return Template{t.New(name).Funcs(Funcs)}
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
