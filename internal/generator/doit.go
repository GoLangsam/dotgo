// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"path/filepath"

	"github.com/golangsam/container/ccsafe/dotpath"
	"github.com/golangsam/container/ccsafe/fileinfocache"
	"github.com/golangsam/container/ccsafe/lsm"
)

// dirS adopts the result of dotpath.DotPathS
type dirS []struct {
	DirPath string
	Recurse bool
}

func (d dirS) Close() error {
	return nil
}

// DoIt performs it all:
//  - file system analysis (where to look)
//  - collection of metadata from templates
//  - execution of all relevant templates
// TODO dirS shall be absolute, if we intend to move os.Getwd during Exec
func DoIt() error {
	patterns := filepath.SplitList(tmplread)
	tmplMatch := matchFunc(patterns...)
	baseMatch := matchFunc(patterns[0])
	execMatch := matchFunc(patterns[len(patterns)-1])
	fakeMatch := matchBool(true)

	roottmpl := filler{fakeMatch, NewTemplate(aDot)} // TODO tmplMatch
	_ = roottmpl
	tmplfile := filler{tmplMatch, MakePile(512, 128)} // files (templates) to handle
	metadata := filler{fakeMatch, MakePile(256, 64)}  // templates with non-empty meta: apply in reverse order!
	executes := filler{baseMatch, lsm.New()}          // templates to execute: basenames found
	writeout := filler{execMatch, lsm.New()}          // folder(s) for execution

	data := NewData(aDot)
	tmpl := NewTemplate(aDot)
	doit := doIt(data, tmpl)

	fsCache := fic.New()

	// lookupData retrieves the data related to file from cache
	lookupData := func(path string) string {
		return fsCache.LookupData(path)
	}

	pathS := dotpath.DotPathS(flagArgs()...)
	flagPrintPathS(pma, pathS, "Args:")

	split := len(pathS) - 1 // at last:
	prepS := pathS[:split]  // - prepare all but last
	execS := pathS[split:]  // - execute only last

	metaParser := maker{metadata.stuff, doit.metaParser(lookupData)}
	_ = metaParser

	if doit.ok() {
		err := doit.prep(prepS, tmplfile, executes, metadata, lookupData) // Analyse
		if err != nil && !doit.ifPrintErrors("Prepare Main:") {           // abort?
			doit.can()
			return err
		}
	}

	if doit.ok() && !nox {
		//	todo.Execute()
		err := doit.exec(execS, tmplfile, executes, metadata, writeout, lookupData) // Execute
		if err != nil && !doit.ifPrintErrors("Prepare Exec:") {                     // abort?
			doit.can()
			return err
		}
	}

	err := doit.ctx.Err()
	doit.can()
	return err
}

func (doit *toDo) prep(
	prepS dirS,
	template filler,
	executes filler,
	metadata filler,
	lookupData func(string) string,
) error {

	if len(prepS) < 1 {
		return nil // nothing to do
	}

	analyse := flagOpen(pm_, "Analyse:")
	flagPrintPathS(pma, prepS, "Prep:")

	tempPile := MakePile(512, 512)
	metaPile := metadata.stuff.(*Pile)
	baseDict := executes.stuff.(*lsm.LazyStringerMap)
	baseMatch := executes.match
	doit.do(doit.dirSWalker(pm_, prepS, template))                           // prevS => tmplPile
	doit.do(doit.fanOut(pm_, template, baseDict, baseMatch, tempPile))       // tmplPile => basePile & tempPile
	doit.do(doit.parseT(pm_, tempPile, metaPile, lookupData))                // tempPile => metaPile & doit.tmpl
	for _, ok := metaPile.Iter(); ok && doit.ok(); _, ok = metaPile.Next() { // wait for metaPile
		flagDot(pm_, dotData) // ...
	}
	doit.wg.Wait() // wait for all
	tempPile = nil // forget

	doit.ifPrintPile(pmf, template.stuff.(*Pile), "File:")
	doit.ifPrintPile(pmf, metaPile, "Meta:")
	// TODO doit.ifPrintPile(pmf, basePile, "Base:")
	doit.ifPrintTemplate(pmt, "Main:")
	if pmd { // build a throw-away DataTree
		doit.do(doit.parseM(pm_, metaPile, lookupData)) // metaPile => doit.data
		doit.wg.Wait()                                  // wait
		doit.ifPrintDataTree(pmd, aDot)                 // show
		doit.data = NewData(aDot)                       // forget
	}

	flagClose(pm_, analyse)

	return doit.ctx.Err()
}

func (doit *toDo) exec(
	execS dirS,
	template filler,
	executes filler,
	metadata filler,
	writeout filler,
	lookupData func(string) string,
) error {

	flagPrintPathS(pxa, execS, "Exec:")

	// we'll recurse!

	for i := range execS {
		_ = execS[i].DirPath
		_ = execS[i].Recurse

		foldPile := MakePile(16, 4) // folders / subdirectories found (not used here)
		foldMatch := matchBool(true)
		var downS dirS
		_, _, _ = foldPile, foldMatch, downS
		// we need to know:
		// new subdirs => downS = append(subS, struct{DirPath:?, Recurse: dirS[i].Recurse}
		// new tmplPile => add to curr
		// new templates => add to curr.Tmpl
		// new metaPile => add to curr
		// new execPile => add to curr
	}

	return doit.ctx.Err()
}
