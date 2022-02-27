package foo

import (
	"log"

	Z "github.com/rwxrob/bonzai"
)

var own = &Z.Cmd{
	Name: `own`,
	Call: func(caller *Z.Cmd, none ...string) error {
		log.Print("I'm in my own file.")
		return nil
	},
}
