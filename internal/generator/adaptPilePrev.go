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
func NewPrev(size, buff int) prevPile {
	return prevPile{gen.MakeStringPile(size, buff)}
}

func (p prevPile) S() []string {
	return <-p.Done()
}

func (p prevPile) Len() int {
	return len(<-p.Done())
}

func (p prevPile) Walker(quit func() bool, out ...*Actor) func() {

	return func() {

		defer ActorsClose(out...)
		itemS := <-p.Done()
		count := len(itemS)
		for i := count - 1; i >= 0 && !quit(); i-- {
			ActorsDo(itemS[i], out...)
		}
	}
}

func (p prevPile) Action(is ...itemIs) *Actor {
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
