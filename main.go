package main

import (
	"os"
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	cli.Run(os.Args)
}
