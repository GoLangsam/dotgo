// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"html/template"
)

type htmlTmpl struct{ *template.Template }

func NewHtmlTemplate(Name string) Template {
	return htmlTmpl{template.New(Name)}
}

func (t htmlTmpl) New(name string) Template {
	return htmlTmpl{t.Template.New(name)}
}

func (t htmlTmpl) Clone() (Template, error) {
	tmpl, err := t.Template.Clone()
	return htmlTmpl{tmpl}, err
}

func (t htmlTmpl) Parse(text string) (Template, error) {
	tmpl, err := t.Template.Parse(text)
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

/*
func (t htmlTmpl) Name() string {
	return t.Template.Name()
}

func (t htmlTmpl) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.Template.ExecuteTemplate(wr, name, data)
}
*/

func (t htmlTmpl) Funcs(funcMap map[string]interface{}) Template {
	return htmlTmpl{t.Template.Funcs(template.FuncMap(funcMap))}
}
