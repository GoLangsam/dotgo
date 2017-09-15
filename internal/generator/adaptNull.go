// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"
)

// Null represents an empty Some
type Null struct{}

// NewNull return a fresh Null
func NewNull() Null {
	return Null{}
}

// Beg implement Some

// Len -
// is zero
func (n Null) Len() int {
	return 0
}

// Close -
// pretend to be a Closer (<=> an io.Closer)
func (n Null) Close() error {
	return nil
}

// Walker -
// pretend to walk the empty content :-)
func (n Null) Walker(quit func() bool, out ...Actor) func() {
	return func() { return }
}

// End implement Some

func (n Null) Action(is ...itemIs) Actor {
	return Actor{n, func(item string) {
		for i := range is {
			if is[i](item) {
				fmt.Println(tab, item, tab, tab)
				return
			}
		}
	}}
}

// End implement SomeWithAction
