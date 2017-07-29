package fsc

import (
	"path/filepath"

	"github.com/golangsam/container/ccsafe/fs"
)

// PipeFsFileGlob forks received fsfiles into directories and files according to fs.MatchDisk
func PipeFsFileGlob(
	inp <-chan fs.FsFile,
	dirS chan<- fs.FsFold,
	filS chan<- fs.FsFile) (
	out <-chan struct{}) {
	cha := make(chan struct{})
	go func() {
		defer close(cha)
		for name := range inp {
			dS, fS, _ := fs.MatchDisk(filepath.Join(name.String(), "*.tmpl"))
			for _, d := range dS {
				dirS <- *d
			}
			for _, f := range fS {
				filS <- *f
			}
		}
		cha <- struct{}{}
	}()
	return cha
}
