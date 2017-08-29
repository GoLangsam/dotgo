// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

func tmplParser(
	t *toDo,
	get func(string) string,
) itemDo {
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
) itemDo {
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
