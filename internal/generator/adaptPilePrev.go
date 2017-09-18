// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"github.com/GoLangsam/dotgo/internal/pile"
)

// prevPile represents a Pile traversed backward
type PrevPile struct {
	*gen.StringPile
}

// NewPrev returns a new dictionary
func NewPrev(size, buff int) PrevPile {
	return PrevPile{gen.MakeStringPile(size, buff)}
}

// Beg implement Some

// Len -
// wait for Done, and return size of pile
func (p PrevPile) Len() int {
	return len(<-p.Done())
}

// Close -
// inherited

// Walker -
// traverse the pile - backward
func (p PrevPile) Walker(quit func() bool, out ...Actor) func() {

	return func() {

		defer ActorsClose(out...)
		itemS := <-p.Done()
		count := len(itemS)
		for i := count - 1; i >= 0 && !quit(); i-- {
			ActorsDo(itemS[i], out...)
		}
	}
}

// End implement Some

func (p PrevPile) Action(is ...itemIs) Actor {
	return Actor{p, func(item string) {
		for i := range is {
			if is[i](item) {
				p.Pile(item)
				return
			}
		}
	}}
}

// End implement SomeWithAction
