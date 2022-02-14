# cmdbox `foo` Command Module (Template)

*🎉 cmdbox uses and requires Go 1.18 ❤*

![WIP](https://img.shields.io/badge/status-wip-red)
[![GoDoc](https://godoc.org/cmdbox-foo?status.svg)](https://godoc.org/cmdbox-foo)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

This repo is a GitHub template designed to help create new cmdbox
command modules quickly. It is kept in sync with the [latest cmdbox
tagged version](https://github.com/rwxrob/cmdbox). See the [check
list](#check-list) to be sure you change everything and don't leave any
of the defaults that might confuse users of your module. Here are some
things to consider while creating your first command module:

* Every place `foo` appears in this document is something that needs to
  change. (Find them easily with `git grep foo`.) Change the `foo.go`
  name and package to your own.

* Keep in mind that the cmdbox modules document themselves with
  embedding usage and other documentation and therefore will likely not
  need a lot of content in this [README.md](README.md) file. 

* You probably want to leave [Install](#install) and [Embedded
  Documentation](#embedded-documentation) alone except for the name.

* You probably only want to change the name in the [Legal](#legal)
  section. Of course, you can put whatever you want in here. But the
  CONTRIBUTING file is an legally precedented way of accepting
  contributions and is used by both the Linux and GitLab projects.

* Keeping the trademark paragraph of this README.md helps preserve the
  identity of the cmdbox project. We appreciate your help.

## Install

This command can be installed as a standalone program or composed into a
"box" of commands combined with other command modules.

Use `go install` to install as a standalone:

```
go install github.com/rwxrob/cmdbox-foo/foo@latest
```

> ⚠️ Note the use of `go get` for installation has been
> deprecated since 1.17 even though many projects still included
> outdated installation documentation.

Use `import` as usual to add to a box:

```go
package main

import (
	"github.com/rwxrob/cmdbox"
	foo "github.com/rwxrob/cmdbox-foo"
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

[GitHub PR]: <https://github.com/rwxrob/register-cmdbox>

## Legal

Copyright (c) 2021 Robert S. Muhlestein
Released under the [Apache 2.0](LICENSE)

Contributors and project participants implicitly accept the 
[Developer Certificate of Authenticity (DCO)](DCO).

"CmdBox" and "cmdbox" are legal trademarks of Robert S. Muhlestein but
can be used freely to refer to the cmdbox project
<https://github.com/rwxrob/cmdbox> without limitation. To avoid
potential developer confusion, intentionally using these trademarks to
refer to other projects --- free or proprietary --- is prohibited.
