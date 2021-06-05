# CmdBox Command Module Template

[![GoDoc](https://godoc.org/cmdbox-_foo?status.svg)](https://godoc.org/cmdbox-_foo)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

This repo is a GitHub template designed to help create new CmdBox
command modules quickly. It is kept in sync with the [latest CmdBox
tagged version](https://github.com/rwxrob/cmdbox). See the [check
list](#check-list) to be sure you change everything and don't leave any
of the defaults that might confuse users of your module. Every place
`foo` or `_foo` appear in this document is something that needs to
change. (Find them easily with `git grep foo`.) Keep in mind that the
CmdBox modules document themselves with embedding usage and other
documentation and therefore will likely not need a lot of content in
this [README.md](README.md) file. Keep the `cmd.go` and `main.go` file
names. Also keep the `cmd` package (which is ignored on import). (Delete
and replace this paragraph with `dap` or `cap` in `vi`.)

## Install 

This command can be installed as a standalone program or composed into a
CmdBox composite monolith.

Use `go install` to install as a standalone:

```
go install github.com/rwxrob/cmdbox-_foo/foo@latest
```

Use `import` with a blank identifier to be composed:

```go
import (
  "github.com/rwxrob/cmdbox"
  _ "github.com/rwxrob/cmdbox-_foo"
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
- [ ] Update the [Install](#install) section
- [ ] Update the [Usage](#usage) section
- [ ] Add more sections (if needed) but avoid
- [ ] Update the [Legal](#legal) section
- [ ] Confirm [LICENSE](LICENSE) file
- [ ] Confirm [DCO](DCO) file
- [ ] Rename the [`foo`](foo) directory
- [ ] Update the import in [`foo/main.go`](foo/main.go) file
- [ ] Update the [`go.mod`](go.mod) file (or run `go mod init`)
- [ ] Optionally remove `help` / `version` imports from [`cmd.go`](cmd.go)
- [ ] Code the [`cmd.go`](cmd.go) file and any dependencies
- [ ] Remove WIP tag when ready

## Legal

Copyright (c) 2021 Robert S. Muhlestein
Released under the [Apache 2.0](LICENSE)

Contributors and project participants implicitly accept the 
[Developer Certificate of Authenticity (DCO)](DCO).

