// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"flag"
)

var (
	a_, aa, ad, ar, at, af, an, am, ap     bool   // print analysis
	aav, adv, arv, atv, afv, anv, amv, apv bool   // ...verbose
	e_, ea, ed, et, el                     bool   // print exec
	eav, edv, etv, elv                     bool   // ...verbose
	w_, wd, wf, wr                         bool   // print write
	wdv, wfv, wrv                          bool   // ...verbose
	exe, ugo, nof, nox, nos                bool   // write
	tmplread                               string // TODO change default (& exf) acc to new 'type'-switch: md txt wiki ...
)

func init() {
	const eol = "\n"
	const lin = "----------"
	const sep = eol + lin + lin + lin + lin + lin + lin + lin + lin + lin + lin + eol

	flag.BoolVar(&a_, "a", false, "print all analysis info:")
	flag.BoolVar(&aa, "aa", false, "print analysis arguments(s)")
	flag.BoolVar(&ap, "ap", false, "print analysis path(s)")
	flag.BoolVar(&af, "af", false, "print analysis files")
	flag.BoolVar(&an, "an", false, "print analysis names")
	flag.BoolVar(&am, "am", false, "print analysis meta files")
	flag.BoolVar(&ad, "ad", false, "print analysis datatree")
	flag.BoolVar(&ar, "ar", false, "print analysis root template names")
	flag.BoolVar(&at, "at", false, "print analysis meta template names")

	flag.BoolVar(&aav, "aav", false, "...verbose")
	flag.BoolVar(&apv, "apv", false, "...verbose")
	flag.BoolVar(&afv, "afv", false, "...verbose")
	flag.BoolVar(&amv, "amv", false, "...verbose")
	flag.BoolVar(&anv, "anv", false, "...verbose")
	flag.BoolVar(&adv, "adv", false, "...verbose")
	flag.BoolVar(&atv, "arv", false, "...verbose")
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

	flag.BoolVar(&nof, "nof", false, "no formatting - do not apply go/format to raw text")
	flag.BoolVar(&nos, "nos", false, "no safe - print resulting text only")
	flag.BoolVar(&nox, "nox", false, "no execute - terminate after main analysis"+sep)

	flag.BoolVar(&exe, "x", false, "execute: safe resulting text")
	flag.BoolVar(&ugo, "xgo", false, "execute: safe raw text (as *.ugo)"+sep)

	flag.StringVar(&tmplread, "ext", ".go.tmpl;.tmpl;dot.go.tmpl",
		"extension list to match template file names"+
			"\tNote:"+
			"\n\tA template gets executed iff it's file name matches the first extension!"+
			"\n\tA directory is writable iff some of it's files name matches the last extension!"+
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

	if aav {
		aa = true
	}
	if adv {
		ad = true
	}
	if arv {
		ar = true
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
	if amv {
		am = true
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
		aa, ap, af, an, am, ad, ar, at = true, true, true, true, true, true, true, true
		a_ = false
	} else if !(aa || ap || af || an || am || ad || ar || at) {
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
