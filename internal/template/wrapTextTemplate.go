// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"text/template"
)

type textTmpl struct{ *template.Template }

func NewTextTemplate(Name string) Template {
	return textTmpl{template.New(Name)}
}

func (t textTmpl) New(name string) Template {
	return textTmpl{t.Template.New(name)}
}

func (t textTmpl) Clone() (Template, error) {
	tmpl, err := t.Template.Clone()
	return textTmpl{tmpl}, err
}

func (t textTmpl) Parse(text string) (Template, error) {
	tmpl, err := t.Template.Parse(text)
	return textTmpl{tmpl}, err
}

func (t textTmpl) Templates() []Template {
	tmpl := []Template{}
	for _, t := range t.Template.Templates() {
		tmpl = append(tmpl, textTmpl{t})
	}
	return tmpl
}

/*
func (t textTmpl) Name() string {
	return t.Template.Name()
}

func (t textTmpl) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.Template.ExecuteTemplate(wr, name, data)
}
*/

func (t textTmpl) Funcs(funcMap map[string]interface{}) Template {
	return textTmpl{t.Template.Funcs(template.FuncMap(funcMap))}
}
