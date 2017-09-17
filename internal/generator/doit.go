// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package gen implements the agnostic template driven text & source generator
package gen

import (
	"github.com/GoLangsam/container/ccsafe/dotpath"
	"github.com/GoLangsam/container/ccsafe/fileinfocache"
)

// DoIt performs it all:
//  - establish the participants
//  - walk the file system (where to look)
//  - collect metadata
//  - execute relevant templates
// TODO use an adaptive cache instead of fic
// TODO dirS shall be absolute, if we intend to move os.Getwd during Exec
func DoIt() error {

	// Beg of Prolog

	pathS := AsDirS(dotpath.DotPathS(flagArgs()...))
	pathS.flagPrint(aa, aav, "aa-Args:")

	split := len(pathS) - 1           // at last:
	prepDirS := AsDirS(pathS[:split]) // - prepare all but last
	execDirS := AsDirS(pathS[split:]) // - execute only last

	fsCache := fic.New()
	lookupData := func(path string) string { return fsCache.LookupData(path) }

	prep := NewStep(lookupData)
	defer prep.todo.can()

	// End of Prolog

	if !prep.done() && len(prepDirS) > 0 {
		analyse := flagOpen(a_, "Prepare")
		prep = prep.prepDo(prepDirS).prepPrint()
		flagClose(a_, analyse)
		if all.NotOk("Prepare Main:") {
			return prep.todo.ctx.Err() // abort
		}
	}

	if !prep.done() {
		execute := flagOpen(e_, "Execute")
		prep = prep.execDo(execDirS)
		flagClose(e_, execute)
		if all.NotOk("Prepare Exec:") {
			return prep.todo.ctx.Err() // abort
		}
	}

	prep.wg.Wait() // wait for all
	return prep.todo.ctx.Err()
}
