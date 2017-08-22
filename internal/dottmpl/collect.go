// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

// Collect: build (and optionally print) template & data
func (t toDo) Collect() {
	t.CollectTmplS()
	t.CollectDataS()
	flagPrintTemplate(pmt, t.tmpl, "Main")
	flagPrintDataTree(pmd, t.data, ".")
}

// CollectTmplS: add files first-to-last to given template
func (t toDo) CollectTmplS() {
	for i := 0; i < len(t.dir.FsFileS); i++ {
		t.CollectTmpl(i)
	}
}

// CollectDataS: add files last-to-first to given data
func (t toDo) CollectDataS() {
	work, err := t.tmpl.Clone() // Duplicate tmpl!
	t.data.SeeError("Collect: Clone: ", t.tmpl.Name(), err)
	if todo, ok := t.doIt(t.data, work, t.dir); ok {
		for i := len(t.dir.FsFileS) - 1; i >= 0; i-- {
			todo.CollectData(i)
		}
	}
}

// CollectTmpl: add file to given template
func (t toDo) CollectTmpl(i int) {
	var err error
	file := t.dir.FsFileS[i]
	text := lookupData(file)
	name := file.BaseLessExt().String()
	t.tmpl, err = t.tmpl.New(name).Parse(text) // Parse the data
	t.data.SeeError("CollectTmpl: Parse:", name, err)
}

// CollectData: extract meta data from file lookupData, parse and execute it
func (t toDo) CollectData(i int) {
	var err error
	file := t.dir.FsFileS[i]
	text := lookupData(file)
	name := file.BaseLessExt().String() + "." + "meta" // new name: append .meta
	meta, err := Meta(text)                            // extract meta-data
	t.data.SeeError("CollectMeta: Extract:", name, err)

	t.tmpl, err = t.tmpl.New(name).Parse(meta) // Parse the meta-data
	t.data.SeeError("CollectMeta: Parse:", name, err)

	_, err = Apply(t.data, t.tmpl, name)
	t.data.SeeError("CollectMeta: Apply:", name, err)
}
