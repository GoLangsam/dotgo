// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"path/filepath"

	"github.com/golangsam/container/ccsafe/dotpath"
	"github.com/golangsam/container/ccsafe/fileinfocache"
)

// DoIt performs it all:
//  - file system analysis (where to look)
//  - collection of metadata from templates
//  - execution of all relevant templates
// TODO dirS shall be absolute, if we intend to move os.Getwd during Exec
func DoIt() error {
	fsCache := fic.New()

	// lookupData retrieves the data related to file from cache
	lookupData := func(path string) string {
		return fsCache.LookupData(path)
	}

	patterns := filepath.SplitList(tmplread)
	tmplMatch := matchFunc(patterns...)
	baseMatch := matchFunc(patterns[0])
	execMatch := matchFunc(patterns[len(patterns)-1])
	//fakeMatch := matchBool(true)

	// roottmpl := filler{fakeMatch, NewTemplate(aDot)} // root template
	// _ = roottmpl
	pathPile := NewNext(512, 128) // files (templates) to handle
	metaPile := NewPrev(256, 64)  // templates with non-empty meta: apply in reverse order!
	baseDict := NewDict()         // templates to execute: basenames found
	execDict := NewDict()         // mathching file(s) identify folder(s) for execution

	pathMake := Actor{pathPile, func(item string) {
		if tmplMatch(item) {
			pathPile.Add(item)
		}
	}}

	metaMake := Actor{metaPile, func(item string) {
		meta, err := Meta(lookupData(item))
		if err == nil && meta != "" {
			pathPile.Add(item)
		}
	}}

	baseMake := Actor{baseDict, func(item string) {
		if baseMatch(item) {
			baseDict.Add(nameLessExt(item))
		}
	}}

	execMake := Actor{execDict, func(item string) {
		if execMatch(item) {
			execDict.Add(nameLessExt(item))
		}
	}}

	data := NewData(aDot)
	tmpl := NewTemplate(aDot)
	doit := doIt(data, tmpl)

	pathS := dotpath.DotPathS(flagArgs()...)
	flagPrintPathS(pma, pathS, "Args:")

	split := len(pathS) - 1        // at last:
	prepS := asDirS(pathS[:split]) // - prepare all but last
	execS := asDirS(pathS[split:]) // - execute only last

	if doit.ok() && len(prepS) > 0 {

		analyse := flagOpen(pm_, "Prepare:")
		flagPrintPathS(pma, prepS, "Prep:")

		flagWalk := Actor{NewNull(), func(string) {
			flagDot(pm_, dotWalk) // ...
		}}
		flagFOut := Actor{NewNull(), func(string) {
			flagDot(pm_, dotFOut) // ...
		}}
		flagTmpl := Actor{NewNull(), func(string) {
			flagDot(pm_, dotTmpl) // ...
		}}
		flagData := Actor{NewNull(), func(string) {
			flagDot(pm_, dotData) // ...
		}}

		tempPile := NewNext(512, 512)
		tempMake := Actor{tempPile, func(item string) {
			tempPile.Add(item)
		}}

		tp := tmplParser(doit, lookupData) //    tmplParser
		tmplMake := Actor{tempPile, tp}    // => doit.data

		doit.do(prepS.Walker(doit, flagWalk, pathMake))              // go prevS => tmplPile
		doit.do(pathMake.Walker(doit, flagFOut, baseMake, tempMake)) // go tmplPile => basePile & tempPile
		doit.do(tempMake.Walker(doit, flagTmpl, tmplMake, metaMake)) // go tempPile => metaPile & doit.tmpl
		doit.do(metaMake.Walker(doit, flagData))                     // go drain meta
		doit.wg.Wait()                                               // wait for all
		tempPile = NewNext(0, 0)                                     // forget
		tempMake = Actor{tempPile, func(string) {}}                  // forget

		doit.ifPrintPile(pmf, pathPile, "File:")
		doit.ifPrintPile(pmf, metaPile, "Meta:")
		// TODO doit.ifPrintPile(pmf, basePile, "Base:")
		doit.ifPrintTemplate(pmt, "Main:")
		if pmd { // build a throw-away DataTree
			mp := metaParser(doit, lookupData) // metaParser
			do := Actor{metaPile, mp}          // parse meta
			do.Walker(doit, flagFOut, do)()    // metaPile => doit.data
			doit.ifPrintDataTree(pmd, aDot)    // show
			doit.data = NewData(aDot)          // forget
		}

		flagClose(pm_, analyse)
	}

	if doit.ok() && !nox && !doit.ifPrintErrors("Prepare Main:") {
		//	todo.Execute()
		err := doit.exec(execS, pathMake, baseMake, metaMake, execMake, lookupData) // Execute
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
	tmplfile Actor,
	executes Actor,
	metadata Actor,
	writeout Actor,
	lookupData func(string) string,
) error {

	flagPrintPathS(pxa, execS, "Exec:")

	// we'll recurse!

	for i := range execS {
		_ = execS[i].DirPath
		_ = execS[i].Recurse

		foldPile := NewNext(16, 4) // folders / subdirectories found (not used here)
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
