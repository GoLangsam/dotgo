// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

type Actor struct {
	it Some   // any collection which implements Closer
	do itemDo // what to do to it
}

func (m Actor) S() []string {
	return m.it.S()
}

func (m Actor) Close() error {
	return m.it.Close()
}

func (m Actor) Walker(t *toDo, out ...Actor) func() {
	return m.it.Walker(t, out...)
}

func ActorsClose(out ...Actor) {
	for i := range out {
		out[i].it.Close()
	}
}

func ActorsDo(item string, out ...Actor) {
	for i := range out {
		out[i].do(item)
	}
}
