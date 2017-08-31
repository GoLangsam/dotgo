// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

type Actor struct {
	it Some   // Some collection
	do itemDo // what to do to it
}

func (m Actor) S() []string {
	return m.it.S()
}

func (m Actor) Close() error {
	return m.it.Close()
}

func (m Actor) Walker(quit func() bool, out ...*Actor) func() {
	return m.it.Walker(quit, out...)
}

func ActorsClose(out ...*Actor) {
	for i := range out {
		out[i].it.Close()
	}
}

func ActorsDo(item string, out ...*Actor) {
	for i := range out {
		out[i].do(item)
	}
}
