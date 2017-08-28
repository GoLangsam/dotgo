// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"path/filepath"

	"github.com/golangsam/container/ccsafe/lsm"
)

func (t *toDo) dirSWalker(
	dot bool,
	inp dirS,
	out filler,
) func() {

	return func() {

		defer out.stuff.Close()

		fh := pathPiler(out.match, out.stuff.(*Pile)) // => populate tmplPile
		for i := 0; i < len(inp) && t.ok(); i++ {
			flagDot(dot, dotWalk) // ...

			dh := ifFlagSkipDirWf(matchBool(inp[i].Recurse)) // Recurse?
			filepath.Walk(inp[i].DirPath, t.isDirWf(dh, fh)) // Walk path
		}
	}
}

func (t *toDo) pileWalker(
	dot bool,
	inp *Pile,
	out ...maker,
) func() {

	return func() {

		defer func() {
			for i := range out {
				out[i].stuff.Close()
			}
		}()
		for item, ok := inp.Iter(); ok && t.ok(); item, ok = inp.Next() {
			flagDot(dot, dotFOut) // ...
			for i := range out {
				out[i].do(item)
			}
		}
	}
}

func (t *toDo) dictWalker(
	dot bool,
	inp *lsm.LazyStringerMap,
	out ...maker,
) func() {

	return func() {

		defer func() {
			for i := range out {
				out[i].stuff.Close()
			}
		}()
		for _, item := range inp.S() {
			if !t.ok() {
				return // bail out
			}
			flagDot(dot, dotFOut) // ...
			for i := range out {
				out[i].do(item)
			}
		}
	}
}

func (t *toDo) fanOut(
	dot bool,
	inp filler,
	out *lsm.LazyStringerMap,
	iff pathIs,
	dup *Pile,
) func() {

	return func() {

		defer dup.Close()
		pile := inp.stuff.(*Pile)
		for item, ok := pile.Iter(); ok && t.ok(); item, ok = pile.Next() {
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

	t.tmpl, err = t.tmpl.Make(name, text) // Parse the data
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

	t.tmpl, err = t.tmpl.Make(name, meta) // Parse the meta-data
	t.data.SeeError("CollectMeta: Parse:", name, err)

	_, err = Apply(t.data, t.tmpl, name)
	t.data.SeeError("CollectMeta: Apply:", name, err)
}

func (t *toDo) metaParser(
	get func(string) string,
) nameDo {
	return func(item string) {

		var err error
		text := get(item)
		name := nameLessExt(item) + ".meta"

		meta, err := Meta(text) // extract meta-data
		t.data.SeeError("CollectMeta: Extract:", name, err)

		t.tmpl, err = t.tmpl.Make(name, meta) // Parse the meta-data
		t.data.SeeError("CollectMeta: Parse:", name, err)

		_, err = Apply(t.data, t.tmpl, name)
		t.data.SeeError("CollectMeta: Apply:", name, err)
	}
}
