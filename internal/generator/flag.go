// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"flag"
)

var (
	a_, ad, at, af, an, ap  bool   // print analysis
	adv, atv, afv, anv, apv bool   // ...verbose
	e_, ea, ed, et, el      bool   // print exec
	eav, edv, etv, elv      bool   // ...verbose
	w_, wd, wf, wr          bool   // print write
	wdv, wfv, wrv           bool   // ...verbose
	exe, ugo, nof, nox, nos bool   // write
	tmplread                string // TODO change default (& exf) acc to new 'type'-switch: md txt wiki ...
)

func init() {
	const new = "\n"
	const lin = "----------"
	const sep = new + lin + lin + lin + lin + lin + lin + lin + lin + lin + lin + new

	flag.BoolVar(&a_, "a", false, "print all analysis info:")
	flag.BoolVar(&ap, "ap", false, "print analysis path(s)")
	flag.BoolVar(&af, "af", false, "print analysis files")
	flag.BoolVar(&an, "an", false, "print analysis names")
	flag.BoolVar(&ad, "ad", false, "print analysis datatree")
	flag.BoolVar(&at, "at", false, "print analysis template names")

	flag.BoolVar(&apv, "apv", false, "...verbose")
	flag.BoolVar(&afv, "afv", false, "...verbose")
	flag.BoolVar(&anv, "anv", false, "...verbose")
	flag.BoolVar(&adv, "adv", false, "...verbose")
	flag.BoolVar(&atv, "atv", false, "...verbose"+sep)

	flag.BoolVar(&e_, "e", false, "print all execution info:")
	flag.BoolVar(&ea, "ea", false, "print execution path")
	flag.BoolVar(&el, "el", false, "print execution line")
	flag.BoolVar(&ed, "ed", false, "print execution datatree")
	flag.BoolVar(&et, "et", false, "print execution template names")

	flag.BoolVar(&eav, "eav", false, "...verbose")
	flag.BoolVar(&elv, "elv", false, "...verbose")
	flag.BoolVar(&edv, "edv", false, "...verbose")
	flag.BoolVar(&etv, "etv", false, "...verbose"+sep)

	flag.BoolVar(&w_, "w", false, "print all writing info:")
	flag.BoolVar(&wd, "wd", false, "print directories")
	flag.BoolVar(&wf, "wf", false, "print formatted text")
	flag.BoolVar(&wr, "wr", false, "print raw unformatted text")

	flag.BoolVar(&wdv, "wdv", false, "...verbose")
	flag.BoolVar(&wfv, "wfv", false, "...verbose")
	flag.BoolVar(&wrv, "wrv", false, "...verbose"+sep)

	flag.BoolVar(&nof, "nof", true, "no formatting - do not apply go/format to raw text")
	flag.BoolVar(&nos, "nos", false, "no safe - print resulting text only")
	flag.BoolVar(&nox, "nox", false, "no execute - terminate after main analysis"+sep)

	flag.BoolVar(&exe, "x", false, "execute: safe resulting text")
	flag.BoolVar(&ugo, "xgo", false, "execute: safe raw text (as *.ugo)"+sep)

	flag.StringVar(&tmplread, "ext", "*.go.tmpl;*.tmpl;dot.go.tmpl",
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

	if adv {
		ad = true
	}
	if atv {
		at = true
	}
	if afv {
		af = true
	}
	if anv {
		an = true
	}
	if apv {
		ap = true
	}

	if eav {
		ea = true
	}
	if edv {
		ed = true
	}
	if etv {
		et = true
	}
	if elv {
		el = true
	}

	if wdv {
		wd = true
	}

	if wfv {
		wf = true
	}
	if wrv {
		wr = true
	}

	if a_ {
		ap, af, an, ad, at = true, true, true, true, true
		a_ = false
	} else if !(ap || af || an || ad || at) {
		a_ = true
	}

	if e_ {
		ea, ed, el, et = true, true, true, true
		e_ = false
	} else if !(ea || ed || el || et) {
		e_ = true
	}

	if w_ {
		wd, wf, wr = true, true, true
		w_ = false
	} else if !(wd || wf || wr) {
		w_ = true
	}

	if nos {
		exe = false
	}
}

func flagArgs() []string {

	return flag.Args()
}
