// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package template implements the interface (as type "Template")
// common to both standard packages: "text/template" and "html/template".
//
// Thus, it exploits the fact, that
//   "package "html/template" has the same interface as package "text/template"
//    but automatically secures HTML output against certain attacks."
// as said in "go doc text/template".
//
// All methods and all package level functions are forewarded,
// except for the method "AddParseTree" (which is considered internal)
// and the two package level functions "ParseFiles" and "ParseGlob",
// which are just conveniences for something as simple as
//   New(name).ParseFiles(name [, morenames...])
//
// Instead of a single New(name) this package unsurprisinlgy provides
// two constructors: template.Text(name) & template.Html(name)
//
// Thus, the exported type "Template" represents the template used,
// be it html or text.
//
// Also, the type FuncMap is forwarded.
//
// Note: Clients in need for any other type such as ExecError (from "text/template")
// or data types such as HTML, CSS, JS and friends as well as Error and ErrorCode (from "html/template")
// are requested to use the respective standard package directly for access to the error and data types.
package template

import (
	"io"
)

// Template represents the template used (html or text)
type Template interface {
	// AddParseTree(name string, tree *parse.Tree) (Template, error) // intentionally not supported
	Clone() (Template, error)
	DefinedTemplates() string
	Delims(left, right string) Template
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Funcs(funcMap map[string]interface{}) Template
	Lookup(name string) Template
	Name() string
	New(name string) Template
	Option(opt ...string) Template
	Parse(text string) (Template, error)
	ParseFiles(filenames ...string) (Template, error)
	ParseGlob(pattern string) (Template, error)
	Templates() []Template
}

// FuncMap is the type of the map defining the mapping from names to functions.
// Each function must have either a single return value, or two return values of
// which the second has type error. In that case, if the second (error)
// return value evaluates to non-nil during execution, execution terminates and
// Execute returns that error.
//
// When template execution invokes a function with an argument list, that list
// must be assignable to the function's parameter types. Functions meant to
// apply to arguments of arbitrary type can use parameters of type interface{} or
// of type reflect.Value. Similarly, functions meant to return a result of arbitrary
// type can return interface{} or reflect.Value.
type FuncMap map[string]interface{}

// Must is a helper that wraps a call to a function returning (Template, error)
// and panics if the error is non-nil. It is intended for use in variable
// initializations such as
//	var t = template.Must(template.NewText|HtmlTemplate("name").Parse("text"))
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}
