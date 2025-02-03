package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

var format = flag.String("f", "jpeg", "output format, required {png, jpeg, gif}")

func main() {
	if err := convertTo(os.Stdin, os.Stdout, "jpg"); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func convertTo(in io.Reader, out io.Writer, outKind string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Входной формат =", kind)
	if kind != outKind {
		switch strings.ToLower(outKind) {
		case "jpeg", "jpg":
			return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
		case "png":
			return png.Encode(out, img)
		case "gif":
			return gif.Encode(out, img, &gif.Options{})
		default:
			return fmt.Errorf("uncknown output wormat:%q", *format)

		}
	}
	return fmt.Errorf("new format is equal of old one format")
}
