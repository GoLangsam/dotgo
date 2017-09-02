// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"

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

// Beg implement Some

// S -
// wait for Done, and return content of pile
// Note: content is *not* reversed!
func (p prevPile) S() []string {
	return <-p.Done()
}

// Len -
// wait for Done, and return size of pile
func (p prevPile) Len() int {
	return len(<-p.Done())
}

// Close -
// inherited

// Walker -
// traverse the pile - backward
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

// flagPrint prints
// the pile (in reverse order),
// iff flag is true
func (p prevPile) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header, tab, cnt, p.Len(), tab, tab)

		if verbose {
			do := func(item string) { fmt.Println(tab, item, tab, tab) }
			p.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
		}
	}
}

// End implement Some

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

// End implement SomeWithAction
