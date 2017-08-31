// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

func tmplParser(
	template Template,
	data Dot,
	get func(string) string,
) itemDo {
	return func(item string) {

		var err error
		text := get(item)
		name := nameLessExt(item)
		_, err = nameParse(template, name, text)
		data.SeeError("CollectTmpl: Parse:", name, err)
	}
}

func metaParser(
	template Template,
	data Dot,
	get func(string) string,
) itemDo {
	return func(item string) {

		var err error
		println("MetaParse: " + item)

		text := get(item)
		name := nameLessExt(item) + ".meta"

		meta, err := Meta(text) // extract meta-data
		data.SeeError("CollectMeta: Extract:", name, err)

		tmpl, err := nameParse(template, name, meta) // Parse the meta-data
		data.SeeError("CollectMeta: Parse:", name, err)

		_, err = Apply(data, tmpl, name)
		data.SeeError("CollectMeta: Apply:", name, err)
	}
}
