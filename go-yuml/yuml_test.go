package yuml

// Copyright 2015, Daniel Garcia <daniel@danielgarcia.info>
// yuml2 is distributed under the MIT license.
//
// Based on the python yuml client https://github.com/wandernauta/yuml
// by Wander Nauta

import (
	"bytes"
	"image"
	_ "image/png"
	"testing"
)

var testOpts = Options{
	Format: "png",
	Use:    "class",
}

func TestGenerate(t *testing.T) {

	input := bytes.NewBuffer([]byte("[User]"))
	output := new(bytes.Buffer)
	if err := Generate(testOpts, input, output); err != nil {
		t.Fatalf("could not generate diagram: %s", err)
	}

	if _, _, err := image.Decode(output); err != nil {
		t.Fatalf("could not decode output image: %s", err)
	}
}
