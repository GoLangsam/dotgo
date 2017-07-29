// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsa

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// AnalyseFlat returns the analysis rooted at *Fold without recursing into any subdirs
func (a *Analysis) AnalyseFlat(dir *fs.FsFold) *Analysis {
	ca := a.Descend(Flat, dir, dir.Recurse())

	dS, fS, _ := fs.MatchDisk(dir.String())
	ca.takeFileS(fS...)

	if len(dS) > 0 {
		for _, d := range dS {
			if dir.Recurse() {
				d = d.AsRecurse()
			}
			ca.collectFlat(d)
		}
	} else {
		ca.collectFlat(dir)
	}

	return ca
}

func (a *Analysis) collectFlat(dir *fs.FsFold) *Analysis {
	a = a.takeFileS(dir.FileS(extPatternS...)...)
	if dir.Recurse() {
		a = a.collectFlatSubDirS(dir)
	}
	return a
}

func (a *Analysis) collectFlatSubDirS(dir *fs.FsFold) *Analysis {
	dirFoldS := dir.SubDirS()
	for _, dirName := range dirFoldS { // SubDirS returned all
		if dirName.IsFold() { // evident - kept for symmetry with ...Down
			downDir := dirName //
			ca := a            // no Descend(downDir, true)
			ca = ca.takeFileS(downDir.FileS(extPatternS...)...)
			// collectDownSubDirS(downDir) // recurse down // no need
		}
	}
	return a
}
