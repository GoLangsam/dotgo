// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"io/ioutil"
	"path/filepath"
)

func (s *step) execDo(d DirS) *step {
	d.flagPrint(ep, epv, "ep-Path:")

	for i := range d {
		s.do(func() { s.Clone().walkDir(d[i].DirPath, d[i].Recurse) })
	}
	return s
}

func (s *step) walkDir(path string, recurse bool) *step {

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
			s.do(func() { s.Clone().walkDir(path, recurse) })
		}

		if hasExecFiles && !nox {
			path := path
			s = s.execReadMetaAndPrint(path)
			s.baseDict.Walker(s.done, // for each base name
				doer(func() { s.do(func() { s.apply(path) }) }), // apply template and show/format/write result
			)()
		}
	}

	s.wg.Wait() // wait for all
	return s
}
