// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"flag"
)

var (
	pm_, pmd, pmt, pmf, pmn, pma bool   // print main
	px_, pxa, pxd, pxt, pxl      bool   // print exec
	pw_, pwd, pwt, pwf, pwr      bool   // print write
	exe, exr, exf, nox           bool   // write
	tmplread                     string // TODO change default (& exf) acc to new 'type'-switch: md txt wiki ...
)

func init() {
	flag.BoolVar(&pm_, "a", false, "print all analysis info")
	flag.BoolVar(&pma, "ap", false, "print analysis path(s)")
	flag.BoolVar(&pmf, "af", false, "print analysis files")
	flag.BoolVar(&pmn, "an", false, "print analysis names")
	flag.BoolVar(&pmd, "ad", false, "print analysis datatree")
	flag.BoolVar(&pmt, "at", false, "print analysis template names"+"\n\t")

	flag.BoolVar(&px_, "e", false, "print all execution info")
	flag.BoolVar(&pxa, "ea", false, "print execution path")
	flag.BoolVar(&pxl, "el", false, "print execution line")
	flag.BoolVar(&pxd, "ed", false, "print execution datatree")
	flag.BoolVar(&pxt, "et", false, "print execution template names"+"\n\t")

	flag.BoolVar(&pw_, "w", false, "print all writing info")
	flag.BoolVar(&pwd, "wd", false, "print writing directories")
	flag.BoolVar(&pwr, "wr", false, "print raw unformatted text"+"\n\t")
	flag.BoolVar(&pwf, "wf", false, "print formatted text")

	flag.BoolVar(&exr, "ugo", false, "execute: write raw text (as *.ugo)")
	flag.BoolVar(&exe, "x", false, "execute: write resulting text")
	flag.BoolVar(&exf, "fmt", true, "apply go/format to raw text") // TODO fmt => nof & negate!
	flag.BoolVar(&pwt, "now", false, "print resulting text only - do not write"+"\n\t")
	flag.BoolVar(&nox, "nox", false, "skip execute, terminate after main analysis")

	flag.StringVar(&tmplread, "patterns", "*.go.tmpl;*.tmpl;dot.go.tmpl",
		"pattern list to match template file names"+
			"\tNote:"+
			"\n\tA template gets executed iff first pattern matches it's file!"+
			"\n\tA directory is writable iff last pattern matches some of it's file!"+
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
		pxa, pxd, pxl, pxt = true, true, true, true
		px_ = false
	} else if !(pxa || pxd || pxl || pxt) {
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
}

func flagArgs() []string {

	return flag.Args()
}
