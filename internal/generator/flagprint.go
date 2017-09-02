// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"
	"runtime"
	"time"
)

const (
	tab = "\t"
	cnt = "#"
	arr = "<-"
	don = " Done!"
	rec = "... recurse"
)

func dots(flag bool) string {
	switch {
	case flag:
		return rec
	default:
		return ""
	}
}

// flagOpen opens a phase and prints a prefix, iff flag is true
// Note: returned time.Now() shall be passed to corresponding flagClose(flag, time,Time)
func flagOpen(flag bool, prefix string) time.Time {
	runtime.GC()
	if flag {
		fmt.Print(prefix, tab, arr, tab)
	}
	return time.Now()
}

// flagDot indicates progress within a phase by printing a typical character, iff flag is true
func flagDot(flag bool, char string) {
	if flag {
		fmt.Print(char)
	}
}

// flagClose closes a phase, and prints the time elepsed since last flagOpen, iff flag is true
func flagClose(flag bool, start time.Time) {
	if flag {
		dur := time.Since(start)
		fmt.Println(tab, arr, don, tab, dur, tab)
	}
	runtime.GC()
}

// flagPrintByteS prints the byteS (as string), iff flag is true
func flagPrintByteS(flag bool, byteS []byte, header string) {
	if flag {
		fmt.Println(header, tab, cnt, len(byteS), tab, tab)

		fmt.Println(tab, string(byteS), tab, tab)

		fmt.Println(tab, tab, tab)
	}
}

// flagPrintString prints prefix & suffix, iff flag is true
func flagPrintString(flag bool, prefix string, suffix string) {
	if flag {
		fmt.Println(prefix, tab, arr, suffix, tab, tab)
	}
}

// flagPrint delegates to it
func (m Actor) flagPrint(flag, verbose bool, header string) {
	m.it.flagPrint(flag, verbose, header)
}

// flagPrint prints the dictionary, iff flag is true
func (d Dict) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header, tab, cnt, d.Len(), tab, tab)

		if verbose {
			do := func(item string) { flagPrintString(flag, "", item) }
			d.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints the path names, iff flag is true
func (d DirS) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header, tab, cnt, len(d), tab, tab)

		if verbose {
			for i := range d {
				flagPrintString(flag, d[i].DirPath, dots(d[i].Recurse))
			}
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints nothing but header, iff flag is true
func (n Null) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header, tab, cnt, n.Len(), tab, tab)
		if verbose {
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints the pile, iff flag is true
func (p nextPile) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header, tab, cnt, p.Len(), tab, tab)

		if verbose {
			do := func(item string) { fmt.Println(tab, item, tab, tab) }
			p.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints the pile, iff flag is true
func (p prevPile) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header, tab, cnt, p.Len(), tab, tab)

		if verbose {
			do := func(item string) { fmt.Println(tab, item, tab, tab) }
			p.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints the template, iff flag is true
func (tmpl Template) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header, tab, cnt, tmpl.Len(), tab, tab)

		if verbose {
			do := func(item string) { fmt.Println(tab, item, tab, tab) }
			tmpl.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
		}
	}
}

/*
// flagPrint  prints the data tree, iff flag is true
func (data Dot) flagPrint(flag bool, header string) {
	if flag {
		flagPrintDataTree(true, data, header)
		fmt.Println()
	}
}
*/

// ifPrintDataTree prints the data tree, iff flag is true
func (t *toDo) ifPrintDataTree(flag, verbose bool, header string) {
	if flag {
		itemS := t.data.S()
		count := len(itemS)
		fmt.Println(header, tab, cnt, count, tab, tab)

		if verbose {
			flagPrintDataTree(verbose, t.data, header)
			fmt.Println()
		}
	}
}

// ifPrintErrors prints the error(s), iff any
func (t *toDo) ifPrintErrors(header string) bool {
	return flagPrintErrors(t.data, header)
}
