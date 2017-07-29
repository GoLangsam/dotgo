// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	"flag"
)

var (
	pm_, pmd, pmt, pmf, pmn, pma bool // print main
	px_, pxd, pxt, pxl           bool // print exec
	pw_, pwd, pwt, pwf, pwr      bool // print write
	exe, exr, exf, nox           bool // write
	seq                          bool // sequential processing
	tmplread                     string
)

func init() {
	flag.BoolVar(&pm_, "pm", false, "print all main info")
	flag.BoolVar(&pma, "pma", false, "print main args")
	flag.BoolVar(&pmf, "pmf", false, "print main files")
	flag.BoolVar(&pmn, "pmn", false, "print main names")
	flag.BoolVar(&pmd, "pmd", false, "print main datatree")
	flag.BoolVar(&pmt, "pmt", false, "print main template names")

	flag.BoolVar(&px_, "px", false, "print all execution info")
	flag.BoolVar(&pxl, "pxl", false, "print execution line")
	flag.BoolVar(&pxd, "pxd", false, "print execution datatree")
	flag.BoolVar(&pxt, "pxt", false, "print execution template names")

	flag.BoolVar(&pw_, "pw", false, "print all writing info")
	flag.BoolVar(&pwd, "pwd", false, "print writing directories")
	flag.BoolVar(&pwr, "pwr", false, "print raw unformatted text")
	flag.BoolVar(&pwf, "pwf", false, "print formatted text")

	flag.BoolVar(&pwt, "p", false, "print resulting text only - do not write")

	flag.BoolVar(&exf, "fmt", true, "apply go/format to raw text")
	flag.BoolVar(&exr, "ugo", false, "execute: write raw text (as *.ugo)")
	flag.BoolVar(&exe, "x", false, "execute: write resulting text")
	flag.BoolVar(&seq, "seq", false, "sequential execution - do not spawn go routines")
	flag.BoolVar(&nox, "xno", false, "skip execute, terminate after main analysis")

	flag.StringVar(&tmplread, "tmplread", "*.go.tmpl;*.tmpl;dot.go.tmpl",
		"pattern(s) to match templates (files & directories!)"+
			"\n"+
			"\tNotes:"+
			"\n"+
			"\tA template is executable if first pattern matches it's file!"+
			"\n"+
			"\tA directory is writable if last pattern matches some of it's file!"+
			"\n\t")

	/*
		flag.StringVar(&skipParse, "skipParse", "",
			"No template is parsed from matching basename")
		flag.StringVar(&skipMeta, "skipMeta", "",
			"No meta info is extracted from matching basename")
	*/

	flagParse()
}

func flagParse() {
	flag.Parse()
	if pm_ {
		pma, pmf, pmn, pmd, pmt = true, true, true, true, true
		pm_ = false
	} else if !(pma || pmf || pmn || pmd || pmt) {
		pm_ = true
	}

	if px_ {
		pxd, pxl, pxt = true, true, true
		px_ = false
	} else if !(pxd || pxl || pxt) {
		px_ = true
	}

	if pw_ {
		pwd, pwf, pwr = true, true, true
		pw_ = false
	} else if !(pwd || pwf || pwr) {
		pw_ = true
	}

	if pwt {
		exe = false
	}

	// seq = true // TODO
}

func flagArgs() []string {
	return flag.Args()
}
