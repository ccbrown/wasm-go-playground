package main

import (
	"flag"
	"os"

	"github.com/ccbrown/go-web-gc/go/cmd/compile"
	"github.com/ccbrown/go-web-gc/go/cmd/link"
)

func main() {
	var LinkFlagSet = flag.CommandLine
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	compile.Compile("main.go")
	flag.CommandLine = LinkFlagSet
	link.Link("main.a")
}
