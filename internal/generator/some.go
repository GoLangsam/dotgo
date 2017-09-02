// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

// Some represents our walkable collections which may contain some strings
type Some interface {
	S() []string                                   // content as stringS
	Len() int                                      // length / count of content
	Walker(quit func() bool, out ...*Actor) func() // provide traversal thru items - interruptable
	Close() error                                  // mimic io.Closer - definded locally in order to avoid explicit dependency
	flagPrint(flag, verbose bool, header string)   // conditional print
}

// SomeWithAction represents some with an Action constructor
type SomeWithAction interface {
	Some
	Action(is ...itemIs) *Actor
}

// itemIs - given some path, returns a bool
type itemIs func(path string) bool

// itemDo - given some name, does sth
type itemDo func(name string)
