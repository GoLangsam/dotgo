// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"io/ioutil"
	"path/filepath"
)

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

	return s
}

func (s *step) execDo(d DirS) *step {
	d.flagPrint(ep, epv, "ep-Path:")

	for i := range d {
		s.do(func() { s.Clone().execDir(d[i].DirPath, d[i].Recurse) })
	}
	return s
}

func (s *step) execDir(path string, recurse bool) *step {

	flagFile := doer(func() { flagDot(e_, ".") })
	s.do(
		s.filePile.Walker(s.done, flagFile, // go file =>
			s.tmplParser(s.lookupData),   // rootTmpl
			s.metaPile.Action(s.hasMeta), // meta (*this* works!?!)
			s.baseDict.Action(s.isBase),  // base
		))

	entries, err := ioutil.ReadDir(path)
	if all.Ok("ReadDir", path, err) {

		subDirS := []string{}
		hasExecFiles := false

		for _, entry := range entries {
			name := entry.Name()
			path := filepath.Join(path, name)

			if entry.IsDir() {
				if recurse && !IsDotNonsense(name) { // No .git or other dot nonsense please.
					subDirS = append(subDirS, path)
				}
			} else {
				if s.isFile(name) {
					s.filePile.Pile(path)
				}
				if !hasExecFiles && s.isExec(name) {
					hasExecFiles = true
				}
			}
		}
		s.filePile.Close()
		<-s.metaPile.Done() // go meta => drain

		for _, path := range subDirS {
			path := path
			s.do(func() { s.Clone().execDir(path, recurse) })
		}

		if hasExecFiles && !nox {
			path := path
			s.do(func() { s.execPath(path) })
		}
	}
	return s
}

// execPath - apply each template 'name' and show/format/write result
func (s *step) execPath(path string) {

	s = s.execReadMetaAndPrint(path)

	exec := s.apply(path) // apply templates
	s.do(s.baseDict.Walker(s.done, exec))
}

// ReadMetaAndPrint

func (s *step) prepReadMetaAndPrint() *step {

	// p f n d m r t => ?p f m n r t d
	s.filePile.flagPrint(af, afv, "af-File:")
	s.metaPile.flagPrint(am, amv, "am-Meta:")
	s.baseDict.flagPrint(an, anv, "an-Name:")
	s.rootTmpl.flagPrint(ar, arv, "ar-Root:")

	if ad || at { // temporary meta data - just to show
		s = s.readMeta(at, atv, "at-Data:")
		s.dataTree.flagPrint(ad, adv, "ad-"+aDot+aDot+aDot+aDot+":")
		s.dataTree = NewData(aDot) // forget
	}

	return s
}

func (s *step) execReadMetaAndPrint(path string) *step {
	// p f n d m r t => p f m n r t d
	flagPrintString(epv, path, "Directory")

	s.filePile.flagPrint(ef, efv, "ef-File:")
	s.metaPile.flagPrint(em, emv, "em-Meta:")
	s.baseDict.flagPrint(en, env, "en-Name:")
	s.rootTmpl.flagPrint(er, erv, "er-Root:")

	s = s.readMeta(et, etv, "et-Data:")
	s.dataTree.flagPrint(ed, edv, "ed-Data: "+path)

	return s
}
