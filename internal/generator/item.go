// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

// Closer mimics io.Closer - definded locally in order to avoid explicit dependency
type Closer interface {
	Close() error
}

type filler struct {
	match pathIs // a filter
	stuff Closer // any collection which implements Closer
}

func (f filler) Close() error {
	return f.stuff.Close()
}

type maker struct {
	stuff Closer // any collection which implements Closer
	do    nameDo // what to do to it
}

func (m maker) Close() error {
	return m.stuff.Close()
}

func (f filler) make(do nameDo) maker {
	return maker{f.stuff, do}
}

type nullcloser struct{}

func (n nullcloser) Close() error {
	return nil
}

func nullCloser() nullcloser {
	return nullcloser{}
}
