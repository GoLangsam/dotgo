// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

type Actor struct {
	it Some   // Some collection
	do itemDo // what to do to it
}

// Beg implement Some

// S -
// delegate to it
func (m Actor) S() []string {
	return m.it.S()
}

// Len -
// delegate to it
func (m Actor) Len() int {
	return m.it.Len()
}

// Close -
// delegate to it
func (m Actor) Close() error {
	return m.it.Close()
}

// Walker -
// delegate to it
func (m Actor) Walker(quit func() bool, out ...*Actor) func() {
	return m.it.Walker(quit, out...)
}

// flagPrint -
// delegate to it
func (m Actor) flagPrint(flag, verbose bool, header string) {
	m.it.flagPrint(flag, verbose, header)
}

// End implement Some

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
