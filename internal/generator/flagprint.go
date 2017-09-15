// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var mu sync.Mutex

const (
	tab = "\t"
	cnt = "#"
	arr = "<-"
	don = "Done!"
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
	mu.Lock()
	defer mu.Unlock()
	runtime.GC()
	if flag {
		fmt.Println(tab, tab, tab)
		fmt.Print(prefix, tab, arr, tab, tab)
	}
	return time.Now()
}

// flagDot indicates progress within a phase by printing a typical character, iff flag is true
func flagDot(flag bool, char string) {
	if flag {
		// mu.Lock()
		// defer mu.Unlock()
		fmt.Print(char)
	}
}

// flagClose closes a phase, and prints the time elepsed since last flagOpen, iff flag is true
func flagClose(flag bool, start time.Time) {
	if flag {
		mu.Lock()
		defer mu.Unlock()
		dur := time.Since(start)
		fmt.Println(tab, tab, tab)
		fmt.Println("", tab, arr, don, tab, dur, tab)
	}
	runtime.GC()
}

// ======

// flagPrintByteS prints the byteS (as string), iff flag is true
func flagPrintByteS(flag bool, byteS []byte, header string) {
	if flag {
		mu.Lock()
		defer mu.Unlock()
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

// ====== implement Some

// flagPrint -
// delegate to it
func (m Actor) flagPrint(flag, verbose bool, header string) {
	m.it.flagPrint(flag, verbose, header)
}

// flagPrint prints
// the dictionary,
// iff flag is true
func (d Dict) flagPrint(flag, verbose bool, header string) {
	if flag {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println(header, tab, cnt, d.Len(), tab, tab)

		if verbose {
			do := func(item string) { flagPrintString(flag, arr, item) }
			d.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints
// the path names,
// iff flag is true
func (d DirS) flagPrint(flag, verbose bool, header string) {
	if flag {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println(header, tab, cnt, len(d), tab, tab)

		if verbose {
			for i := range d { // do not use Walker as we have two items
				flagPrintString(flag, d[i].DirPath, dots(d[i].Recurse))
			}
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints
// the data tree,
// iff flag is true
func (d Data) flagPrint(flag, verbose bool, header string) {
	if flag {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println(header, tab, cnt, d.Len(), tab, tab)

		if verbose {
			d.PrintTree(">>")
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints
// nothing but header,
// iff flag is true
func (n Null) flagPrint(flag, verbose bool, header string) {
	if flag {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println(header, tab, cnt, n.Len(), tab, tab)

		if verbose {
			do := func(item string) {}
			n.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints
// the pile,
// iff flag is true
func (p NextPile) flagPrint(flag, verbose bool, header string) {
	if flag {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println(header, tab, cnt, p.Len(), tab, tab)

		if verbose {
			do := func(item string) { fmt.Println(tab, item, tab, tab) }
			p.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints
// the pile (in reverse order),
// iff flag is true
func (p PrevPile) flagPrint(flag, verbose bool, header string) {
	if flag {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println(header, tab, cnt, p.Len(), tab, tab)

		if verbose {
			do := func(item string) { fmt.Println(tab, item, tab, tab) }
			p.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
		}
	}
}

// flagPrint prints
// the template,
// iff flag is true
func (template Template) flagPrint(flag, verbose bool, header string) {
	if flag {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println(header, tab, cnt, template.Len(), tab, tab)

		if verbose {
			do := func(item string) { fmt.Println(tab, item, tab, tab) }
			template.Walker(noquit, doit(do))()
			fmt.Println(tab, tab, tab)
		}
	}
}
