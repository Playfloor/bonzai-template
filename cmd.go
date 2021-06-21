package cmd

import (
	"github.com/rwxrob/cmdbox"
	_ "github.com/rwxrob/cmdbox-help"
	_ "github.com/rwxrob/cmdbox-version"
)

func init() {
	x := cmdbox.Add("foo", "h|help", "version", "stuff")
	x.Usage = `[h|help|version|stuff]`
	x.Summary = `just a sample foo command`

	x.Description = `
		The foo command does foo stuff. You can start the description here
		and wrap it to look nice and it will just work. Otherwise, just
		follow the same guidelines as for Go documentation. Note that the
		x.Method here is commented out since the main work is delegated to
		the stuff subcommand. The help command, however, is the x.Default
		because it is first. `

	// 	x.Method = func(args []string) error {
	// 		fmt.Println("would do foo stuff")
	// 		return nil
	// 	}

}
