// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"html/template"
	"io"
)

type htmlTmpl template.Template

func NewHtmlTemplate(Name string) Template {
	return htmlTmpl(*template.New(Name))
}

func (t htmlTmpl) New(name string) Template {
	return t.New(name)
}

func (t htmlTmpl) Clone() (Template, error) {
	return t.Clone()
}

func (t htmlTmpl) Parse(text string) (Template, error) {
	return t.Parse(text)
}

func (t htmlTmpl) Templates() []Template {
	return t.Templates()
}

func (t htmlTmpl) Name() string {
	return t.Name()
}

func (t htmlTmpl) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.ExecuteTemplate(wr, name, data)
}
