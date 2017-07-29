// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	t "github.com/golangsam/dotgo/internal/texttmpl" // adapter to "text/template"
	//"github.com/golangsam/dotgo/internal/htmltmpl" // adapter to "html/template"
)

// Template represents the template used (html or text)
type Template interface {
	t.Template
}

func flagPrintTemplate(flag bool, tmpl Template, prefix string) {
	t.IfPrintTemplate(flag, tmpl, prefix)
}

func NewTemplate(name string) Template {
	return t.New(name)
}

func Meta(text string) (string, error) {
	return t.Meta(text)
}
