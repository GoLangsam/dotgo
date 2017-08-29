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

// Close - pretend to be a Closer (<=> an io.Closer)
func (d Dict) Close() error {
	return nil
}

func (d Dict) Walker(t *toDo, out ...maker) func() {

	return func() {

		defer closeMaker(out...)
		for _, item := range d.S() {
			if !t.ok() {
				return // bail out
			}
			for i := range out {
				out[i].do(item)
			}
		}
	}
}

func (d Dict) Add(item string) {
	d.Assign(item, nil)
}

func (d Dict) Adder() nameDo {
	return func(item string) {
		d.Assign(item, nil)
	}
}
