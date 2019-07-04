package link

import (
	"os"

	"github.com/ccbrown/go-web-gc/go/cmd/link/internal/ld"
	"github.com/ccbrown/go-web-gc/go/cmd/link/internal/wasm"
)

func Link(file string) {
	os.Args = []string{os.Args[0], "-importcfg", "importcfg.link", "-buildmode=exe", file}
	ld.Main(wasm.Init())
}
