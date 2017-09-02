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
	pathS.flagPrint(aa, aav, "aa-Args:")

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
	_ = isTrue

	// some Null Actors - for flagDot dotter  // ...
	flagPrep := doer(func() { flagDot(a_, ".") })
	flagTemp := doer(func() { flagDot(a_, "-") })
	flagFile := doer(func() { flagDot(a_, "~") })
	flagMeta := doer(func() { flagDot(a_, "'") })

	show := func() Actor { return Actor{NewNull(), func(item string) { println("Debug:" + tab + item) }} }
	_ = show

	// Actors - Some containers, and how to populate each
	filePile := NewNext(512, 128).Action(isFile)                   // files (templates) to handle
	metaPile := NewPrev(256, 064).Action(hasMeta)                  // templates with non-empty meta: apply in reverse order!
	baseDict := NewDict().Action(isBase)                           // templates to execute: basenames found
	execDict := NewDict().Action(isExec)                           // TODO this is wrong: we need the directory! nameLessExt get's appended to base            // mathching file(s) identify folder(s) for execution
	rootData := NewData(aDot)                                      // data - a Dot
	rootTmpl := NewTemplate(aDot).tmplParser(rootData, lookupData) // text/template
	doit := doIt(rootData)                                         // carries context, and data

	// End of Prolog

	if doit.ctx.Err() == nil && len(prepDirS) > 0 { // Beg of prep Analysis
		analyse := flagOpen(a_, "Prepare:")
		prepDirS.flagPrint(ap, apv, "ap-Prep:")

		// a temp - for fan-out file names
		tempPile := NewNext(512, 128).Action(isFile)
		testPile := NewPrev(256, 064).Action(hasMeta) // templates with non-empty meta: apply in reverse order!

		quit := func() bool { return doit.ctx.Err() != nil }                   // quit, iff not doit.ok
		doit.do(prepDirS.Walker(quit, flagPrep, tempPile, filePile))           // go prepS => temp & file path
		doit.do(tempPile.Walker(quit, flagTemp, execDict, baseDict, testPile)) // go temp => exec & meta & test
		doit.do(testPile.Walker(quit))                                         // TODO go test => ./.
		doit.do(filePile.Walker(quit, flagFile, rootTmpl, metaPile))           // go file => rootTmpl & meta (*this* works!?!)
		doit.do(metaPile.Walker(quit, flagMeta))                               // go meta => drain
		doit.wg.Wait()                                                         // wait for all

		x := metaPile.S()
		y := testPile.S()
		if len(x) != len(y) {
			println("Sizes differ: meta =", len(x), "test =", len(y))
		}

		x = tempPile.S()
		y = filePile.S()
		if len(x) != len(y) {
			println("Sizes differ: temp =", len(x), "file =", len(y))
		}

		tmpl, err := rootTmpl.it.(Template).Clone()                 // Clone rootTmpl
		rootData.SeeError("Clone", "Root", err)                     // err? ignore for now
		metaTmpl := Template{tmpl}.metaParser(rootData, lookupData) // text/template from meta
		doit.do(metaPile.Walker(quit, flagMeta, metaTmpl))          // go meta => metaTmpl & metaData
		doit.wg.Wait()                                              // wait for all
		tempPile = NewNext(0, 0).Action(isTrue)                     // TODO forget temp

		filePile.flagPrint(af, afv, "af-File:")
		metaPile.flagPrint(am, amv, "am-Meta:")
		baseDict.flagPrint(an, anv, "an-Name:")
		rootTmpl.flagPrint(ar, arv, "ar-Root:")
		metaTmpl.flagPrint(at, atv, "at-Data:")
		doit.ifPrintDataTree(ad, adv, "ad-"+aDot+aDot+aDot+aDot+":")

		doit.data = NewData(aDot) // forget
		flagClose(a_, analyse)
	} // End of prep Analysis

	if doit.ctx.Err() == nil && !nox && !doit.ifPrintErrors("Prepare Main:") {
		//	todo.Execute()
		err := doit.exec(execDirS, filePile, baseDict, metaPile, execDict, lookupData) // Execute
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
