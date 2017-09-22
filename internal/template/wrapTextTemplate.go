// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"text/template"
)

type textTmpl struct{ *template.Template }

// Text returns a new "text/template" Template
func Text(Name string) Template {
	return textTmpl{template.New(Name)}
}

// method wrappers - in alphabetical order

func (t textTmpl) Clone() (Template, error) {
	tmpl, err := t.Template.Clone()
	return textTmpl{tmpl}, err
}

func (t textTmpl) Delims(left, right string) Template {
	return textTmpl{t.Template.Delims(left, right)}
}

/* inherited:
func (t textTmpl) Execute(wr io.Writer, data interface{}) error {
	return t.Template.Execute(wr, data)
}

func (t textTmpl) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.Template.ExecuteTemplate(wr, name, data)
}
*/

func (t textTmpl) Funcs(funcMap map[string]interface{}) Template {
	return textTmpl{t.Template.Funcs(template.FuncMap(funcMap))}
}

func (t textTmpl) Lookup(name string) Template {
	return textTmpl{t.Template.Lookup(name)}
}

/* inherited:
func (t textTmpl) Name() string {
	return t.Template.Name()
}
*/

func (t textTmpl) New(name string) Template {
	return textTmpl{t.Template.New(name)}
}

func (t textTmpl) Option(opt ...string) Template {
	return textTmpl{t.Template.Option(opt...)}
}

func (t textTmpl) Parse(text string) (Template, error) {
	tmpl, err := t.Template.Parse(text)
	return textTmpl{tmpl}, err
}

func (t textTmpl) ParseFiles(filenames ...string) (Template, error) {
	tmpl, err := t.Template.ParseFiles(filenames...)
	return textTmpl{tmpl}, err
}

func (t textTmpl) ParseGlob(pattern string) (Template, error) {
	tmpl, err := t.Template.ParseGlob(pattern)
	return textTmpl{tmpl}, err
}

func (t textTmpl) Templates() []Template {
	tmps := t.Template.Templates()
	news := make([]Template, 0, len(tmps))
	for i := range tmps {
		news = append(news, textTmpl{tmps[i]})
	}
	return news
}
