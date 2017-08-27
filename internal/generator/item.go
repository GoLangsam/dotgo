// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

// Closer mimics io.Closer - in order to avoid explicit dependency
type Closer interface {
	Close() error
}

type item struct {
	stuff Closer // any collection which implements Closer
	match pathIs // a filter
}

type doit struct {
	item        // any collection which implements Closer
	do   pathDo // what to do to it
}
