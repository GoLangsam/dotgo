// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"os"
	"path/filepath"

	"github.com/golangsam/container/ccsafe/lsm"
)

func (t *toDo) dirSWalker(
	dot bool,
	inp dirS,
	out ...filler,
) func() {

	return func() {

		defer func() {
			for i := range out {
				out[i].stuff.Close()
			}
		}()

		fh := func(path string, info os.FileInfo, err error) error {
			for i := range out {
				if out[i].match(path) {
					out[i].stuff.(*Pile).Pile(path)
				}
			}
			return nil
		}

		for i := 0; i < len(inp) && t.ok(); i++ {
			flagDot(dot, dotWalk) // ...

			dh := ifFlagSkipDirWf(matchBool(inp[i].Recurse)) // Recurse?
			filepath.Walk(inp[i].DirPath, t.isDirWf(dh, fh)) // Walk path
		}
	}
}

func (t *toDo) iter(inp Closer, out ...maker) func() {
	switch i := inp.(type) {
	case filler:
		return t.iter(i.stuff, out...)
	case maker:
		return t.iter(i.stuff, out...)
	case *Pile:
		return t.iterPile(i, out...)
	case *lsm.LazyStringerMap:
		return t.iterDict(i, out...)
	default:
		panic("No walker for this type")
	}
}

func (t *toDo) iterPile(inp *Pile, out ...maker) func() {

	return func() {

		defer closeMaker(out...)
		for item, ok := inp.Iter(); ok && t.ok(); item, ok = inp.Next() {
			for i := range out {
				out[i].do(item)
			}
		}
	}
}

func (t *toDo) iterDict(inp *lsm.LazyStringerMap, out ...maker) func() {

	return func() {

		defer closeMaker(out...)
		for _, item := range inp.S() {
			if !t.ok() {
				return // bail out
			}
			for i := range out {
				out[i].do(item)
			}
		}
	}
}

func closeFiller(out ...filler) {
	for i := range out {
		out[i].stuff.Close()
	}
}

func closeMaker(out ...maker) {
	for i := range out {
		out[i].stuff.Close()
	}
}

// ========================== old style ==========================

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

// ========================== pathDo ==========================

func tmplParser(
	t *toDo,
	get func(string) string,
	out *Pile,
) nameDo {
	return func(item string) {

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
}

func metaParser(
	t *toDo,
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
