// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

func tmplParser(
	t *toDo,
	template Template,
	get func(string) string,
) itemDo {
	return func(item string) {

		var err error
		text := get(item)
		name := nameLessExt(item)
		_, err = nameParse(template, name, text)
		t.data.SeeError("CollectTmpl: Parse:", name, err)
	}
}

func metaParser(
	t *toDo,
	template Template,
	get func(string) string,
) itemDo {
	return func(item string) {

		var err error
		println("MetaParse: " + item)

		text := get(item)
		name := nameLessExt(item) + ".meta"

		meta, err := Meta(text) // extract meta-data
		t.data.SeeError("CollectMeta: Extract:", name, err)

		tmpl, err := nameParse(template, name, meta) // Parse the meta-data
		t.data.SeeError("CollectMeta: Parse:", name, err)

		_, err = Apply(t.data, tmpl, name)
		t.data.SeeError("CollectMeta: Apply:", name, err)
	}
}
