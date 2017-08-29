// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"github.com/golangsam/dotgo/internal/pile"
)

// nextPile represents a forward traversed Pile
type nextPile struct {
	*gen.StringPile
}

// NewNext returns a new dictionary
func NewNext(size, buff int) nextPile {
	return nextPile{gen.MakeStringPile(size, buff)}
}

func (p nextPile) Add(item string) {
	p.Pile(item)
}

func (p nextPile) Adder() nameDo {
	return func(item string) {
		p.Pile(item)
	}
}

func (p nextPile) S() []string {
	return <-p.Done()
}

func (p nextPile) Walker(t *toDo, out ...maker) func() {

	return func() {

		defer closeMaker(out...)
		for item, ok := p.Iter(); ok && t.ok(); item, ok = p.Next() {
			for i := range out {
				out[i].do(item)
			}
		}
	}
}