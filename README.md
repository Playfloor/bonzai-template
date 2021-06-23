# `cmdbox-foo` CmdBox Command Module Template

![WIP](https://img.shields.io/badge/status-wip-red)
[![GoDoc](https://godoc.org/cmdbox-foo?status.svg)](https://godoc.org/cmdbox-foo)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

This repo is a GitHub template designed to help create new CmdBox
command modules quickly. It is kept in sync with the [latest CmdBox
tagged version](https://github.com/rwxrob/cmdbox). See the [check
list](#check-list) to be sure you change everything and don't leave any
of the defaults that might confuse users of your module. Here are some
things to consider while creating your first command module:

* Every place `foo` appears in this document is something that needs to
  change. (Find them easily with `git grep foo`.) 

* Keep in mind that the CmdBox modules document themselves with
  embedding usage and other documentation and therefore will likely not
  need a lot of content in this [README.md](README.md) file. 

* Keep the conventional `cmd.go` and `main.go` file names.

* Keep the conventional `cmd` package name. This package name is never
  used in practice because it is for side-effect only (`import _
  <package>`). Where it ever needed, it could be easily named something
  other than `_` on the import line.

* Only remove `help` and `version` import if you absolutely must to
  minimize binary size as much as possible. Keep in mind that you
  probably want them even in multicall, minimal, `FROM scratch`
  containers because this provides at least some help documentation when
  using `exec` into the container/pod.

* You probably want to leave [Install](#install) and [Embedded
  Documentation](#embedded-documentation) alone except for the name.

* You probably only want to change the name in the [Legal](#legal)
  section. Of course, you can put whatever you want in here. But the DCO
  is an accepted way of handling legal entanglements with contributors
  and is used by both the Linux and GitLab projects. Keeping the
  trademark paragraph helps preserve the identity of the CmdBox project.

* Use fully qualified subcommands in their own files (ex: `stuff.go`).
  While this seems like overkill for small commands, it provides for
  better modularity and separation of concerns in terms of execution and
  documentation. This way a subcommand can easily be ported to another
  command module just by moving the file.

## Install

This command can be installed as a standalone program or composed into a
CmdBox composite monolith.

Use `go install` to install as a standalone:

```
go install github.com/rwxrob/cmdbox-foo/foo@latest
```

Use `import` with a blank identifier to be composed:

```go
import (
  "github.com/rwxrob/cmdbox"
  _ "github.com/rwxrob/cmdbox-foo"
)
```

## Embedded Documentation

See the [`cmd.go`](cmd.go) file itself for additional embedded
documentation about this command.

## Usage

```
foo
foo help
foo version
```

## Check List

- [ ] Update GoDoc link
- [ ] Update starting summary paragraph
- [ ] Change name in the [Install](#install) section
- [ ] Update the [Usage](#usage) section
- [ ] Add more sections (if needed) but avoid
- [ ] Change names in the [Legal](#legal) section
- [ ] Confirm [LICENSE](LICENSE) file
- [ ] Confirm [DCO](DCO) file
- [ ] Rename the [`foo`](foo) directory
- [ ] Update the import in [`foo/main.go`](foo/main.go) file
- [ ] Update the [`go.mod`](go.mod) file (or run `go mod init`)
- [ ] Optionally remove `help` / `version` imports from [`cmd.go`](cmd.go)
- [ ] Code the [`cmd.go`](cmd.go) file and any dependencies
* [ ] Remove or update the `stuff.go` file
- [ ] Remove WIP tag when ready
* [ ] Optionally register your repo by submitting [GitHub PR]

[GitHub PR]: <https://github.com/rwxrob/register-cmdbox>

## Legal

Copyright (c) 2021 Robert S. Muhlestein
Released under the [Apache 2.0](LICENSE)

Contributors and project participants implicitly accept the 
[Developer Certificate of Authenticity (DCO)](DCO).

"CmdBox" and "cmdbox" are legal trademarks of Robert S. Muhlestein but
can be used freely to refer to the CmdBox project
<https://github.com/rwxrob/cmdbox> without limitation. To avoid
potential developer confusion, intentionally using these trademarks to
refer to other projects --- free or proprietary --- is prohibited.
