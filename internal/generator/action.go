// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

func nameParse(template Template, name, body string) (Template, error) {

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

func tmplParser(
	t *toDo,
	get func(string) string,
) itemDo {
	return func(item string) {

		var err error
		text := get(item)
		name := nameLessExt(item)
		t.tmpl, err = nameParse(t.tmpl, name, text)
		t.data.SeeError("CollectTmpl: Parse:", name, err)
	}
}

func metaParser(
	t *toDo,
	get func(string) string,
) itemDo {
	return func(item string) {

		var err error
		println("MetaParse: " + item)

		text := get(item)
		name := nameLessExt(item) + ".meta"

		meta, err := Meta(text) // extract meta-data
		t.data.SeeError("CollectMeta: Extract:", name, err)

		t.tmpl, err = nameParse(t.tmpl, name, meta) // Parse the meta-data
		t.data.SeeError("CollectMeta: Parse:", name, err)

		_, err = Apply(t.data, t.tmpl, name)
		t.data.SeeError("CollectMeta: Apply:", name, err)
	}
}
