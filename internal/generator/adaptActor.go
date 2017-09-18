// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

type Actor struct {
	it   Some          // Some collection
	do   itemDo        // what to do to it
	done chan struct{} // done?
}

func Act(it Some, do itemDo) Actor {
	return Actor{
		it,
		do,
		make(chan struct{}),
	}
}

// Beg implement Some

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
func (m Actor) Walker(quit func() bool, out ...Actor) func() {
	return m.it.Walker(quit, out...)
}

// End implement Some

func (m Actor) Done() <-chan struct{} {
	return m.done
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
