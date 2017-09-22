// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	ds "github.com/GoLangsam/do/strings" // Extract
)

const ( // restrict comment to "{{/*-" & "-*/}}" for Meta-Comments
	dash  = "-"
	metaL = tmplL + commL + dash
	metaR = dash + commR + tmplR
)

// Meta returns the meta-text extraced from text
func Meta(text string) (string, error) {
	return ds.Extract(text, metaL, metaR) // extract meta-data
}
