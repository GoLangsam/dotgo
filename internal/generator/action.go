// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"os"
	"path/filepath"
)

func (t *toDo) dirSWalker(
	dot bool,
	inp dirS,
	out ...maker,
) func() {

	return func() {

		defer closeMaker(out...)
		fh := func(path string, info os.FileInfo, err error) error {
			for i := range out {
				out[i].do(path)
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

// ========================== old style ==========================

func (t *toDo) fanOut(
	dot bool,
	inp maker,
	out maker,
	iff itemIs,
	dup maker,
) func() {

	return func() {

		defer dup.Close()
		pile := inp.stuff.(nextPile)
		for item, ok := pile.Iter(); ok && t.ok(); item, ok = pile.Next() {
			flagDot(dot, dotFOut) // ...
			t.itemFanOut(item, out, iff, dup)
		}
	}
}

func (t *toDo) itemFanOut(
	item string,
	out maker,
	iff itemIs,
	dup maker,
) {
	// TODO dup.Pile(item)
	if iff(item) {
		// TODO out.Assign(nameLessExt(item), nil)
	}
}

// ========================== pathDo ==========================

func tmplParser(
	t *toDo,
	get func(string) string,
) nameDo {
	return func(item string) {

		var err error
		text := get(item)
		name := nameLessExt(item)

		t.tmpl, err = t.tmpl.Make(name, text) // Parse the data
		t.data.SeeError("CollectTmpl: Parse:", name, err)
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
