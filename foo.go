package foo

import (
	"log"

	"github.com/rwxrob/bonzai"
)

var Cmd = &bonzai.Cmd{

	Name:      `foo`,
	Summary:   `just a sample foo command`,
	Usage:     `[bar|own|h|help]`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2`,
	Commands:  []*bonzai.Cmd{bonzai.Help, bar, own},

	Description: `
		The foo commands do foo stuff. You can start the description here
		and wrap it to look nice and it will just work. Otherwise, just
		follow the same guidelines as for Go documentation. Note that the
		x.Method here is commented out since the main work is delegated to
		the stuff subcommand. The help command, however, is the x.Default
		because it is first. `,

	// no Method since has Commands, if had Method would only call if
	// commands didn't match
}

var bar = &bonzai.Cmd{
	Name:     `bar`,
	Commands: []*bonzai.Cmd{bonzai.Help},
	Method: func(none ...string) error {
		log.Printf("would bar stuff")
		return nil
	},
}
