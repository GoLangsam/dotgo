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

func (p nextPile) S() []string {
	return <-p.Done()
}

func (p nextPile) Len() int {
	return len(<-p.Done())
}

func (p nextPile) Walker(quit func() bool, out ...*Actor) func() {

	return func() {

		defer ActorsClose(out...)
		for item, ok := p.Iter(); ok && !quit(); item, ok = p.Next() {
			ActorsDo(item, out...)
		}
	}
}

func (p nextPile) Action(is ...itemIs) *Actor {
	actor := Actor{p, func(item string) {
		for i := range is {
			if is[i](item) {
				p.Pile(item)
				return
			}
		}
	}}
	return &actor
}
