// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"path/filepath"
	"sync"
)

// step represents data relevant for processing a node
type step struct {
	wg *sync.WaitGroup

	lookupData func(string) string

	isFile, isBase, isExec, hasMeta itemIs

	todo *toDo

	filePile NextPile
	baseDict Dict
	metaPile PrevPile
	rootTmpl Template
	dataTree Data
}

func NewStep(lookupData func(string) string) *step {

	s := new(step)

	s.wg = new(sync.WaitGroup)

	s.lookupData = lookupData

	// itemIs - some properties
	suffix := filepath.SplitList(tmplread)
	s.isFile = matchFunc(suffix...)
	s.isBase = matchFunc(suffix[0])
	s.isExec = matchFunc(suffix[len(suffix)-1])
	// s.isTrue = matchBool(true)
	s.hasMeta = func(item string) bool {
		meta, err := Meta(s.lookupData(item))
		all.Ok("hasMeta", item, err)
		return meta != ""
	}

	s.todo = doIt() // carries context & waitgroup

	// Some containers
	s.filePile = NewNext(512, 128) // files (templates) to handle
	s.metaPile = NewPrev(256, 000) // templates with non-empty meta: apply in reverse order!
	s.baseDict = NewDict()         // templates to execute: basenames found
	s.rootTmpl = NewTextTemplate(aDot) // text/template
	s.dataTree = NewData(aDot)     // data - a Dot
	return s
}

func (s *step) Clone() *step {
	n := new(step)

	n.lookupData = s.lookupData

	n.isFile = s.isFile
	n.isBase = s.isBase
	n.isExec = s.isExec
	n.hasMeta = s.hasMeta

	n.todo = s.todo // TODO todo !!!
	n.wg = s.wg     // TODO todo !!!

	n.filePile = NextPile{s.filePile.Clone()}
	n.baseDict = Dict{s.baseDict.Clone()}
	n.metaPile = PrevPile{s.metaPile.Clone()}

	tmpl, err := s.rootTmpl.Clone()
	if all.Ok("Clone", "template", err) {
		n.rootTmpl = Template{tmpl}
	} else {
		panic(err)
	}

	n.dataTree = NewData(aDot) // TODO Clone?
	return n
}

func (s *step) do(do func()) {
	s.wg.Add(1)

	go func(s *step, do func()) {
		defer s.wg.Done()
		do()
	}(s, do)
}
