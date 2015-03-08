package yuml

// Copyright 2015, Daniel Garcia <daniel@danielgarcia.info>
// yuml2 is distributed under the MIT license.
//
// Based on the python yuml client https://github.com/wandernauta/yuml
// by Wander Nauta

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Options to generate a yuml.me diagram.
type Options struct {
	// Direction of the digram LR (left to right), RL (right to left),
	// or TD (top to down).
	Direction string
	// Scale is the percentage to scale the image by. 0 is default, no scaling.
	Scale int
	// Style is the style of the diagram (scruffy, nofunky, plain). Scruffy is
	// the default.
	Style string
	// Format is the output type of the image (png, pdf, jpg, svg).
	Format string
	// Use is the type of digram (class, activity, usecase)
	Use string
}

var BaseURL = "http://yuml.me/diagram"

var (
	// ErrUnkUse indicates that the given Option.Use is not valid
	ErrUnkUse = errors.New("use is not one of class, activity, or usecase")
	// ErrDirection indicates diagram direction is not one of RL, LR, or TD."
	ErrDirection = errors.New("unknown diagram direction")
	// ErrFormat indicates output type is unknown (png, pdf, jpg, svg).
	ErrFormat = errors.New("unknown format type")
)

func validate(opts Options) (Options, error) {
	switch opts.Use {
	case "class":
	case "activity":
	case "usecase":
	default:
		return opts, ErrUnkUse
	}

	switch opts.Direction {
	case "LR":
	case "RL":
	case "TD":
	case "":
		opts.Direction = "LR"
	default:
		return opts, ErrDirection
	}

	switch opts.Format {
	case "png":
	case "pdf":
	case "jpg":
	case "svg":
	default:
		return opts, ErrFormat
	}
	return opts, nil
}

// Generate a yuml diagram. The input is new line seperated
// yuml statements. The output is a stream of the desired image format.
func Generate(opts Options, input io.Reader, output io.Writer) error {

	// validate options and set defaults
	opts, err := validate(opts)
	if err != nil {
		return err
	}
	urlopts := opts.Style
	if opts.Scale != 0 {
		urlopts += fmt.Sprintf(";scale:%d", opts.Scale)
	}
	if len(opts.Direction) > 0 {
		urlopts += fmt.Sprintf(";dir:%s", opts.Direction)
	}
	targetURL := fmt.Sprintf("%s/%s/%s/", BaseURL, urlopts, opts.Use)

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
	v.Set("dsl_text", strings.Join(olines, ", ")+"."+opts.Format)
	resp, err := http.PostForm(targetURL, v)
	if err != nil {
		return err
	}
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	filename := strings.SplitN(string(bs), ".", 2)[0] + "." + opts.Format

	targetURL = targetURL + filename
	res, err := http.Get(targetURL)
	if err != nil {
		return err
	}
	_, err = io.Copy(output, res.Body)
	return err
}
