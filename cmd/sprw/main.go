package main

import (
	"os"

	"github.com/varunamachi/vaali/vmgo"

	"github.com/varunamachi/sprw"
	"github.com/varunamachi/vaali/vapp"
	"github.com/varunamachi/vaali/vcmn"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := vapp.NewWebApp(
		"sprw",
		vcmn.Version{
			Major: 0,
			Minor: 0,
			Patch: 0,
		},
		"0",
		[]cli.Author{
			cli.Author{
				Name: "Varuna Amachi",
			},
		},
		"Sprw entity manager",
	)
	app.Modules = append(app.Modules, sprw.NewModule())
	vmgo.SetDefaultDB("sprw")
	app.Exec(os.Args)
}
