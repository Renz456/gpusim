package emu

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/akita/gcn3/insts"
)

var _ = Describe("ALU", func() {

	var (
		alu   *ALUImpl
		state *mockInstState
	)

	BeforeEach(func() {
		alu = NewALUImpl(nil)

		state = new(mockInstState)
		state.scratchpad = make([]byte, 4096)
	})

	It("should run V_ADD_U32 VOP3b", func() {
		state.inst = insts.NewInst()
		state.inst.FormatType = insts.VOP3b
		state.inst.Opcode = 281

		sp := state.Scratchpad().AsVOP3B()
		sp.SRC0[0] = 1
		sp.SRC1[0] = 2
		sp.SRC0[1] = 0xffffffff
		sp.SRC1[1] = 2
		sp.EXEC = 3

		alu.Run(state)

		Expect(sp.DST[0]).To(Equal(uint64(3)))
		Expect(sp.DST[1] & 0xffffffff).To(Equal(uint64(1)))
		Expect(sp.VCC).To(Equal(uint64(0x2)))
	})

	It("should run V_ADDC_U32", func() {
		state.inst = insts.NewInst()
		state.inst.FormatType = insts.VOP3b
		state.inst.Opcode = 284

		sp := state.scratchpad.AsVOP3B()
		sp.SRC0[0] = uint64(0xfffffffd)
		sp.SRC1[0] = uint64(2)
		sp.SRC2[0] = uint64(1)
		sp.SRC0[1] = uint64(0xfffffffd)
		sp.SRC1[1] = uint64(1)
		sp.SRC2[1] = uint64(1)
		sp.EXEC = 0x3

		alu.Run(state)

		Expect(sp.DST[0]).To(Equal(uint64(0)))
		Expect(sp.DST[1]).To(Equal(uint64(0xfffffffe)))
		Expect(sp.SDST).To(Equal(uint64(1)))
	})

})
