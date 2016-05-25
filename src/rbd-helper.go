package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

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

