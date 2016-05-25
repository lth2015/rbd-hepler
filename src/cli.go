package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "rbd"

	app.Usage = "Used to create, format, delete a rbd image!"

	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"a"},
			Usage:   "Create a given-size block and format it using a given filesystem",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Value: "rbd",
					Usage: "set the name of the rbd block",
				},
				cli.StringFlag{
					Name:  "size, s",
					Value: "10240",
					Usage: "set the size of the rbd block",
				},
				cli.StringFlag{
					Name:  "filesystem, f",
					Value: "ext4",
					Usage: "set the filesystem of the rbd block",
				},
				cli.StringFlag{
					Name:  "pool, p",
					Value: "rbd",
					Usage: "set the pool of the rbd block",
				},
			},
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					fmt.Println("Option number: ", len(c.Args()))
				}

				// Print the value of name option
				fmt.Println("Name Option: ", c.String("name"))
				fmt.Println("Size Option: ", c.String("size"))
				fmt.Println("Filesystem Option: ", c.String("filesystem"))

				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete a rbd block",
			Action: func(c *cli.Context) error {
				fmt.Println("delete the rbd: ", c.Args().First())
				return nil
			},
		},
	}

	app.Run(os.Args)
}
