package main

import (
	"fmt"
	"os"

	"github.com/varunamachi/vaali/vsec"

	"github.com/varunamachi/vaali/vlog"

	"github.com/varunamachi/vaali/vnet"

	"github.com/varunamachi/vaali/vapp"
	"github.com/varunamachi/vaali/vcmn"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := vapp.NewSimpleApp(
		"sprwcli",
		vcmn.Version{
			Major: 0,
			Minor: 0,
			Patch: 1,
		},
		"0",
		[]cli.Author{
			cli.Author{
				Name: "varunamachi",
			},
		},
		true,
		"Sparrow client",
	)
	app.Commands = append(app.Commands, cli.Command{
		Name: "login",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "server",
				Value: "http://localhost:8000",
				Usage: "Server address and port - <address>:<port>",
			},
			cli.StringFlag{
				Name:  "userID",
				Value: "",
				Usage: "User ID",
			},
			cli.StringFlag{
				Name:   "password",
				Value:  "",
				Usage:  "Password",
				Hidden: true,
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			ag := vcmn.NewArgGetter(ctx)
			host := ag.GetRequiredString("server")
			userID := ag.GetRequiredString("userID")
			password := ag.GetOptionalString("password")
			if len(password) == 0 {
				password = vcmn.AskPassword("Password")
			}
			if err = ag.Err; err == nil {
				c := vnet.NewClient(host, "sprw", "v0")
				err = c.Login(userID, password)
				if err == nil {
					fmt.Println("Login successful. User: ")
					vcmn.DumpJSON(c.User)
					data := vnet.M{}
					res := c.Get(&data, vsec.Public, "ping")
					if res.Err == nil {
						res.Read(&data)
						vcmn.DumpJSON(data)
					} else {
						fmt.Println(res.Err)
					}
				}
			}
			return vlog.LogError("Client", err)
		},
	})
	app.Exec(os.Args)
}
