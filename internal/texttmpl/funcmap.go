// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

// see also "do/dot/funcs"

import (
	"os"
	"runtime"
	"text/template"
)

// Null - return empty string - useful to pipe method-results away
func null() interface{} {
	return ""
}

// os.Args as TempFunc "CmdLine"
func cmdline() interface{} {
	return os.Args
}

// os.Environ as TempFunc - use range on it, as it returns a slice
func environ() interface{} {
	return os.Environ()
}

// os.TempDir as TempFunc
func tempdir() interface{} {
	return os.TempDir()
}

// os.Getwd as TempFunc - silent on error
func getwd() interface{} {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	} else {
		return wd
	}
}

// os.Hostname as TempFunc - silent on error
func hostname() interface{} {
	hn, err := os.Hostname()
	if err != nil {
		return ""
	} else {
		return hn
	}
}

// runtime.MemStats as TempFunc - refreshed on each call
func memstats() interface{} {
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	return *stats
}

var Funcs = template.FuncMap{
	"null":     null,     // returns an empty string - useful to pipe method-results away
	"memstats": memstats, // runtime.MemStats - refreshed on each call
	"cmdline":  cmdline,  // os.Args
	"getwd":    getwd,    // os.Getwd - silent on error
	"hostname": hostname, // os.Hostname - silent on error
	"environ":  environ,  // os.Environ - use range on it, as it returns a slice
	// "tempdir":  tempdir,
}
