package cmd

import (
	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("foo", "stuff")
	x.Summary = `just a sample foo command`
	x.Usage = `[h|help|version|stuff]`
	x.Version = `v0.0.1`
	x.Copyright = `Copyright 2021 Robert S Muhlestin`
	x.License = `Apache-2`
	x.Site = `https://rwxrob.github.io/cmdbox-foo`
	x.Source = `https://github.com/rwxrob/cmdbox-foo`
	x.Issues = `https://github.com/rwxrob/cmdbox-foo/issues`

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
