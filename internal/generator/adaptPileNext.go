// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"github.com/golangsam/dotgo/internal/pile"
)

// nextPile represents a Pile traversed forward
type NextPile struct {
	*gen.StringPile
}

// NewNext returns a new dictionary
func NewNext(size, buff int) NextPile {
	return NextPile{gen.MakeStringPile(size, buff)}
}

// Beg implement Some

// Len -
// wait for Done, and return size of pile
func (p NextPile) Len() int {
	return len(<-p.Done())
}

// Close -
// inherited

// Walker -
// traverse the pile forward
// Note: may be used in parallel to p being populated, e.g. as some out *Actor
func (p NextPile) Walker(quit func() bool, out ...*Actor) func() {

	return func() {

		defer ActorsClose(out...)
		for item, ok := p.Iter(); ok && !quit(); item, ok = p.Next() {
			ActorsDo(item, out...)
		}
	}
}

// End implement Some

func (p NextPile) Action(is ...itemIs) *Actor {
	return &Actor{p, func(item string) {
		for i := range is {
			if is[i](item) {
				p.Pile(item)
				return
			}
		}
	}}
}

// End implement SomeWithAction
