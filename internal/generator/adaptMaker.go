// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

type maker struct {
	stuff Item   // any collection which implements Closer
	do    nameDo // what to do to it
}

func (m maker) S() []string {
	return m.stuff.S()
}

func (m maker) Close() error {
	return m.stuff.Close()
}

func (m maker) Walker(t *toDo, out ...maker) func() {
	return m.stuff.Walker(t, out...)
}

func (m maker) Add(item string) {
	m.stuff.Add(item)
}

func closeMaker(out ...maker) {
	for i := range out {
		out[i].stuff.Close()
	}
}
