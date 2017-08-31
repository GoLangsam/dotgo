// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"
	"runtime"
	"time"
)

type dotter string

const (
	dotWalk dotter = "." // .
	dotFOut dotter = "-" // -
	dotTmpl dotter = "~" // ~
	dotData dotter = "'" // '

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
		fmt.Print(prefix + tab + arr + tab)
	}
	return time.Now()
}

// flagDot indicates progress within a phase by printing a typical character, iff flag is true
func flagDot(flag bool, char dotter) {
	if flag {
		fmt.Print(char)
	}
}

// flagClose closes a phase, and prints the time elepsed since last flagOpen, iff flag is true
func flagClose(flag bool, start time.Time) {
	if flag {
		dur := time.Since(start)
		fmt.Println(tab+arr+don+tab, dur, tab)
	}
	runtime.GC()
}

// flagPrintByteS prints the byteS (as string), iff flag is true
func flagPrintByteS(flag bool, byteS []byte, header string) {
	if flag {
		fmt.Println(header+tab+cnt, len(byteS), tab)

		fmt.Println(tab + string(byteS) + tab)

		fmt.Println()
	}
}

// flagPrintString prints prefix & suffix, iff flag is true
func flagPrintString(flag bool, prefix string, suffix string) {
	if flag {
		fmt.Println(prefix+tab+arr, suffix+tab)
	}
}

// flagPrint delegates to it
func (m Actor) flagPrint(flag, verbose bool, header string) {
	m.it.flagPrint(flag, verbose, header)
}

// flagPrint prints the dictionary, iff flag is true
func (d Dict) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header+tab+cnt, d.Len(), tab, tab)

		if verbose {
			for _, s := range d.S() {
				flagPrintString(flag, "", s)
			}
			fmt.Println()

		}
	}
}

// flagPrint prints the path names, iff flag is true
func (d DirS) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header+tab+cnt, len(d), tab, tab)

		if verbose {
			for i := range d {
				flagPrintString(flag, tab+d[i].DirPath, dots(d[i].Recurse))
			}
			fmt.Println()
		}
	}
}

// flagPrint prints nothing but header, iff flag is true
func (n Null) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header + tab)
		if verbose {
			fmt.Println()
		}
	}
}

// flagPrint prints the pile, iff flag is true
func (p nextPile) flagPrint(flag, verbose bool, header string) {
	if flag {
		itemS := <-p.Done()
		count := len(itemS)
		fmt.Println(header+tab+cnt, count, tab)

		if verbose {
			for i := range itemS {
				fmt.Println(tab + itemS[i] + tab)
			}
			fmt.Println()
		}
	}
}

// flagPrint prints the pile, iff flag is true
func (p prevPile) flagPrint(flag, verbose bool, header string) {
	if flag {
		itemS := <-p.Done()
		count := len(itemS)
		fmt.Println(header+tab+cnt, count, tab)

		if verbose {
			for i := count - 1; i >= 0; i-- {
				fmt.Println(tab + itemS[i] + tab)
			}
			fmt.Println()
		}
	}
}

// flagPrint prints the template, iff flag is true
func (tmpl Template) flagPrint(flag, verbose bool, header string) {
	if flag {
		itemS := tmpl.S()
		count := len(itemS)
		fmt.Println(header+tab+cnt, count, tab)

		if verbose {
			for i := range itemS {
				fmt.Println(tab + itemS[i] + tab)
			}
			fmt.Println()
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
func (t *toDo) ifPrintDataTree(flag bool, header string) {
	if flag {
		flagPrintDataTree(true, t.data, header)
		if true {
			fmt.Println()
		}
	}
}

// ifPrintErrors prints the error(s), iff any
func (t *toDo) ifPrintErrors(header string) bool {
	return flagPrintErrors(t.data, header)
}
