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

func (s *step) tmplParser() *Actor {
	actor := Actor{s.rootTmpl, func(item string) {

		var err error
		text := s.lookupData(item)
		name := nameLessExt(item)

		_, err = s.rootTmpl.ParseName(name, text)
		all.Ok("CollectTmpl: Parse:", name, err)
	}}
	return &actor
}

func (s *step) metaReader(tmpl Template) *Actor {
	actor := Actor{tmpl, func(item string) {

		var err error
		text := s.lookupData(item)
		name := nameLessExt(item) + ".meta"

		meta, err := Meta(text) // extract meta-data
		all.Ok("CollectMeta: Extract:", name, err)

		tmpl, err := tmpl.ParseName(name, meta) // Parse the meta-data
		all.Ok("CollectMeta: Parse:", name, err)

		_, err = Apply(s.dataTree, tmpl, name)
		all.Ok("CollectMeta: Apply:", name, err)
	}}
	return &actor
}

func (s *step) readMeta(flag, verbose bool, header string) *step {

	tmpl, err := s.rootTmpl.Clone()          // Clone rootTmpl
	all.Ok("Clone", "Root", err)             // err? ignore for now
	metaData := s.metaReader(Template{tmpl}) // text/template from meta
	s.metaPile.Walker(s.done, metaData)()    // meta => metaTmpl & metaData
	metaData.flagPrint(flag, verbose, header)

	return s
}

func (s *step) apply(path string) *Actor {
	actor := Actor{s.rootTmpl, func(item string) {
		flagPrintString(wd, "Apply", path+tab+arr+item)

		// path - where we are
		// item - template name
		// id - root node name
		for _, node := range s.dataTree.DownS() {
			name := node.String()
			node := Data{node}
			byteS, err := Apply(node, s.rootTmpl, item)
			if all.Ok("Execute", item, err) {
				flagPrintByteS(wr, byteS, ">>>>Raw text of "+item+" & "+name)
				if ugo {
					filename := filepath.Join(path, node.FileName(nameLessExt(item)+".ugo"))
					all.Ok("Write Raw", filename, ioutil.WriteFile(filename, byteS, 0644))
				}
				if !nof {
					byteS, err = Source(byteS)
					all.Ok("Format", item, err)
				}
				flagPrintByteS(wf || nos, byteS, ">>>>Final text of "+item+" & "+name)
				filename := filepath.Join(path, node.FileName(item))
				if exe {
					all.Ok("Write", filename, ioutil.WriteFile(filename, byteS, 0644))
				}
			}

		}
	}}
	return &actor
}
