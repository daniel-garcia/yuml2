package main

// Copyright 2015, Daniel Garcia <daniel@danielgarcia.info>
// yuml is distributed under the MIT license.
//
// Based on the python yuml client https://github.com/wandernauta/yuml
// by Wander Nauta

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type optionsT struct {
	direction string
	scale     int
	style     string
	format    string
	use       string
}

var (
	options = optionsT{}
	baseURL = "http://yuml.me/diagram"
)

func generate(opts optionsT, input io.Reader, output io.Writer) error {
	urlopts := opts.style
	if opts.scale != 0 {
		urlopts += fmt.Sprintf(";scale=%d", opts.scale)
	}
	if len(opts.direction) > 0 {
		urlopts += fmt.Sprintf(";dir=%s", opts.direction)
	}
	targetURL := fmt.Sprintf("%s/%s/%s/", baseURL, urlopts, opts.use)

	bs, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}
	olines := make([]string, 0)
	for _, line := range strings.Split(string(bs), "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		olines = append(olines, line)
	}
	v := url.Values{}
	v.Set("dsl_text", strings.Join(olines, ", ")+"."+opts.format)
	resp, err := http.PostForm(targetURL, v)
	if err != nil {
		return err
	}
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	filename := strings.SplitN(string(bs), ".", 2)[0] + opts.format

	targetURL = targetURL + filename
	res, err := http.Get(targetURL)
	if err != nil {
		return err
	}
	_, err = io.Copy(output, res.Body)
	return err
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [INPUT_FILE] [OUTPUT_FILE]\n", os.Args[0])
		flag.PrintDefaults()
	}

	if endpoint := os.Getenv("YUML_ENDPOINT"); len(endpoint) > 0 {
		baseURL = endpoint
	}
	flag.IntVar(&options.scale, "scale", 0, "percentage to scale output")
	flag.StringVar(&options.direction, "direction", "LR", "text direction (LR, RL, TD)")
	flag.StringVar(&options.style, "style", "scruffy", "style of image (scruffy, nofunky, plain)")
	flag.StringVar(&options.format, "format", "", "format of output (png, pdf, jpg, svg)")
	flag.StringVar(&options.use, "t", "class", "type of diagram (class, activity, usecase)")
	flag.StringVar(&baseURL, "u", baseURL, "base url for service (YUML_ENDPOINT) ")
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
		f, err := os.Create(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open %s: %s\n", args[1], err)
		}
		defer f.Close()
		output = f
	}
	if err := generate(options, input, output); err != nil {
		fmt.Fprintf(os.Stderr, "could not generate diagram: %s", err)
	}
}
