// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"path/filepath"

	"github.com/golangsam/container/ccsafe/lsm"
)

func (t *toDo) walkFS(
	dot bool,
	inp dirS,
	out *Pile,
	iff pathIs,
) func() {

	return func() {

		defer out.Close()

		fh := pathPiler(iff, out) // => populate tmplPile
		for i := 0; i < len(inp) && t.ok(); i++ {
			flagDot(dot, dotWalk) // ...

			dh := ifFlagSkipDirWf(matchBool(inp[i].Recurse)) // Recurse?
			filepath.Walk(inp[i].DirPath, t.isDirWf(dh, fh)) // Walk path
		}
	}
}

func (t *toDo) fanOut(
	dot bool,
	inp *Pile,
	out *lsm.LazyStringerMap,
	iff pathIs,
	dup *Pile,
) func() {

	return func() {

		defer dup.Close()

		for item, ok := inp.Iter(); ok && t.ok(); item, ok = inp.Next() {
			flagDot(dot, dotFOut) // ...
			t.itemFanOut(item, out, iff, dup)
		}
	}
}

func (t *toDo) parseT(
	dot bool,
	inp *Pile,
	out *Pile,
	get func(string) string,
) func() {

	return func() {

		defer out.Close()

		for item, ok := inp.Iter(); ok && t.ok(); item, ok = inp.Next() {
			flagDot(dot, dotTmpl) // ...
			t.itemParseT(item, out, get)
		}
	}
}

func (t *toDo) parseM(
	dot bool,
	inp *Pile,
	get func(string) string,
) func() {

	return func() {

		for item, ok := inp.Iter(); ok && t.ok(); item, ok = inp.Next() {
			flagDot(dot, dotData) // ...
			t.itemParseM(item, get)
		}
	}
}

func (t *toDo) itemFanOut(
	item string,
	out *lsm.LazyStringerMap,
	iff pathIs,
	dup *Pile,
) {
	dup.Pile(item)
	if iff(item) {
		out.Assign(nameLessExt(item), nil)
	}
}

func (t *toDo) itemParseT(
	item string,
	out *Pile,
	get func(string) string,
) {
	var err error
	text := get(item)
	name := nameLessExt(item)

	t.tmpl, err = t.tmpl.New(name).Parse(text) // Parse the data
	t.data.SeeError("CollectTmpl: Parse:", name, err)

	meta, err := Meta(text) // extract meta-data
	t.data.SeeError("CollectMeta: Extract:", name, err)

	if meta != "" { // has meta?
		out.Pile(item) // => populate metaPile
	}
}

func (t *toDo) itemParseM(
	item string,
	get func(string) string,
) {
	var err error
	text := get(item)
	name := nameLessExt(item) + ".meta"

	meta, err := Meta(text) // extract meta-data
	t.data.SeeError("CollectMeta: Extract:", name, err)

	t.tmpl, err = t.tmpl.New(name).Parse(meta) // Parse the meta-data
	t.data.SeeError("CollectMeta: Parse:", name, err)

	_, err = Apply(t.data, t.tmpl, name)
	t.data.SeeError("CollectMeta: Apply:", name, err)
}
