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

	roottmpl := filler{fakeMatch, NewTemplate(aDot)} // root template
	_ = roottmpl
	tmplfile := filler{tmplMatch, MakePile(512, 128)} // files (templates) to handle
	metaFill := filler{fakeMatch, MakePile(256, 64)}  // templates with non-empty meta: apply in reverse order!
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

	if doit.ok() && len(prepS) > 0 {

		analyse := flagOpen(pm_, "Prepare:")
		flagPrintPathS(pma, prepS, "Prep:")

		flagFOut := maker{filler{fakeMatch, nullCloser()}, func(string) {
			flagDot(pm_, dotFOut) // ...
		}}
		flagTmpl := maker{filler{fakeMatch, nullCloser()}, func(string) {
			flagDot(pm_, dotTmpl) // ...
		}}
		flagData := maker{filler{fakeMatch, nullCloser()}, func(string) {
			flagDot(pm_, dotData) // ...
		}}

		tempPile := MakePile(512, 512)
		baseDict := executes.stuff.(*lsm.LazyStringerMap)
		baseMatch := executes.match
		doit.do(doit.dirSWalker(pm_, prepS, tmplfile))                     // go prevS => tmplPile
		doit.do(doit.fanOut(pm_, tmplfile, baseDict, baseMatch, tempPile)) // go tmplPile => basePile & tempPile
		tp := tmplParser(doit, lookupData, metaFill.stuff.(*Pile))         //    tmplParser
		tmplMake := maker{tempPile, tp}                                    // => doit.data
		doit.do(doit.iter(tempPile, tmplMake, flagTmpl))                   // go tempPile => metaPile & doit.tmpl
		doit.do(doit.iter(metaFill, flagData))                             // go drain meta
		doit.wg.Wait()                                                     // wait for all
		tempPile = nil                                                     // forget

		doit.ifPrintPile(pmf, tmplfile.stuff.(*Pile), "File:")
		doit.ifPrintPile(pmf, metaFill.stuff.(*Pile), "Meta:")
		// TODO doit.ifPrintPile(pmf, basePile, "Base:")
		doit.ifPrintTemplate(pmt, "Main:")
		if pmd { // build a throw-away DataTree
			mp := metaParser(doit, lookupData)  // metaParser
			do := metaFill.make(mp)             // parse meta
			doit.iter(do.stuff, flagFOut, do)() // metaPile => doit.data
			doit.ifPrintDataTree(pmd, aDot)     // show
			doit.data = NewData(aDot)           // forget
		}

		flagClose(pm_, analyse)
	}

	if doit.ok() && !nox && !doit.ifPrintErrors("Prepare Main:") {
		//	todo.Execute()
		err := doit.exec(execS, tmplfile, executes, metaFill, writeout, lookupData) // Execute
		if err != nil && !doit.ifPrintErrors("Prepare Exec:") {                     // abort?
			doit.can()
			return err
		}
	}

	err := doit.ctx.Err()
	doit.can()
	return err
}

func (doit *toDo) exec(
	execS dirS,
	tmplfile filler,
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
