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

// IsDotNonsense matches .git or other dot nonsense.
func IsDotNonsense(name string) bool { return strings.HasPrefix(name, ".") }

// matchBool - iff flag
func matchBool(flag bool) itemIs {
	return func(path string) bool {
		return flag
	}
}

// matchFunc - iff filename matches any pattern
func matchFunc(pattern ...string) itemIs {
	return func(path string) bool {
		for i := range pattern {
			if strings.HasSuffix(path, pattern[i]) {
				return true
			}
		}
		return false
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

func isDirWf(quit func() bool, dirWf, filWf filepath.WalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		switch {
		case quit():
			return filepath.SkipDir
		case info.IsDir():
			return dirWf(path, info, err)
		default:
			return filWf(path, info, err)
		}
	}
}
