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
