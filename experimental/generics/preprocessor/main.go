package main

import (
	"fmt"
	"go/go2go"
	"io/ioutil"
	"os"
)

func main() {
	source, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fout := os.Stdout
	if len(os.Args) > 2 {
		f, err := os.Create(os.Args[2])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fout = f
	}

	importerTmpdir, err := ioutil.TempDir("", "go2go")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(importerTmpdir)

	importer := go2go.NewImporter(importerTmpdir)

	buf, err := go2go.RewriteBuffer(importer, os.Args[1], source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(2)
	}

	if _, err := fout.Write(buf); err != nil {
		panic(err)
	}
}
