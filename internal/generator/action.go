// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

// doer - just do something
func doer(do func()) *Actor       { a := Actor{NewNull(), func(item string) { do() }}; return &a }
func doit(do func(string)) *Actor { a := Actor{NewNull(), func(item string) { do(item) }}; return &a }

var noquit = func() bool { return false }

func (template Template) tmplParser(
	data Dot,
	get func(string) string,
) *Actor {
	actor := Actor{template, func(item string) {

		var err error
		text := get(item)
		name := nameLessExt(item)

		_, err = template.nameParse(name, text)
		data.SeeError("CollectTmpl: Parse:", name, err)
	}}
	return &actor
}

func (template Template) metaParser(
	data Dot,
	get func(string) string,
) *Actor {
	actor := Actor{template, func(item string) {

		var err error
		text := get(item)
		name := nameLessExt(item) + ".meta"

		meta, err := Meta(text) // extract meta-data
		data.SeeError("CollectMeta: Extract:", name, err)

		tmpl, err := template.nameParse(name, meta) // Parse the meta-data
		data.SeeError("CollectMeta: Parse:", name, err)

		_, err = Apply(data, tmpl, name)
		data.SeeError("CollectMeta: Apply:", name, err)
	}}
	return &actor
}

// nameParse is slightly similar to ParseFiles
func (template Template) nameParse(name, body string) (Template, error) {

	var err error
	var tmpl Template
	if name == template.Name() {
		tmpl = template
	} else {
		tmpl = Template{template.New(name)}
	}

	_, err = tmpl.Parse(body) // Parse the data
	return tmpl, err
}
