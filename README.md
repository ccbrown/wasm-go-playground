# go-web-gc

This is the Go compiler ("gc") compiled for WASM, running in your browser! It can be used to run a simple playground, à la [play.golang.org](https://play.golang.org/) entirely in your browser!

You can try it out here: https://ccbrown.github.io/go-web-gc/server/

#### ⚠️ Important ⚠️

* Safari works, but is unbearably slow. **Chrome or Firefox for desktop is highly recommended.**
* Packages are not supported. Making them work would be straightforward, but I'm not compelled to spend time on it right now since this probably has no practical uses.

## Code

This repo very much has "proof of concept" / "make it work" organization. Here's a low effort explanation of what things are:

* ./go – Almost everything in this directory was copied directly from the Go source. The only exceptions are go/cmd/link/link.go and go/cmd/compile/compile.go. These files basically turn the "link" and "compile" commands into libraries that can be used to link and compile an executable.
* ./main.go – This uses the "compile" and "link" libraries to compile a single Go file into a WASM file. This is compiled to WASM and copied to ./server/build.wasm.
* ./server – This contains all the static files for an in-browser Go playground. Most of the files are either precompiled WASM or lightly modified copies of bits and pieces from [play.golang.org](https://play.golang.org/). The most substantial work in this directory is in index.html and wasm_exec.js as wasm_exec.js needed a virtual filesystem implementation.
