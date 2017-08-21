// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	a "github.com/golangsam/dotgo/internal/fsa" // adapter to file system analysis
)

func (t toDo) Execute() {
	//	ToDo.Assign(Data, namS...)

	flagOpen(px_, "execute")

	for i := range t.dir.SubDirS { // should be just one
		t.dir.SubDirS[i].FsBaseS = append(t.dir.FsBaseS, t.dir.SubDirS[i].FsBaseS...)
		a.IfPrintAnalysis(pwd, "<- Target Directory", t.dir.SubDirS[i])
		flagPrintAnalysisTree(pwd, t.dir.SubDirS[i], "<- Target")
		todo := doIt(t.data, t.tmpl, t.dir.SubDirS[i])
		todo.CollectDown()
	}

	flagClose(px_)
}

func (t toDo) CollectDown() {
	if len(t.dir.MatchingFileS(t.dir.FsFileS, execPattern)) > 0 { // # have execPattern matches
		t.CollectFold()
	}

	for i := range t.dir.SubDirS {
		t.dir.SubDirS[i].FsFileS = append(t.dir.FsFileS, t.dir.SubDirS[i].FsFileS...)
		t.dir.SubDirS[i].FsBaseS = append(t.dir.FsBaseS, t.dir.SubDirS[i].FsBaseS...)
		todo := doIt(t.data, t.tmpl, t.dir.SubDirS[i])
		todo.CollectDown()
	}
}

func (t toDo) CollectFold() {
	exec, err := t.tmpl.Clone()
	if !SeeError(t.data, err, "Clone Main:") {
		fold := t.data.G(t.dir.String())
		todo := doIt(fold, exec, t.dir)

		flagPrintString(pxl, todo.dir.String(), "Directory")
		todo.CollectTmplS()
		flagPrintTemplate(pxt, todo.tmpl, todo.dir.String()+" - Execution")

		names := nameSnotEmpty(todo.dir.FsBaseS)
		for i := range names {
			name := names[i].String()
			file := fold.G(name)
			todo := doIt(file, todo.tmpl, todo.dir)
			todo.CollectDataS()
			flagPrintDataTree(pxd, todo.data, todo.dir.String())
			if !flagPrintErrors(file, "Collect "+t.dir.String()+" for name "+name+":") {
				dotS := file.DownS()
				for i := range dotS {
					todo := doIt(dotS[i], todo.tmpl, todo.dir)
					todo.ExecuteTmpl(name)
				}
			}
			flagDot(px_)
		}
	}
	// flagPrintDataTree(pxd, t.data, t.dir.String())
}

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
