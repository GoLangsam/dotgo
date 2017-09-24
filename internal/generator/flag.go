// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"flag"
)

var (
	a_, aa                            bool   // print analysis
	ap, af, an, ad, am, ar, at        bool   // p f n d m r t
	aav                               bool   // ...verbose
	apv, afv, anv, adv, amv, arv, atv bool   // p f n d m r t
	e_, el                            bool   // print exec
	ep, ef, en, ed, em, er, et        bool   // p f n d m r t
	elv                               bool   // ...verbose
	epv, efv, env, edv, emv, erv, etv bool   // p f n d m r t
	w_                                bool   // print write
	wd, wf, wr                        bool   // dir file raw
	wdv, wfv, wrv                     bool   // ...verbose
	exe, nof, nox, nos, ugo           bool   // write no-format no-exec no-save ugly-go
	tmplread                          string // TODO change default (& exf) acc to new 'type'-switch: md txt wiki ...
)

func init() {
	const eol = "\n"
	const lin = "----------"
	const sep = eol + lin + lin + lin + lin + lin + lin + lin + lin + lin + lin + eol

	flag.BoolVar(&a_, "a", false, "print all analysis info:")
	flag.BoolVar(&aa, "aa", false, "print analysis arguments(s)")
	// p f n d m r t
	flag.BoolVar(&ap, "ap", false, "print analysis path(s)")
	flag.BoolVar(&af, "af", false, "print analysis files")
	flag.BoolVar(&an, "an", false, "print analysis names")
	flag.BoolVar(&ad, "ad", false, "print analysis datatree")
	flag.BoolVar(&am, "am", false, "print analysis meta files")
	flag.BoolVar(&ar, "ar", false, "print analysis root template names")
	flag.BoolVar(&at, "at", false, "print analysis meta template names")

	flag.BoolVar(&aav, "aav", false, "...verbose")
	// p f n d m r t
	flag.BoolVar(&apv, "apv", false, "...verbose")
	flag.BoolVar(&afv, "afv", false, "...verbose")
	flag.BoolVar(&anv, "anv", false, "...verbose")
	flag.BoolVar(&adv, "adv", false, "...verbose")
	flag.BoolVar(&amv, "amv", false, "...verbose")
	flag.BoolVar(&atv, "arv", false, "...verbose")
	flag.BoolVar(&atv, "atv", false, "...verbose"+sep)

	flag.BoolVar(&e_, "e", false, "print all execution info:")
	flag.BoolVar(&el, "el", false, "print execution line")
	// p f n d m r t
	flag.BoolVar(&ep, "ep", false, "print execution path(s)")
	flag.BoolVar(&ef, "ef", false, "print execution files")
	flag.BoolVar(&en, "en", false, "print execution names")
	flag.BoolVar(&ed, "ed", false, "print execution datatree")
	flag.BoolVar(&em, "em", false, "print execution meta files")
	flag.BoolVar(&er, "er", false, "print execution root template names")
	flag.BoolVar(&et, "et", false, "print execution meta template names")

	flag.BoolVar(&elv, "elv", false, "...verbose")
	// p f n d m r t
	flag.BoolVar(&epv, "epv", false, "...verbose")
	flag.BoolVar(&efv, "efv", false, "...verbose")
	flag.BoolVar(&env, "env", false, "...verbose")
	flag.BoolVar(&edv, "edv", false, "...verbose")
	flag.BoolVar(&emv, "emv", false, "...verbose")
	flag.BoolVar(&erv, "erv", false, "...verbose")
	flag.BoolVar(&etv, "etv", false, "...verbose"+sep)

	flag.BoolVar(&w_, "w", false, "print all writing info:")
	flag.BoolVar(&wd, "wd", false, "print directories")
	flag.BoolVar(&wf, "wf", false, "print formatted text")
	flag.BoolVar(&wr, "wr", false, "print raw unformatted text")

	flag.BoolVar(&wdv, "wdv", false, "...verbose")
	flag.BoolVar(&wfv, "wfv", false, "...verbose")
	flag.BoolVar(&wrv, "wrv", false, "...verbose"+sep)

	flag.BoolVar(&nof, "nof", false, "no formatting - do not apply go/format to raw text")
	flag.BoolVar(&nos, "nos", false, "no save - print resulting text only")
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
	// p f n d m r t
	if apv {
		ap = true
	}
	if afv {
		af = true
	}
	if anv {
		an = true
	}
	if adv {
		ad = true
	}
	if amv {
		am = true
	}
	if arv {
		ar = true
	}
	if atv {
		at = true
	}

	if elv {
		el = true
	}
	// p f n d m r t
	if epv {
		ep = true
	}
	if efv {
		ef = true
	}
	if env {
		en = true
	}
	if edv {
		ed = true
	}
	if emv {
		em = true
	}
	if erv {
		et = true
	}
	if erv {
		et = true
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
		aa = true
		// p f n d m r t
		ap, af, an, ad, am, ar, at = true, true, true, true, true, true, true
		a_ = false
	} else if !(aa ||
		// p f n d m r t
		ap || af || an || ad || am || ar || at) {
		a_ = true
	}

	if e_ {
		el = true
		// p f n d m r t
		ep, ef, en, ed, em, er, et = true, true, true, true, true, true, true
		e_ = false
	} else if !(el ||
		// p f n d m r t
		ep || ef || en || ed || em || er || et) {
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
