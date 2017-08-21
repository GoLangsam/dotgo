// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	"github.com/golangsam/dotgo/internal/args"
	f "github.com/golangsam/dotgo/internal/fs"  // adapter to file system (via "container/ccsafe/fs"
	a "github.com/golangsam/dotgo/internal/fsa" // adapter to file system analysis

	"github.com/golangsam/container/ccsafe/fs"
)

const (
	aDot = fs.Dot // just a "."
)

// DoIt performs it all:
//  - file system analysis (where to look)
//  - collection of metadata from templates
//  - execution of all relevant templates
func DoIt() {
	data := NewData(aDot)
	tmpl := NewTemplate(aDot)

	flagOpen(pm_, "analyse")
	args.IfPrintFlagArgs(pma, flagArgs()...)
	root := analyseFS(args.ToFolds(flagArgs()...)) // here the work is done

	f.IfPrintFsFileS(pmf, root.FsFileS, "File:")
	f.IfPrintFsBaseS(pmn, root.FsBaseS, "Base:")
	flagClose(pm_)

	todo := doIt(data, tmpl, root)
	todo.Collect()
	if !nox && !flagPrintErrors(data, "Prepare Main:") {
		todo.Execute()
	}
}

// These variables support the file system analysis.
var (
	extPatternS fs.PatternS // patterns for file system analysis
	execPattern *fs.Pattern // pattern for analysis := last pattern
)

func setPatterns(patterns ...string) {
	extPatternS = fs.NewPatternS(patterns...)
	execPattern = extPatternS[len(extPatternS)-1]
}

// Results from PrepareMain
func analyseFS(argS fs.FsFoldS) *a.Analysis {
	setPatterns(tmplread)
	root := a.Open(extPatternS) // Open
	last := len(argS) - 1
	for i, arg := range argS {
		if i < last {
			ca := root.AnalyseFlat(arg)                        // NotDown
			flagPrintAnalysisTree(false, ca, "Flat")           // ...flat
			root.FsFileS = append(root.FsFileS, ca.FsFileS...) // collect fileS
			root.FsBaseS = append(root.FsBaseS, ca.FsBaseS...) // collect baseS
		} else {
			fileS, baseS := root.FsFileS, root.FsBaseS // save fileS, baseS
			root = root.Reset()                        // hide fileS, baseS
			ca := root.AnalyseDown(arg)                // Descend
			flagPrintAnalysisTree(false, ca, "Down")   // ...down
			root.FsFileS, root.FsBaseS = fileS, baseS  // restore fileS, baseS
			// root.SubDirS = append(root.SubDirS, ca)    // attach sub tree
		}
		flagDot(pm_)
	}
	fsCache = a.Close() // Close & assign cache

	return root
}

func toNameS(files fs.FsBaseS) (nameS []f.FsName) {
	for i := range files { // convert to f.FsName
		nameS = append(nameS, files[i].BaseLessExt()) // less ext !
	}
	return nameS
}

func nameSnotEmpty(files fs.FsBaseS) (nameS []f.FsName) {
	nameS = toNameS(files)
	if len(nameS) < 1 { // no BaseNameS yet:
		nameS = append(nameS, forceBase(aDot)) // have "." as BaseName, at least
	}
	return nameS
}

func forceBase(name string) *fs.FsBase {
	return fs.ForceBase(name)
}

func flagPrintAnalysisTree(flag bool, ca *a.Analysis, prefix string) {
	if flag {
		if len(ca.FsFileS) > 0 {
			// flagPrintString(flag, "\t "+ca.String(), prefix+" Files:")
			f.IfPrintFsFileS(flag, ca.FsFileS, prefix)
		}
		if len(ca.FsBaseS) > 0 {
			// flagPrintString(flag, "\t "+ca.String(), prefix+" Names:")
			f.IfPrintFsBaseS(flag, ca.FsBaseS, prefix)
		}
		if len(ca.FsFileS) < 1 && len(ca.FsBaseS) < 1 {
			flagPrintString(flag, "----\t "+ca.String(), prefix)
		}
		for i := range ca.SubDirS {
			flagPrintAnalysisTree(flag, ca.SubDirS[i], prefix)
		}
		// println("===============================================================")
	}
}
