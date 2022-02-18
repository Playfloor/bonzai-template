package foo

import (
	"log"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmd"
)

var Cmd = &bonzai.Cmd{

	Name:      `foo`,
	Summary:   `just a sample foo command`,
	Usage:     `[B|bar|own|h|help]`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2`,
	Commands:  []*bonzai.Cmd{cmd.Help, bar, own},

	Description: `
		The foo commands do foo stuff. You can start the description here
		and wrap it to look nice and it will just work. Otherwise, just
		follow the same guidelines as for Go documentation. Note that the
		x.Call Method here is omitted since the main work is delegated to
		the subcommands in the command tree. The help command, however, is
		the default because it is first. `,

	// no Call since has Commands, if had Call would only call if
	// commands didn't match
}

var bar = &bonzai.Cmd{
	Name:     `bar`,
	Aliases:  []string{"B", "notbar"}, // to make a point
	Commands: []*bonzai.Cmd{cmd.Help},
	Call: func(none ...string) error {
		log.Printf("would bar stuff")
		return nil
	},
}
