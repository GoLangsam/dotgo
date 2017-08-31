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
//  - establish the participators
//  - file system analysis (where to look)
//  - collection of metadata from templates
//  - execution of all relevant templates
// TODO dirS shall be absolute, if we intend to move os.Getwd during Exec
func DoIt() error {

	// Beg of Prolog - we have something to declare

	pathS := AsDirS(dotpath.DotPathS(flagArgs()...))
	pathS.flagPrint(ap, apv, "Args:")

	split := len(pathS) - 1           // at last:
	prepDirS := AsDirS(pathS[:split]) // - prepare all but last
	execDirS := AsDirS(pathS[split:]) // - execute only last

	fsCache := fic.New()
	lookupData := func(path string) string { return fsCache.LookupData(path) }

	suffix := filepath.SplitList(tmplread)
	isFile := matchFunc(suffix...)
	isBase := matchFunc(suffix[0])
	isExec := matchFunc(suffix[len(suffix)-1])
	isTrue := matchBool(true)
	hasMeta := func(item string) bool {
		meta, err := Meta(lookupData(item))
		if err != nil {
			panic(err) // TODO don't panic
		}
		return meta != ""
	}

	// Actors - Some containers, and how to populate each
	fileMake := NewNext(512, 128).Action(isFile)                   // files (templates) to handle
	metaMake := NewPrev(256, 064).Action(hasMeta)                  // templates with non-empty meta: apply in reverse order!
	baseMake := NewDict().Action(isBase)                           // templates to execute: basenames found
	execMake := NewDict().Action(isExec)                           // TODO this is wrong: we need the directory! nameLessExt get's appended to base            // mathching file(s) identify folder(s) for execution
	rootData := NewData(aDot)                                      // data - a Dot
	metaData := NewData(aDot)                                      // data for metaParser
	rootTmpl := NewTemplate(aDot).tmplParser(rootData, lookupData) // text/template
	metaTmpl := NewTemplate(aDot).metaParser(metaData, lookupData) // text/template from meta
	doit := doIt(rootData)                                         // carries context, and data

	// doer - just do something
	doer := func(do func()) *Actor { a := Actor{NewNull(), func(string) { do() }}; return &a }

	// some Null Actors - for flagDot dotter  // ...
	flagWalk := doer(func() { flagDot(a_, dotWalk) })
	flagFOut := doer(func() { flagDot(a_, dotFOut) })
	flagTmpl := doer(func() { flagDot(a_, dotTmpl) })
	flagData := doer(func() { flagDot(a_, dotData) })

	// End of Prolog

	if doit.ctx.Err() == nil && len(prepDirS) > 0 { // Beg of prep Analysis
		analyse := flagOpen(a_, "Prepare:")
		prepDirS.flagPrint(ap, ap, "Prep:")

		// a temp - for fan-out file names
		tempMake := NewNext(128, 32).Action(isTrue)

		quit := func() bool { return doit.ctx.Err() != nil }         // quit, iff not doit.ok
		doit.do(prepDirS.Walker(quit, flagWalk, tempMake, fileMake)) // go prepS => temp & file path
		doit.do(tempMake.Walker(quit, flagFOut, metaMake, baseMake)) // go temp => meta & base
		doit.do(fileMake.Walker(quit, flagTmpl, rootTmpl))           // go file => rootTmpl
		doit.do(metaMake.Walker(quit, flagData, metaTmpl))           // go meta => metaTmpl & metaData
		doit.wg.Wait()                                               // wait for all
		tempMake = NewNext(0, 0).Action(isTrue)                      // forget temp

		rootTmpl.flagPrint(at, atv, "Main:")
		metaTmpl.flagPrint(at, atv, "Meta:")
		doit.ifPrintDataTree(ad, aDot)
		fileMake.flagPrint(af, afv, "File:")
		metaMake.flagPrint(am, amv, "Meta:")
		baseMake.flagPrint(an, anv, "Base:")

		doit.data = NewData(aDot) // forget
		flagClose(a_, analyse)
	} // End of prep Analysis

	if doit.ctx.Err() == nil && !nox && !doit.ifPrintErrors("Prepare Main:") {
		//	todo.Execute()
		err := doit.exec(execDirS, fileMake, baseMake, metaMake, execMake, lookupData) // Execute
		if err != nil && !doit.ifPrintErrors("Prepare Exec:") {                        // abort?
			doit.can()
			return err
		}
	}

	err := doit.ctx.Err()
	doit.can()
	return err
}

func (doit *toDo) exec(
	execS DirS,
	tmplfile *Actor,
	executes *Actor,
	metadata *Actor,
	writeout *Actor,
	lookupData func(string) string,
) error {

	execS.flagPrint(ea, eav, "Exec:")

	// we'll recurse!

	for i := range execS {
		_ = execS[i].DirPath
		_ = execS[i].Recurse

		foldPile := NewNext(16, 4) // folders / subdirectories found (not used here)
		foldMatch := matchBool(true)
		var downS DirS
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
