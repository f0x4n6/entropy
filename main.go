// File entropy calculation tool.
//
// Usage:
//
//	entropy path
//
// The arguments are:
//
//	path
//	    File or folder to calculate (required).
package main

import (
	"fmt"
	"io"
	"io/fs"
	"math"
	"os"
	"path/filepath"
)

// https://gist.github.com/n2p5/4eda328b080c9f09eff928ad47228ab1
func entropy(name string) (n float64, err error) {
	f, err := os.Open(name)

	if err != nil {
		return
	}

	defer func() { _ = f.Close() }()

	buf, err := io.ReadAll(f)

	if err != nil {
		return
	}

	a := make([]float64, 256)

	for _, b := range buf {
		a[b]++
	}

	l := float64(len(buf))

	for i := range 256 {
		if a[i] != 0 {
			v := a[i] / l
			n -= v * math.Log2(v)
		}
	}

	return
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		_, _ = fmt.Fprintln(os.Stderr, "usage: entropy PATH")
		os.Exit(2)
	}

	err := filepath.WalkDir(os.Args[1], func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return nil
		}

		if d.IsDir() {
			return nil
		}

		val, err := entropy(path)

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return nil
		}

		abs, err := filepath.Abs(path)

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return nil
		}

		_, _ = fmt.Printf("%.10f  %s\n", val, abs)
		return nil
	})

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
