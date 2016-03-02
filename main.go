package main

import (
  "fmt"
  "os"
  "github.com/codegangsta/cli"
  "github.com/duranmla/avatarme/avatar"
  "github.com/duranmla/avatarme/cmdutil"
)

var (
  Stdout *os.File = os.Stdout
)

func main() {
	app := cli.NewApp()

	app.Name = "avatarme"
	app.Usage = "CLI tool to generate hashes from your email"

	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "prints out a User struct with hash",
			Action: func(c *cli.Context) {
        email := requestCredentials();
        user := avatar.New(email)
				fmt.Println(user)
			},
		},
	}

	app.Run(os.Args)
}

func requestCredentials() (email string){
  fmt.Fprint(Stdout, "email: ")
  email = cmdutil.ReadLine()

  return email
}
