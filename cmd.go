package cmd

import (
	"github.com/rwxrob/cmdbox"
	_ "github.com/rwxrob/cmdbox-help"
	_ "github.com/rwxrob/cmdbox-version"
)

func init() {
	x := cmdbox.New("foo", "help", "version")
	x.Usage = ``
	x.Summary = `just a sample foo command`

	// TODO consider removing for minimal builds
	x.Description = `
		The *foo* command does cool stuff. You can start the description
		here and wrap it to look nice and it will just work. You can use
		minimal markdown with one, two, or three stars and wrap things with
		< and > to get uppercase with underline, as with options in a man
		page. That is, if you use the cmdbox-help command module. Other
		cmdbox-help modules might handle this Description differently but,
		by convention, all are supposed to support the same minimal
		CmdBox markdown.`

}
