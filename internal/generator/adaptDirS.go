// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
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

func (d DirS) S() []string {
	s := []string{}
	for i := range d {
		s = append(s, d[i].DirPath+tab+dots(d[i].Recurse)+tab)
	}
	return s
}

func (d DirS) Close() error {
	return nil
}

func (d DirS) Walker(quit func() bool, out ...Actor) func() {

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

// AsDirS - a helper for type conversion
func AsDirS(i DirS) DirS {
	return i
}
