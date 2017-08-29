// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"os"
	"path/filepath"
)

// dirS adopts the result of dotpath.DotPathS
type dirS []struct {
	DirPath string
	Recurse bool
}

func (d dirS) Close() error {
	return nil
}

func (d dirS) Walker(t *toDo, out ...Actor) func() {

	return func() {

		defer ActorsClose(out...)
		fh := func(path string, info os.FileInfo, err error) error {
			ActorsDo(path, out...)
			return nil
		}

		for i := 0; i < len(d) && t.ok(); i++ {
			dh := ifFlagSkipDirWf(matchBool(d[i].Recurse)) // Recurse?
			filepath.Walk(d[i].DirPath, t.isDirWf(dh, fh)) // Walk path
		}
	}
}

// asDirS - a helper for type conversion
func asDirS(i dirS) dirS {
	return i
}
