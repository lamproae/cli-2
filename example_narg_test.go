package cli_test

import (
	"fmt"

	"github.com/mkideal/cli"
)

func ExampleNumArgFunc_exactN() {
	type argT struct {
		cli.Helper
		I int `cli:"i" usage:"int"`
	}

	app := func() *cli.Command {
		return &cli.Command{
			Name:    "hahaha",
			Argv:    func() interface{} { return new(argT) },
			NumArg:  cli.ExactN(1),
			UsageFn: func() string { return "usage function" },
			Fn: func(ctx *cli.Context) error {
				return nil
			},
		}
	}
	fmt.Println(app().Run([]string{}))
	fmt.Println(app().Run([]string{"-i", "1", "b"}))
	fmt.Println(app().Run([]string{"-i", "1", "b", "c"}))

	// Output:
	// usage function<nil>
	// <nil>
	// usage function<nil>
}

func ExampleNumArgFunc_atLeast() {
	type argT struct {
		cli.Helper
		I int `cli:"i" usage:"int"`
	}

	app := func() *cli.Command {
		return &cli.Command{
			Name:    "hahaha",
			Argv:    func() interface{} { return new(argT) },
			NumArg:  cli.AtLeast(1),
			UsageFn: func() string { return "usage function" },
			Fn: func(ctx *cli.Context) error {
				return nil
			},
		}
	}
	fmt.Println(app().Run([]string{}))
	fmt.Println(app().Run([]string{"-i", "1", "b"}))
	fmt.Println(app().Run([]string{"-i", "1", "b", "c"}))

	// Output:
	// usage function<nil>
	// <nil>
	// <nil>
}

func ExampleNumArgFunc_atMost() {
	type argT struct {
		cli.Helper
		I int `cli:"i" usage:"int"`
	}

	app := func() *cli.Command {
		return &cli.Command{
			Name:    "hahaha",
			Argv:    func() interface{} { return new(argT) },
			NumArg:  cli.AtMost(2),
			UsageFn: func() string { return "usage function" },
			Fn: func(ctx *cli.Context) error {
				return nil
			},
		}
	}
	fmt.Println(app().Run([]string{}))
	fmt.Println(app().Run([]string{"-i", "1", "b"}))
	fmt.Println(app().Run([]string{"-i", "1", "b", "c"}))
	fmt.Println(app().Run([]string{"-i", "1", "b", "c", "d"}))

	// Output:
	// <nil>
	// <nil>
	// <nil>
	// usage function<nil>
}

func ExampleNumOptionFunc_exactN() {
	type argT struct {
		cli.Helper
		I int `cli:"i" usage:"int"`
	}

	app := func() *cli.Command {
		return &cli.Command{
			Name:      "hahaha",
			Argv:      func() interface{} { return new(argT) },
			NumOption: cli.ExactN(1),
			UsageFn:   func() string { return "usage function" },
			Fn: func(ctx *cli.Context) error {
				return nil
			},
		}
	}
	fmt.Println(app().Run([]string{}))
	fmt.Println(app().Run([]string{"-i", "1"}))

	// Output:
	// usage function<nil>
	// <nil>
}

func ExampleNumOptionFunc_atLeast() {
	type argT struct {
		cli.Helper
		I int `cli:"i" usage:"int"`
	}

	app := func() *cli.Command {
		return &cli.Command{
			Name:      "hahaha",
			Argv:      func() interface{} { return new(argT) },
			NumOption: cli.AtLeast(1),
			UsageFn:   func() string { return "usage function" },
			Fn: func(ctx *cli.Context) error {
				return nil
			},
		}
	}
	fmt.Println(app().Run([]string{}))
	fmt.Println(app().Run([]string{"-i", "1"}))

	// Output:
	// usage function<nil>
	// <nil>
}

func ExampleNumOptionFunc_atMost() {
	type argT struct {
		cli.Helper
		I int `cli:"i" usage:"int"`
		J int `cli:"j" usage:"int"`
		K int `cli:"k" usage:"int"`
	}

	app := func() *cli.Command {
		return &cli.Command{
			Name:      "hahaha",
			Argv:      func() interface{} { return new(argT) },
			NumOption: cli.AtMost(2),
			UsageFn:   func() string { return "usage function" },
			Fn: func(ctx *cli.Context) error {
				return nil
			},
		}
	}
	fmt.Println(app().Run([]string{}))
	fmt.Println(app().Run([]string{"-i", "1"}))
	fmt.Println(app().Run([]string{"-i", "1", "-j=2"}))
	fmt.Println(app().Run([]string{"-i", "1", "-j=2", "-k=3"}))

	// Output:
	// <nil>
	// <nil>
	// <nil>
	// usage function<nil>
}
