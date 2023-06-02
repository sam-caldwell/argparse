package main

import (
	"github.com/sam-caldwell/go-argparse/v2/argparse"
	"github.com/sam-caldwell/go-argparse/v2/types"
	"log"
)

// main - Demonstration program

func main() {
	var err error
	var args argparse.Arguments
	log.Println("starting example...")

	log.Println("Initialize Arguments")
	args.Init("This is an example program implementing argparse")
	args.EnableDebugging()

	log.Println("Add positional argument pos1")
	err = args.AddArgument(
		"pos1",
		types.String,
		false,
		"pos1value",
		"This is pos1 help")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Add positional argument pos2")
	err = args.AddArgument(
		"pos2",
		types.Integer,
		false,
		1337,
		"This is pos2 help")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Add optional argument -a")
	err = args.AddArgument(
		"-a",
		types.Integer,
		true,
		-1,
		"This is -a <integer> argument help")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Add optional argument --flag")
	err = args.AddArgument(
		"--flag",
		types.Flag,
		true,
		false,
		"This is --flag argument help")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Parse arguments against os.Args")
	err = args.ParseArgs()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("-------Dump() parsed args--------")
	for i, v := range args.Dump() {
		log.Println(i, v)
	}
	log.Println("---------------------------------")

	if value, err := args.Get("pos1"); err != nil {
		log.Fatalf("expected pos1 arg not found: %v", err)
	} else {
		log.Printf("expected pos1 arg value: %v", value)
	}

	log.Println("done...")
}
