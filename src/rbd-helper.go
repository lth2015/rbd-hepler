package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func CreateImage(rbd string, size string, pool string) {

	fmt.Println("\nCreate a RBD image: name=%s, size=%s, pool=%s\n", rbd, size, pool)

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

	fmt.Printf("The result is %s\n", out)
}


func MapImage(rbd string) {

	fmt.Println("\nMap a RBD image: name=%s", rbd)

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

	fmt.Println("Show mapped a given image\n")

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

	fmt.Println("Show all rbd blocks mapped\n")

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

	fmt.Println("Unmapped the block: path=%s\n", path)

	// Exec create command
	cmd := "rbd unmapped " + path

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

	fmt.Println("Format the block device: fs=%s, path=%s\n", fs, path)

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

	fmt.Println("Delete the block: rbd=%s\n", rbd)

	// Exec create command
	cmd := "rbd unmapped " + rbd

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

	// rbd command
	// cmd := "rbd create ttt -s 1024"
	cmd := "rbd showmapped"

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The result is %s\n", out)
}

