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

// ifPrintPile prints all items of the pile, iff flag is true
func (t *toDo) ifPrintPile(flag bool, pile *Pile, prefix string) {
	if flag && t.ok() {
		fmt.Println(prefix+tab+cnt, len(<-pile.Done()), tab)
		for item, ok := pile.Iter(); ok && t.ok(); item, ok = pile.Next() {
			fmt.Println(tab + item + tab)
		}
		fmt.Println()
	}
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

// flagPrintPathS prints the path names, iff flag is true
func flagPrintPathS(flag bool, pathS dirS, header string) {
	if flag {
		fmt.Println(header+tab+cnt, len(pathS), tab, tab)

		dots := func(flag bool) string {
			switch {
			case flag:
				return rec
			default:
				return ""
			}
		}
		for i := range pathS {
			flagPrintString(flag, tab+pathS[i].DirPath, dots(pathS[i].Recurse))
			// fmt.Println(tab + pathS[i].DirPath + tab + dots(pathS[i].Recurse))
		}
		fmt.Println()
	}
}

// ifPrintTemplate prints the template, iff flag is true
func (t *toDo) ifPrintTemplate(flag bool, prefix string) {
	if flag && t.ok() {
		flagPrintTemplate(true, t.tmpl, prefix)
		fmt.Println()
	}
}

// ifPrintDataTree prints the data tree, iff flag is true
func (t *toDo) ifPrintDataTree(flag bool, prefix string) {
	if flag && t.ok() {
		flagPrintDataTree(true, t.data, prefix)
		fmt.Println()
	}
}

// ifPrintErrors prints the error(s), iff any
func (t *toDo) ifPrintErrors(prefix string) bool {
	return flagPrintErrors(t.data, prefix)
}
