// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsa

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type Kind int

const (
	Root Kind = iota
	Flat
	Down
)

type Analysis struct {
	*fs.FsFold              // folder / directory
	Kind       Kind         // kind of analysis
	FsFileS    []*fs.FsFile // Files for data collection
	FsBaseS    []*fs.FsBase // Base names for tmpl execution
	Root       *Analysis    // root of analysis
	SubDirS    []*Analysis  // directories directly below
}

func newAnalysis(kind Kind, fsInfo *fs.FsFold) *Analysis {
	a := &Analysis{fsInfo, kind, []*fs.FsFile{}, []*fs.FsBase{}, nil, []*Analysis{}}
	a.Root = a
	return a
}

// Reset clears the data (files, bases & waydown)
func (a *Analysis) Reset() *Analysis {
	a.FsFileS = []*fs.FsFile{} // Files for data collection
	a.FsBaseS = []*fs.FsBase{} // Base names for tmpl execution
	a.SubDirS = []*Analysis{}  // way down
	return a
}

// Descend & inherit

func (a *Analysis) Descend(kind Kind, d *fs.FsFold, recurse bool) *Analysis {
	var ca *Analysis
	if recurse && !d.Recurse() {
		ca = newAnalysis(kind, d.AsRecurse())
	} else {
		ca = newAnalysis(kind, d)
	}
	a.SubDirS = append(a.SubDirS, ca) // append self to parent
	ca.Root = a.Root                  // carry root from above
	return ca
}

// helpers

func (a *Analysis) takeFileS(files ...*fs.FsFile) *Analysis {
	a.FsFileS = append(a.FsFileS, a.fileNameS(files...)...)
	a.FsBaseS = append(a.FsBaseS, a.baseNameS(files...)...)
	cacheFiles(files...)
	return a
}

func (a *Analysis) fileNameS(files ...*fs.FsFile) (nameS fs.FsFileS) {
	nameS = a.MatchingFileS(files, extPatternS...)
	return nameS
}

func (a *Analysis) baseNameS(files ...*fs.FsFile) (nameS fs.FsBaseS) {
	nameS = a.MatchingBaseS(files, namePattern)
	if len(nameS) < 1 && !a.IsFold() { // no files && not a fold ?
		nameS = append(nameS, a.Base()) // use basename
	}
	return nameS
}
