// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

func (s *step) prepDo(d DirS) *step {
	d.flagPrint(ap, apv, "ap-Prep:")

	// some Null Actors - for flagDot dotter  // ...
	flagPrep := doer(func() { flagDot(a_, ".") })
	flagTemp := doer(func() { flagDot(a_, "-") })
	flagFile := doer(func() { flagDot(a_, "~") })

	tempPile := NewNext(512, 128) // a temp - for fan-out file names
	testPile := NewPrev(256, 064) // templates with non-empty meta: apply in reverse order!

	s.do(
		d.Walker(s.done, flagPrep, // go dirS =>
			s.filePile.Action(s.isFile), // file
			tempPile.Action(s.isFile),   // temp
		))
	s.do(
		tempPile.Walker(s.done, flagTemp, // go temp =>
			s.baseDict.Action(s.isBase), // base
			testPile.Action(s.hasMeta),  // test (meta)
		))
	s.do(
		s.filePile.Walker(s.done, flagFile, // go file =>
			s.tmplParser(s.lookupData),   // rootTmpl
			s.metaPile.Action(s.hasMeta), // meta (*this* works!?!)
		))
	s.do(
		s.metaPile.Walker(s.done)) // go meta => drain
	s.do(
		testPile.Walker(s.done)) // go test => drain
	s.wg.Wait() // wait for all

	x := <-s.metaPile.Done()
	y := <-testPile.Done()
	if len(x) != len(y) {
		println("Sizes differ: meta =", len(x), "test =", len(y))
	}

	x = <-s.filePile.Done()
	y = <-tempPile.Done()
	if len(x) != len(y) {
		println("Sizes differ: file =", len(x), "temp =", len(y))
	}

	s.wg.Wait() // wait for all
	return s.prepReadMetaAndPrint()
}
