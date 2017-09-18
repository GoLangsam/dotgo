// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"io"
	"text/template"
)

type textTmpl template.Template

func NewTextTemplate(Name string) Template {
	return textTmpl(*template.New(Name))
}

func (t textTmpl) New(name string) Template {
	return t.New(name)
}

func (t textTmpl) Clone() (Template, error) {
	return t.Clone()
}

func (t textTmpl) Parse(text string) (Template, error) {
	return t.Parse(text)
}

func (t textTmpl) Templates() []Template {
	return t.Templates()
}

func (t textTmpl) Name() string {
	return t.Name()
}

func (t textTmpl) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.ExecuteTemplate(wr, name, data)
}
