package main

// Copyright 2015, Daniel Garcia <daniel@danielgarcia.info>
// yuml2 is distributed under the MIT license.
//
// Based on the python yuml client https://github.com/wandernauta/yuml
// by Wander Nauta

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/daniel-garcia/yuml2/go-yuml"
)

const version = "0.2.0"

var (
	options = yuml.Options{}
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s v%s\n", os.Args[0], version)
		fmt.Fprintf(os.Stderr, "Usage: %s [options] INPUT_FILE OUTPUT_FILE\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "'-' can be used for stdin or stdout\n")
		flag.PrintDefaults()
	}

	if endpoint := os.Getenv("YUML_ENDPOINT"); len(endpoint) > 0 {
		yuml.BaseURL = endpoint
	}
	flag.IntVar(&options.Scale, "scale", 0, "percentage to scale output")
	flag.StringVar(&options.Direction, "direction", "LR", "text direction (LR, RL, TD)")
	flag.StringVar(&options.Style, "style", "scruffy", "style of image (scruffy, nofunky, plain)")
	flag.StringVar(&options.Format, "format", "", "format of output (png, pdf, jpg, svg)")
	flag.StringVar(&options.Use, "t", "class", "type of diagram (class, activity, usecase)")
	flag.StringVar(&yuml.BaseURL, "u", yuml.BaseURL, "base url for service (YUML_ENDPOINT) ")
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	var input io.Reader
	var output io.Writer

	if args[0] == "-" {
		input = os.Stdin
	} else {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open %s: %s\n", args[0], err)
			os.Exit(1)
		}
		defer f.Close()
		input = f
	}

	if args[1] == "-" {
		output = os.Stdout
	} else {
		var err error
		options.Format, err = getExtensionType(options.Format, args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not determine output type: %s\n", err)
		}
		f, err := os.Create(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open %s: %s\n", args[1], err)
		}
		defer f.Close()
		output = f
	}
	if err := yuml.Generate(options, input, output); err != nil {
		fmt.Fprintf(os.Stderr, "could not generate diagram: %s\n", err)
	}
}
