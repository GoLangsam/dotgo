// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"
	"os"
	"path/filepath"
)

// DirS adopts the result of dotpath.DotPathS
//
// Hint: use AsDirS for type conversion
type DirS []struct {
	DirPath string
	Recurse bool
}

// Beg implement Some

// S -
// return all DirPath & dots(Recurse)
func (d DirS) S() []string {
	s := []string{}
	for i := range d {
		s = append(s, d[i].DirPath+tab+dots(d[i].Recurse)+tab)
	}
	return s
}

// Len -
// how many directories
func (d DirS) Len() int {
	return len(d)
}

// Close -
// pretend to be a Closer (<=> an io.Closer)
func (d DirS) Close() error {
	return nil
}

// Walker -
// traverse each DirPath
// apply ActorsDo to each file,
// and descend into its sub-directories
// iff d[i].Recurse
// Note: uses filepath.Walk
func (d DirS) Walker(quit func() bool, out ...*Actor) func() {

	return func() {

		defer ActorsClose(out...)
		fh := func(path string, info os.FileInfo, err error) error {
			ActorsDo(path, out...)
			return nil
		}

		for i := 0; i < len(d) && !quit(); i++ {
			dh := ifFlagSkipDirWf(matchBool(d[i].Recurse))     // Recurse?
			filepath.Walk(d[i].DirPath, isDirWf(quit, dh, fh)) // Walk path
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

// End implement Some

// AsDirS - a helper for type conversion
func AsDirS(i DirS) DirS {
	return i
}
