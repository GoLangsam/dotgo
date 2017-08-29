// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	aDot = "."
)

func nameLessExt(path string) string {
	name := filepath.Base(path)
	return strings.TrimSuffix(name, filepath.Ext(name))
}

// matchBool - iff flag
func matchBool(flag bool) itemIs {
	return func(path string) bool {
		return flag
	}
}

// matchFunc - iff filename matches any pattern
func matchFunc(pattern ...string) itemIs {
	return func(path string) (matched bool) {
		for i := range pattern {
			matched, _ = filepath.Match(pattern[i], filepath.Base(path)) // ignore errors
			if matched {
				break
			}
		}
		return matched
	}
}

func ifFlagSkipDirWf(match itemIs) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if match(path) {
			return nil
		}
		return filepath.SkipDir
	}
}

func (t *toDo) isDirWf(dirWf, filWf filepath.WalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		switch {
		case t.ctx.Err() != nil:
			return filepath.SkipDir
		case info.IsDir():
			return dirWf(path, info, err)
		default:
			return filWf(path, info, err)
		}
	}
}
