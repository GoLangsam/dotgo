// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"github.com/golangsam/container/ccsafe/lsm"
)

// Dict represents a dictionary
type Dict struct {
	*lsm.LazyStringerMap
}

// NewDict returns a new dictionary
func NewDict() Dict {
	return Dict{lsm.New()}
}

// Beg implement Some

// Len -
// inherited

// Close -
// pretend to be a Closer (<=> an io.Closer)
func (d Dict) Close() error {
	return nil
}

// Walker -
// traverse the (sorted) keys of the dictionary
func (d Dict) Walker(quit func() bool, out ...Actor) func() {

	return func() {

		defer ActorsClose(out...)
		for _, item := range d.S() {
			if quit() {
				return // bail out
			}
			ActorsDo(item, out...)
		}
	}
}

// End implement Some

func (d Dict) Action(is ...itemIs) Actor {
	return Actor{d, func(item string) {
		for i := range is {
			if is[i](item) {
				d.Assign(nameLessExt(item), nil)
				return
			}
		}
	}}
}

// End implement SomeWithAction
