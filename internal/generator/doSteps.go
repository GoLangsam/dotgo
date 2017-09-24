// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

// ReadMetaAndPrint

func (s *step) prepReadMetaAndPrint() *step {

	// p f n d m r t => ?p f m n r t d
	s.filePile.flagPrint(af, afv, "af-File:")
	s.metaPile.flagPrint(am, amv, "am-Meta:")
	s.baseDict.flagPrint(an, anv, "an-Name:")
	s.rootTmpl.flagPrint(ar, arv, "ar-Root:")

	if ad || at { // temporary meta data - just to show
		s = s.readMeta(at, atv, "at-Data:")
		s.dataTree.flagPrint(ad, adv, "ad-"+aDot+aDot+aDot+aDot+":")
		s.dataTree = NewData(aDot) // forget
	}

	return s
}

func (s *step) execReadMetaAndPrint(path string) *step {
	// p f n d m r t => p f m n r t d
	flagPrintString(epv, path, "Directory")

	s.filePile.flagPrint(ef, efv, "ef-File:")
	s.metaPile.flagPrint(em, emv, "em-Meta:")
	s.baseDict.flagPrint(en, env, "en-Name:")
	s.rootTmpl.flagPrint(er, erv, "er-Root:")

	s = s.readMeta(et, etv, "et-Data:")
	s.dataTree.flagPrint(ed, edv, "ed-Data: "+path)

	return s
}

// readMeta - for each metaPile:
// read metadata (using a cloned template)
// and optionally print it
func (s *step) readMeta(flag, verbose bool, header string) *step {

	tmpl, err := s.rootTmpl.Clone() // Clone rootTmpl
	if all.Ok("Clone", "Root", err) {
		metaData := s.metaReader(s.lookupData, Template{tmpl}) // text/template from meta
		s.metaPile.Walker(s.done, metaData)()                  // meta => metaTmpl & metaData
		metaData.flagPrint(flag, verbose, header)
	}

	return s
}
