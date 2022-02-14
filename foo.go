package foo

import (
	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/cmdbox/util"
)

var Cmd = cmdbox.Command{

	Name:      `foo`,
	Summary:   `just a sample foo command`,
	Usage:     `[bar|help]`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 Robert S Muhlestin`,
	License:   `Apache-2`,
	Commands:  []cmdbox.Cmd{cmdbox.Help, bar},

	Description: `
		The foo command does foo stuff. You can start the description here
		and wrap it to look nice and it will just work. Otherwise, just
		follow the same guidelines as for Go documentation. Note that the
		x.Method here is commented out since the main work is delegated to
		the stuff subcommand. The help command, however, is the x.Default
		because it is first. `,

	Method: func(none ...string) error {
		util.Log("would foo stuff")
		return nil
	},
}

var bar = cmdbox.Command{
	Name:     `bar`,
	Commands: []cmdbox.Cmd{cmdbox.Help},
	Method: func(none ...string) error {
		util.Log("would bar stuff")
		return nil
	},
}
