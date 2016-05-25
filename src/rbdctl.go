package main

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func CreateImage(rbd string, size string, pool string) {

	fmt.Printf("\nCreate a RBD image: name=%s, size=%s, pool=%s\n", rbd, size, pool)
	fmt.Println("==========================================================================")

	// Exec create command
	cmd := "rbd create " + rbd + " -s " + size + " -p " + pool
	// cmd := "rbd showmapped"

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}

func MapImage(rbd string) {

	fmt.Printf("\nMap a RBD image: name=%s", rbd)
	fmt.Println("==========================================================================")

	// Exec map command
	cmd := "rbd map " + rbd

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The result is %s\n", out)
}

func ShowMappedImage(rbd string) string {

	fmt.Printf("Show mapped a given image\n")
	fmt.Println("==========================================================================")

	// Exec map command
	rbd_cmd := "rbd showmapped "

	// RBD shommaped
	parts := strings.Fields(rbd_cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	rbd_out := exec.Command(head, parts...)

	// Grep rbd
	grep_cmd := "grep " + rbd

	parts = strings.Fields(grep_cmd)
	head = parts[0]
	parts = parts[1:len(parts)]

	grep_out := exec.Command(head, parts...)

	// Pipe rbd and grep command
	var b2 bytes.Buffer
	r, w := io.Pipe()
	rbd_out.Stdout = w
	grep_out.Stdin = r
	grep_out.Stdout = &b2

	rbd_out.Start()
	grep_out.Start()
	rbd_out.Wait()
	w.Close()
	grep_out.Wait()

	buffer := b2.String()
	buffer = strings.Trim(buffer, " \n\t")
	for {
		if strings.Contains(buffer, "  ") {
			buffer = strings.Replace(buffer, "  ", " ", -1)
		} else {
			break
		}
	}

	paths := strings.Split(buffer, " ")

	path := paths[len(paths)-1]

	return path

}

func ShowMappedAllIamges() {

	fmt.Printf("Show all rbd blocks mapped\n")
	fmt.Println("==========================================================================")

	// Exec create command
	cmd := "rbd showmapped"

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}

func UnmapImage(path string) {

	fmt.Printf("Unmapped the block: path=%s\n", path)

	// Exec create command
	cmd := "rbd unmap " + path

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}

func MakeFileSystem(fs, path string) {

	fmt.Printf("Format the block device: fs=%s, path=%s\n", fs, path)
	fmt.Println("==========================================================================")

	// Exec create command
	cmd := "mkfs." + fs + " " + path

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}

func DeleteImage(rbd string) {

	fmt.Printf("Delete the block: rbd=%s\n", rbd)
	fmt.Println("==========================================================================")

	// Exec create command
	cmd := "rbd rm " + rbd

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}

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

				// Print the value of name option
				name := c.String("name")
				size := c.String("size")
				pool := c.String("pool")
				fs := c.String("filesystem")
				fmt.Println("Name Option: ", c.String("name"))
				fmt.Println("Size Option: ", c.String("size"))
				fmt.Println("Filesystem Option: ", c.String("filesystem"))

				// Create rbd
				CreateImage(name, size, pool)
				MapImage(name)
				path := ShowMappedImage(name)
				MakeFileSystem(fs, path)
				UnmapImage(path)

				fmt.Println("RBD is created successfully!")

				return nil
			},
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Show blocks mapped",
			Action: func(c *cli.Context) error {
				ShowMappedAllIamges()
				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete a rbd block",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Value: "rbd",
					Usage: "set the name of the rbd block",
				},
			},
			Action: func(c *cli.Context) error {
				name := c.String("name")
				DeleteImage(name)
				return nil
			},
		},
	}

	app.Run(os.Args)
}

