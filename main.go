package main

import (
	"os"

	"github.com/sabidrome/sabidrome/core"
)

func main() {
    command := os.Args[1]

    if command == "add" {
        path := os.Args[2]

    } else if command == "rm" {
        fmt.println("Oh no")

    } else {
        fmt.println("Ah ha")
    }
}
