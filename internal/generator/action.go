// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"io/ioutil"
	"path/filepath"
)

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

func (template Template) apply(
	path string,
	data Dot,
) *Actor {
	actor := Actor{template, func(item string) {
		flagPrintString(wd, "Apply", data.String()+tab+arr+item)
		byteS, err := Apply(data, template, item)
		if !data.SeeError("Execute", item, err) {
			flagPrintByteS(wr, byteS, ">>>>Raw text of "+item+" & "+data.String())
			if ugo {
				filename := filepath.Join(path, FileName(data, item+".ugo"))
				data.SeeError("Write Raw", filename, ioutil.WriteFile(filename, byteS, 0644))
			}
			if !nof {
				byteS, err = Source(byteS)
				data.SeeError("Format", item, err)
			}
			flagPrintByteS(wf || nos, byteS, ">>>>Final text of "+item+" & "+data.String())
			filename := filepath.Join(path, FileName(data, item))
			if exe {
				data.SeeError("Write", filename, ioutil.WriteFile(filename, byteS, 0644))
			}
		}
	}}
	return &actor
}

// FileName resolves name as a template, executed against data
func FileName(data Dot, name string) string {
	id := "FileName"
	fileName := name
	template := NewTemplate(id)
	if tmpl, err := template.Parse(fileName); err == nil {
		if byteS, err := Apply(data, Template{tmpl}, id); err == nil {
			fileName = string(byteS)
		} else {
			panic(id + ": Apply: Error: " + err.Error())
		}
	} else {
		panic(id + ": Parse: Error: " + err.Error())
	}
	return fileName
}
