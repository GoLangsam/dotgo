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

// ifPrintErrors prints the error(s), iff any
func (t *toDo) ifPrintErrors(header string) bool {
	return flagPrintErrors(t.data, header)
}
