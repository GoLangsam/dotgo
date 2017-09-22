// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"html/template"
)

type htmlTmpl struct{ *template.Template }

// Html returns a new "html/template" Template
func Html(Name string) Template {
	return htmlTmpl{template.New(Name)}
}

// method wrappers - in alphabetical order

func (t htmlTmpl) Clone() (Template, error) {
	tmpl, err := t.Template.Clone()
	return htmlTmpl{tmpl}, err
}

func (t htmlTmpl) Delims(left, right string) Template {
	return htmlTmpl{t.Template.Delims(left, right)}
}

/* inherited:
func (t htmlTmpl) Execute(wr io.Writer, data interface{}) error {
	return t.Template.Execute(wr, data)
}

func (t htmlTmpl) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.Template.ExecuteTemplate(wr, name, data)
}
*/

func (t htmlTmpl) Funcs(funcMap map[string]interface{}) Template {
	return htmlTmpl{t.Template.Funcs(template.FuncMap(funcMap))}
}

func (t htmlTmpl) Lookup(name string) Template {
	return htmlTmpl{t.Template.Lookup(name)}
}

/* inherited:
func (t htmlTmpl) Name() string {
	return t.Template.Name()
}
*/

func (t htmlTmpl) New(name string) Template {
	return htmlTmpl{t.Template.New(name)}
}

func (t htmlTmpl) Option(opt ...string) Template {
	return htmlTmpl{t.Template.Option(opt...)}
}

func (t htmlTmpl) Parse(text string) (Template, error) {
	tmpl, err := t.Template.Parse(text)
	return htmlTmpl{tmpl}, err
}

func (t htmlTmpl) ParseFiles(filenames ...string) (Template, error) {
	tmpl, err := t.Template.ParseFiles(filenames...)
	return htmlTmpl{tmpl}, err
}

func (t htmlTmpl) ParseGlob(pattern string) (Template, error) {
	tmpl, err := t.Template.ParseGlob(pattern)
	return htmlTmpl{tmpl}, err
}

func (t htmlTmpl) Templates() []Template {
	tmps := t.Template.Templates()
	news := make([]Template, 0, len(tmps))
	for i := range tmps {
		news = append(news, htmlTmpl{tmps[i]})
	}
	return news
}
