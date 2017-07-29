# dotgo
An agnostic tool to create `*.go` sources (and other text) given some template(s) - generic types made easy.

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoLangsam/dotgo)](https://goreportcard.com/report/github.com/GoLangsam/dotgo)
[![Build Status](https://travis-ci.org/GoLangsam/dotgo.svg?branch=master)](https://travis-ci.org/GoLangsam/dotgo)
[![GoDoc](https://godoc.org/github.com/GoLangsam/dotgo?status.svg)](https://godoc.org/github.com/GoLangsam/dotgo)

## Simplicity ain't easy

Many weeks were spent in order to come up with a tool named `dotgo` which is easy to use.

And: `dotgo` offers a simple yet powerful way to generate [go sources](https://golang.org) in particular (and other kinds of texts in general).

## 2017-07-12
Today [github](https://github.com/) invites us to [Join GitHub in support of the open internet, again](https://github.com/blog/2396-join-github-in-support-of-the-open-internet-again).

*[This](https://www.battleforthenet.com/) is important to us, as is [Open Software](http://www.OpenSoftware.org/)*.

As are *Freedom* and *Transparency*. And we love to contribute, and to share. Thus:

Today `dotgo` goes public here - in the hope it shall become useful to the community.

## Please be patient
As of this writing - some refactoring is still under way in order to achive this, and in order to improve explanations and examples.

( Note: a couple of days ago, an example of what can be achieved was already published [here](https://github.com/GoLangsam/AnyType) )

More will follow soon - **Thank You** for Your patience!

## Basic Usage
Imagine, You have Your templates and definitions in place. (We'll show You later, how to achieve this with ease.), 

`dotgo` just needs to know:
- the template(s) to be used
- the location for resulting output file(s).

Thus, a simple 

	dotgo templates-dir target-dir

will do. Or - if these are same, and You already went there, use

	dotgo .

and the magic shall happen.

Hint: It's use is intentionally kept so super-simple in order to allow ease of use, e.g. in Your source files, with the `generate` tool and it's workflow.

## Remarks

### `dotgo` uses `text/template`
`dotgo` builds heavily on the `text/template` package from the standard library. For reasons to be shared later elswhere.

### `dottxt` - currently: No Need
And, as it's 'go awarness' can be switched off, there is currently no need for a twin such as `dottxt`.
(This may change, if we feel need to add awarness for go specific stuff such as packages or vendoring.
If so, there shall be a plain vanilla `dottxt` - also useful e.g. for `*.md`)

### `dothtml` - coming soon
Now -as You might now- the standard library provides a second template package: `html/template`. Even so being isomorphic, it has other benefits when applied to xml/html.
Thus, soon after having `dotgo` public and stable, there shall be it's companion `dothtml` - just using the other package.

### same name - different meanings
When I chose the name `dotgo` I was not aware of the website [dotgo.eu](https://www.dotgo.eu/) and it's related events.
Please accepty my apologies if this creates any kind of confusion in Your mind.
And *Yes* - I may deserve blames and flames not having taken time to research the name beforehand ...
