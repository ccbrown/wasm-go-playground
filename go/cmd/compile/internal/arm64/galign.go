// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arm64

import (
	"github.com/ccbrown/go-web-gc/go/cmd/compile/internal/gc"
	"github.com/ccbrown/go-web-gc/go/cmd/compile/internal/ssa"
	"github.com/ccbrown/go-web-gc/go/cmd/internal/obj/arm64"
)

func Init(arch *gc.Arch) {
	arch.LinkArch = &arm64.Linkarm64
	arch.REGSP = arm64.REGSP
	arch.MAXWIDTH = 1 << 50

	arch.PadFrame = padframe
	arch.ZeroRange = zerorange
	arch.ZeroAuto = zeroAuto
	arch.Ginsnop = ginsnop
	arch.Ginsnopdefer = ginsnop

	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
