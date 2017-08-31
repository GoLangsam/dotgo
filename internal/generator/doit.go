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

	// lookupData retrieves the data related to file from cache
	lookupData := func(path string) string {
		return fsCache.LookupData(path)
	}

	suffix := filepath.SplitList(tmplread)
	isFile := matchFunc(suffix...)
	isBase := matchFunc(suffix[0])
	isExec := matchFunc(suffix[len(suffix)-1])
	//isTrue := matchBool(true)

	filePile := NewNext(512, 128) // files (templates) to handle
	//metaPile := NewPrev(256, 64)     // templates with non-empty meta: apply in reverse order!
	metaPile := NewPrev(256, 0)      // templates with non-empty meta: apply in reverse order!
	baseDict := NewDict()            // templates to execute: basenames found
	execDict := NewDict()            // mathching file(s) identify folder(s) for execution
	rootData := NewData(aDot)        // data - a Dot
	rootTmpl := NewTemplate(aDot)    // text/template
	doit := doIt(rootData, rootTmpl) // carries context, and data & tmpl

	// Actors - how to populate each Container
	fileMake := Actor{filePile, func(item string) {
		if isFile(item) {
			filePile.Pile(item)
		}
	}}

	metaMake := Actor{metaPile, func(item string) {
		println("\nCheck Meta " + item)
		meta, err := Meta(lookupData(item))
		if err == nil && meta != "" {
			metaPile.Pile(item)
			println("\nFound Meta " + item)
		} else if err != nil {
			panic(err)
		}
	}}

	baseMake := Actor{baseDict, func(item string) {
		if isBase(item) {
			baseDict.Assign(nameLessExt(item), nil)
		}
	}}

	execMake := Actor{execDict, func(item string) {
		if isExec(item) {
			execDict.Assign(nameLessExt(item), nil) // TODO this is wrong: we need the directory! nameLessExt get's appended to base
		}
	}}

	// some Null Actors - for flagDot dotter
	flagWalk := Actor{NewNull(), func(item string) {
		flagDot(a_, dotWalk) // ...
	}}
	flagFOut := Actor{NewNull(), func(item string) {
		flagDot(a_, dotFOut) // ...
		println(item)
	}}
	flagTmpl := Actor{NewNull(), func(item string) {
		flagDot(a_, dotTmpl) // ...
	}}
	flagData := Actor{NewNull(), func(item string) {
		flagDot(a_, dotData) // ...
	}}

	// End of Prolog

	if doit.ok() && len(prepDirS) > 0 { // Beg of prep Analysis
		analyse := flagOpen(a_, "Prepare:")
		prepDirS.flagPrint(ap, ap, "Prep:")

		// a temp Pile - fan out file names
		tempPile := NewNext(128, 32)
		tempMake := Actor{tempPile, func(item string) {
			tempPile.Pile(item)
		}}

		tmplParse := Actor{tempPile, tmplParser(doit, lookupData)} // => doit.tmpl
		metaParse := Actor{metaPile, metaParser(doit, lookupData)} // => doit.tmpl

		doit.do(prepDirS.Walker(doit, flagWalk, tempMake, fileMake)) // go prepS => temp & file path
		doit.do(tempMake.Walker(doit, flagTmpl, tmplParse))          // go temp => doit.tmpl
		doit.do(fileMake.Walker(doit, flagFOut, metaMake, baseMake)) // go path => meta & base
		doit.do(metaMake.Walker(doit, flagData, metaParse))          // go meta => drain
		doit.wg.Wait()                                               // wait for all

		tempPile = NewNext(0, 0)                    // forget
		tempMake = Actor{tempPile, func(string) {}} // forget

		doit.tmpl.flagPrint(at, atv, "Main:")
		doit.ifPrintDataTree(ad, aDot)
		filePile.flagPrint(af, afv, "File:")
		metaPile.flagPrint(am, amv, "Meta:")
		baseDict.flagPrint(an, anv, "Base:")

		doit.data = NewData(aDot) // forget
		flagClose(a_, analyse)
	} // End of prep Analysis

	if doit.ok() && !nox && !doit.ifPrintErrors("Prepare Main:") {
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
	tmplfile Actor,
	executes Actor,
	metadata Actor,
	writeout Actor,
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
