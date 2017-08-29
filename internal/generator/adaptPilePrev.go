// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"github.com/golangsam/dotgo/internal/pile"
)

// prevPile represents a forward traversed Pile
type prevPile struct {
	*gen.StringPile
}

// NewPrev returns a new dictionary
func NewPrev(size, buff int) nextPile {
	return nextPile{gen.MakeStringPile(size, buff)}
}

func (p prevPile) Add(item string) {
	p.Pile(item)
}

func (p prevPile) Adder() itemDo {
	return func(item string) {
		p.Pile(item)
	}
}

func (p prevPile) S() []string {
	return <-p.Done()
}

func (p prevPile) Walker(t *toDo, out ...Actor) func() {

	return func() {

		defer ActorsClose(out...)
		for item, ok := p.Iter(); ok && t.ok(); item, ok = p.Next() { // TODO must reverse!
			ActorsDo(item, out...)
		}
	}
}
