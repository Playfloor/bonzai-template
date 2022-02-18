# Bonzai™ Sample `foo` Command (Template)

*🎉 Bonzai requires Go 1.18 ❤*

![WIP](https://img.shields.io/badge/status-wip-red)
[![GoDoc](https://godoc.org/bonzai-foo?status.svg)](https://godoc.org/bonzai-foo)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

This repo is a GitHub template designed to help create new completable
Bonzai command trees quickly. It is kept in sync with the [latest bonzai
tagged version](https://github.com/rwxrob/bonzai). See the [check
list](#check-list) to be sure you change everything and don't leave any
of the defaults that might confuse users of your module. Here are some
things to consider while creating your first command module:

* Every place `foo` appears in this document is something that needs to
  change. (Find them easily with `git grep foo`.) Change the `foo.go`
  name and package to your own.

* Keep in mind that Bonzai commands document themselves with
  embedding usage and other documentation and therefore will likely not
  need a lot of content in this [README.md](README.md) file. 

* You probably want to leave [Install](#install) and [Embedded
  Documentation](#embedded-documentation) alone except for the name.

* You probably only want to change the name in the [Legal](#legal)
  section. Of course, you can put whatever you want in here. But the
  CONTRIBUTING file is an legally precedented way of accepting
  contributions and is used by both the Linux and GitLab projects.

* Keeping the trademark paragraph of this README.md to help preserve the
  identity of the Bonzai™ project. We appreciate your help.

## Install

*Change this for your own install instructions and delete this line.*

This command can be installed as a standalone program, composed into a
"tree" of commands.

Use `go install` to install as a standalone:

```
go install github.com/rwxrob/bonzai-foo/foo@latest
```

> ⚠️ Note the use of  `go get` for installation has been
> deprecated since 1.17 even though many projects still included
> outdated installation documentation.

Use `import` as usual to add to a box:

```go
package main

import (
	"github.com/rwxrob/bonzai"
	foo "github.com/rwxrob/bonzai-foo"
)

func main() {
  box := new(cmdbox.Box)
	box.Cmd = foo.Cmd
	box.Run()
}
```

## Embedded Documentation

See the [`foo.go`](foo.go) file itself for additional embedded
documentation about this command.

## Usage

```
foo
foo bar
foo help
```

## Check List

- [ ] Update GoDoc link
- [ ] Update starting summary paragraph
- [ ] Change name in the [Install](#install) section
- [ ] Update the [Usage](#usage) section
- [ ] Add more sections (if needed) but avoid
- [ ] Change names in the [Legal](#legal) section
- [ ] Confirm [LICENSE](LICENSE) file
- [ ] Confirm [CONTRIBUTING](CONTRIBUTING) file
- [ ] Rename the [`foo`](foo) directory
- [ ] Rename the [`foo.go`](foo.go) file
- [ ] Update the import in [`foo/main.go`](foo/main.go) file
- [ ] Update the [`go.mod`](go.mod) file (or run `go mod init`)
- [ ] Optionally remove `help` imports from [`foo.go`](foo.go)
- [ ] Code the [`foo.go`](foo.go) file and any dependencies
- [ ] Remove WIP tag from README.md when ready
* [ ] Optionally register your repo by submitting [GitHub PR]

[GitHub PR]: <https://github.com/rwxrob/register-bonzai>

## Legal

Copyright (c) 2021 Robert S. Muhlestein
Released under the [Apache 2.0](LICENSE)

Contributors and project participants implicitly accept the 
[Developer Certificate of Authenticity (DCO)](DCO).

"Bonzai" and "bonzai" (purposefully misspelled) are legal trademarks of
Robert S. Muhlestein but can be used freely to refer to the Bonzai
project <https://github.com/rwxrob/bonzai> without limitation. To avoid
potential developer confusion, intentionally using these trademarks to
refer to other projects --- free or proprietary --- is prohibited.
