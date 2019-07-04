// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mips64

import (
	"github.com/ccbrown/go-web-gc/go/cmd/compile/internal/gc"
	"github.com/ccbrown/go-web-gc/go/cmd/compile/internal/ssa"
	"github.com/ccbrown/go-web-gc/go/cmd/internal/obj/mips"
	"github.com/ccbrown/go-web-gc/go/cmd/internal/objabi"
)

func Init(arch *gc.Arch) {
	arch.LinkArch = &mips.Linkmips64
	if objabi.GOARCH == "mips64le" {
		arch.LinkArch = &mips.Linkmips64le
	}
	arch.REGSP = mips.REGSP
	arch.MAXWIDTH = 1 << 50
	arch.SoftFloat = objabi.GOMIPS64 == "softfloat"
	arch.ZeroRange = zerorange
	arch.ZeroAuto = zeroAuto
	arch.Ginsnop = ginsnop
	arch.Ginsnopdefer = ginsnop

	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
