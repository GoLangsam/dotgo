// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"github.com/golangsam/container/ccsafe/fs"
	"github.com/golangsam/dotgo/internal/fsc" // Fs Channels
)

// IfPrint helpers - using fsc: Fs Channels

// IfPrintFsNameS is a deprechated dummy
func IfPrintFsNameS(flag bool, nameS []FsName, prefix string) {
	dummy := fs.FsBaseS{} // TODO(apa): obsolete this dummy
	for _, name := range nameS {
		dummy = append(dummy, name.(*fs.FsBase))
	}
	IfPrintFsBaseS(flag, dummy, prefix)
}

// IfPrintFsBaseS prints the base names, iff flag is true
func IfPrintFsBaseS(flag bool, fsBaseS fs.FsBaseS, prefix string) {
	if flag {
		baseS := fsc.DoneFsBaseSlice(fsc.PipeFsBaseFunc(fsc.ChanFsBase(fsBaseS...), fsc.FsBasePrintFunc(prefix+"\t")))
		_ = <-baseS
	}
}

// IfPrintFsFileS prints the file names, iff flag is true
func IfPrintFsFileS(flag bool, fsFileS fs.FsFileS, prefix string) {
	if flag {
		fileS := fsc.DoneFsFileSlice(fsc.PipeFsFileFunc(fsc.ChanFsFile(fsFileS...), fsc.FsFilePrintFunc(prefix+"\t")))
		_ = <-fileS
	}
}

// IfPrintFsFoldS prints the fold names, iff flag is true
func IfPrintFsFoldS(flag bool, fsFoldS fs.FsFoldS, prefix string) {
	if flag {
		foldS := fsc.DoneFsFoldSlice(fsc.PipeFsFoldFunc(fsc.ChanFsFold(fsFoldS...), fsc.FsFoldPrintFunc(prefix+"\t")))
		_ = <-foldS
	}
}
