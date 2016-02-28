package main

import (
	"github.com/mkideal/cli"
)

var _ = app.Register(&cli.Command{
	Name: "clean",
	Desc: "Clean build data",
	Argv: func() interface{} { return new(clean_t) },
	Fn:   clean,
})

type clean_t struct {
	Help      bool `cli:"h,help" usage:"display help information"`
	Recursion bool `cli:"R,recursion" usage:"clean recursion or not" dft:"true"`
}

func clean(ctx *cli.Context) error {
	argv := ctx.Argv().(*clean_t)

	if argv.Help {
		ctx.String(ctx.Usage())
		return nil
	}
	ctx.String("%s: %v", ctx.Path(), jsonIndent(argv))
	return nil
}