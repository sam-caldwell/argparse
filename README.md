argparse
===========

## Description

A python-approximate golang commandline argument parser. The goal is to create a Command-line argument
parser similar to python argparse. Obviously, golang will have some differences.

## Usage

See [example](example/main.go).
```golang
package main

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"log"
)

// main - Demonstration program

func main() {
	log.Println("starting example...")
	var args argparse.Arguments
	args.
		ProgramName().
		Copyright(2023, "Sam Caldwell", "mail@samcaldwell.net>").
		Preamble("This is a description of our Arguments Object.").
		Postscript("This is the postfix string after help is dumped.").
		ShowErrors().
		Add("all", "-a", "--all", types.Flag, true, false, "All flag").
		Add("number", "-n", "--number", types.Integer, true, 1337, "set number").
		Add("bool", "-b", "--bool", types.Boolean, true, false, "set boolean").
		ExitOnError().
		Parse()

	fmt.Println(args.Help())

}
```

