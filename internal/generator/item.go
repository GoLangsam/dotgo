// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

// Item represents our string-based collections
type Item interface {
	Add(item string)                     // add item to content
	S() []string                         // content as stringS
	Walker(t *toDo, out ...maker) func() // provide traversal - interruptable
	Close() error                        // mimic io.Closer - definded locally in order to avoid explicit dependency
}
