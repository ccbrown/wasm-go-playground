package compile

import (
	"flag"
	"os"

	"github.com/ccbrown/go-web-gc/go/cmd/compile/internal/gc"
	"github.com/ccbrown/go-web-gc/go/cmd/compile/internal/wasm"
	"github.com/ccbrown/go-web-gc/go/cmd/internal/objabi"
)

func Compile(file string) {
	objabi.GOARCH = "wasm"
	objabi.GOOS = "js"
	os.Args = []string{os.Args[0], "-p", "main", "-complete", "-dwarf=false", "-pack", file}
	gc.Main(wasm.Init)
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
