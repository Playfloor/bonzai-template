package foo

import (
	"log"

	"github.com/rwxrob/bonzai"
)

var own = &bonzai.Cmd{
	Name: `own`,
	Call: func(none ...string) error {
		log.Print("I'm in my own file.")
		return nil
	},
}
