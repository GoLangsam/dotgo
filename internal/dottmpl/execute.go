// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	a "github.com/golangsam/dotgo/internal/fsa" // adapter to file system analysis
)

// Execute - given the current (= last) directory
// CollectDown
// while carrying known names of templates down.
func (t toDo) Execute() {
	//	ToDo.Assign(Data, namS...)

	flagOpen(px_, "execute")

	for i := range t.dir.SubDirS { // should be just one
		t.dir.SubDirS[i].FsBaseS = append(t.dir.FsBaseS, t.dir.SubDirS[i].FsBaseS...)
		a.IfPrintAnalysis(pwd, "<- Target Directory", t.dir.SubDirS[i])
		flagPrintAnalysisTree(pwd, t.dir.SubDirS[i], "<- Target")
		if todo, ok := t.doIt(t.data, t.tmpl, t.dir.SubDirS[i]); ok {
			todo.CollectDown()
		}
	}

	flagClose(px_)
}

// CollectDown - CollectFold iff there are matching files,
// and recurse CollectDown into all subdirs
// while carrying known names of files and templates down.
func (t toDo) CollectDown() {
	if len(t.dir.MatchingFileS(t.dir.FsFileS, execPattern)) > 0 { // # have execPattern matches
		t.CollectFold()
	}

	for i := range t.dir.SubDirS {
		t.dir.SubDirS[i].FsFileS = append(t.dir.FsFileS, t.dir.SubDirS[i].FsFileS...)
		t.dir.SubDirS[i].FsBaseS = append(t.dir.FsBaseS, t.dir.SubDirS[i].FsBaseS...)
		if todo, ok := t.doIt(t.data, t.tmpl, t.dir.SubDirS[i]); ok {
			todo.CollectDown()
		}
	}
}

// CollectFold - using a clone of the template,
// collect data from each target template
// of the current folder
// and ExecuteTmpl with it's (shortened) name
func (t toDo) CollectFold() {
	exec, err := t.tmpl.Clone()
	if !SeeError(t.data, err, "Clone Main:") {
		fold := t.data.G(t.dir.String())
		todo, ok := t.doIt(fold, exec, t.dir)
		if ok {
			flagPrintString(pxl, todo.dir.String(), "Directory")
			todo.CollectTmplS()
			flagPrintTemplate(pxt, todo.tmpl, todo.dir.String()+" - Execution")

			names := nameSnotEmpty(todo.dir.FsBaseS)
			for i := range names {
				name := names[i].String()
				file := fold.G(name)
				if todo, ok := todo.doIt(file, todo.tmpl, todo.dir); ok {
					todo.CollectDataS()
					flagPrintDataTree(pxd, todo.data, todo.dir.String())
					if !flagPrintErrors(file, "Collect "+t.dir.String()+" for name "+name+":") {
						dotS := file.DownS()
						for i := range dotS {
							if todo, ok := todo.doIt(dotS[i], todo.tmpl, todo.dir); ok {
								todo.ExecuteTmpl(name)
							}
						}
					}
				}
			}
			flagDot(px_)
		}
	}
}

// ExecuteTempl - apply template 'name' and show/format/write result
func (t toDo) ExecuteTmpl(name string) {
	var err error
	flagPrintString(pxl, "Apply", t.data.String()+"\t<- "+name+"\t")
	byteS, err := Apply(t.data, t.tmpl, name)
	if !SeeError(t.data, err, "Execute") {
		flagPrintByteS(pwr, byteS, ">>>>Raw text of "+name+" & "+t.data.String())
		if exr {
			filename := t.dir.JoinWith(FileName(t.data, name+".ugo"))
			SeeError(t.data, Write(t.data, filename, byteS), "Write Raw")
		}
		if exf { // if strings.HasSuffix(fileName, ".go") {
			byteS, err = Source(byteS)
			SeeError(t.data, err, "Format")
		}
		flagPrintByteS(pwt || pwf, byteS, ">>>>Final text of "+name+" & "+t.data.String())
		filename := t.dir.JoinWith(FileName(t.data, name))
		if exe {
			SeeError(t.data, Write(t.data, filename, byteS), "Write")
		}
	}
}
