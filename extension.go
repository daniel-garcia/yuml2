package main

import (
	"errors"
	"path/filepath"
	"strings"
)

var (
	unknownExtension = errors.New("unknown extension")
)

func getExtensionType(defaultType, filename string) (string, error) {

	if len(defaultType) == 0 {
		ext := filepath.Ext(filename)
		if len(ext) < 2 || !strings.HasPrefix(ext, ".") {
			return "", unknownExtension
		}
		ext = ext[1:]
		defaultType = ext
	}
	defaultType = strings.ToLower(defaultType)
	switch defaultType {
	case "png":
		fallthrough
	case "pdf":
		fallthrough
	case "jpg":
		fallthrough
	case "svg":
		return defaultType, nil
	}

	return "", unknownExtension
}
