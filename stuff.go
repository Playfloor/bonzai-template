package cmd

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("foo stuff")
	x.Summary = `does foo stuff`
	x.Description = `This subcommand does the foo stuff.`

	x.Method = func(args []string) error {
		fmt.Println("would do foo stuff")
		return nil
	}

}
