package tensor

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CPUOperator", func() {
	var (
		to CPUOperator
	)

	It("should transpose", func() {
		in := to.CreateWithData(
			[]float64{
				15, 18, 9, 30, 36, 18, 25, 30, 15,
				15, 18, 9, 30, 36, 18, 25, 30, 15,
			},

			[]int{1, 2, 3, 3}, "CNHW")

		outData := []float32{
			15, 18, 9, 30, 36, 18, 25, 30, 15,
			15, 18, 9, 30, 36, 18, 25, 30, 15,
		}

		out := to.Transpose(in, []int{1, 0, 2, 3})

		outV := out.Vector()
		for i := range outData {
			Expect(outV[i]).To(BeNumerically("~", outData[i], 1e-3))
		}
		Expect(out.Descriptor()).To(Equal("NCHW"))
	})

	It("should roate 180", func() {
		in := to.CreateWithData(
			[]float64{
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 10, 11, 12,
			},
			[]int{1, 1, 3, 4}, "")

		out := to.Rotate180(in)

		Expect(out.Vector()).To(Equal(
			[]float64{
				12, 11, 10, 9,
				8, 7, 6, 5,
				4, 3, 2, 1,
			}))
	})

	It("should roate 180, test 2", func() {
		in := to.CreateWithData(
			[]float64{
				1.111, 1.112, 1.113, 1.114,
				1.121, 1.122, 1.123, 1.124,
				1.131, 1.132, 1.133, 1.134,

				1.211, 1.212, 1.213, 1.214,
				1.221, 1.222, 1.223, 1.224,
				1.231, 1.232, 1.233, 1.234,

				1.311, 1.312, 1.313, 1.314,
				1.321, 1.322, 1.323, 1.324,
				1.331, 1.332, 1.333, 1.334,

				2.111, 2.112, 2.113, 2.114,
				2.121, 2.122, 2.123, 2.124,
				2.131, 2.132, 2.133, 2.134,

				2.211, 2.212, 2.213, 2.214,
				2.221, 2.222, 2.223, 2.224,
				2.231, 2.232, 2.233, 2.234,

				2.311, 2.312, 2.313, 2.314,
				2.321, 2.322, 2.323, 2.324,
				2.331, 2.332, 2.333, 2.334,
			}, []int{2, 3, 3, 4}, "")

		out := to.Rotate180(in)

		goldOut := []float64{
			1.134, 1.133, 1.132, 1.131,
			1.124, 1.123, 1.122, 1.121,
			1.114, 1.113, 1.112, 1.111,

			1.234, 1.233, 1.232, 1.231,
			1.224, 1.223, 1.222, 1.221,
			1.214, 1.213, 1.212, 1.211,

			1.334, 1.333, 1.332, 1.331,
			1.324, 1.323, 1.322, 1.321,
			1.314, 1.313, 1.312, 1.311,

			2.134, 2.133, 2.132, 2.131,
			2.124, 2.123, 2.122, 2.121,
			2.114, 2.113, 2.112, 2.111,

			2.234, 2.233, 2.232, 2.231,
			2.224, 2.223, 2.222, 2.221,
			2.214, 2.213, 2.212, 2.211,

			2.334, 2.333, 2.332, 2.331,
			2.324, 2.323, 2.322, 2.321,
			2.314, 2.313, 2.312, 2.311,
		}

		outV := out.Vector()
		for i := 0; i < len(goldOut); i++ {
			Expect(outV[i]).To(BeNumerically("~", goldOut[i], 1e-3))
		}
	})

	It("should dilate", func() {
		in := to.CreateWithData(
			[]float64{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			[]int{1, 1, 3, 3}, "")

		out := to.Dilate(in, []int{3, 2})

		goldOut := []float64{
			1, 0, 2, 0, 3,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			4, 0, 5, 0, 6,
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			7, 0, 8, 0, 9,
		}

		outV := out.Vector()
		for i := 0; i < len(goldOut); i++ {
			Expect(outV[i]).To(BeNumerically("~", goldOut[i], 1e-3))
		}

	})

	It("should do general transpose", func() {
		in := to.CreateWithData(
			[]float64{
				1.111, 1.112, 1.113,
				1.121, 1.122, 1.123,
				1.131, 1.132, 1.133,

				1.211, 1.212, 1.213,
				1.221, 1.222, 1.223,
				1.231, 1.232, 1.233,

				1.311, 1.312, 1.313,
				1.321, 1.322, 1.323,
				1.331, 1.332, 1.333,

				1.411, 1.412, 1.413,
				1.421, 1.422, 1.423,
				1.431, 1.432, 1.433,

				2.111, 2.112, 2.113,
				2.121, 2.122, 2.123,
				2.131, 2.132, 2.133,

				2.211, 2.212, 2.213,
				2.221, 2.222, 2.223,
				2.231, 2.232, 2.233,

				2.311, 2.312, 2.313,
				2.321, 2.322, 2.323,
				2.331, 2.332, 2.333,

				2.411, 2.412, 2.413,
				2.421, 2.422, 2.423,
				2.431, 2.432, 2.433,
			},

			[]int{2, 4, 3, 3}, "CNHW")

		outData := []float32{
			1.111, 1.112, 1.113,
			1.121, 1.122, 1.123,
			1.131, 1.132, 1.133,

			2.111, 2.112, 2.113,
			2.121, 2.122, 2.123,
			2.131, 2.132, 2.133,

			1.211, 1.212, 1.213,
			1.221, 1.222, 1.223,
			1.231, 1.232, 1.233,

			2.211, 2.212, 2.213,
			2.221, 2.222, 2.223,
			2.231, 2.232, 2.233,

			1.311, 1.312, 1.313,
			1.321, 1.322, 1.323,
			1.331, 1.332, 1.333,

			2.311, 2.312, 2.313,
			2.321, 2.322, 2.323,
			2.331, 2.332, 2.333,

			1.411, 1.412, 1.413,
			1.421, 1.422, 1.423,
			1.431, 1.432, 1.433,

			2.411, 2.412, 2.413,
			2.421, 2.422, 2.423,
			2.431, 2.432, 2.433,
		}

		out := to.Transpose(in, []int{1, 0, 2, 3})
		outV := out.Vector()

		for i := range outData {
			Expect(outV[i]).To(BeNumerically("~", outData[i], 1e-3))
		}
		Expect(out.Descriptor()).To(Equal("NCHW"))
	})

	It("should do im2col", func() {
		input := to.CreateWithData(
			[]float64{
				1, 1, 1,
				2, 2, 2,
				3, 3, 3,
			},
			[]int{1, 1, 3, 3},
			"NCHW",
		)
		kernelSize := []int{3, 3}
		padding := []int{1, 1}
		stride := []int{1, 1}
		dilation := []int{1, 1}

		output := to.Im2Col(input, kernelSize, padding, stride, dilation)

		Expect(output.Size()).To(Equal([]int{9, 9}))
		Expect(output.Vector()).To(Equal([]float64{
			0, 0, 0, 0, 1, 1, 0, 2, 2,
			0, 0, 0, 1, 1, 1, 2, 2, 2,
			0, 0, 0, 1, 1, 0, 2, 2, 0,
			0, 1, 1, 0, 2, 2, 0, 3, 3,
			1, 1, 1, 2, 2, 2, 3, 3, 3,
			1, 1, 0, 2, 2, 0, 3, 3, 0,
			0, 2, 2, 0, 3, 3, 0, 0, 0,
			2, 2, 2, 3, 3, 3, 0, 0, 0,
			2, 2, 0, 3, 3, 0, 0, 0, 0,
		}))
	})

	It("should do im2col dilation", func() {
		input := to.CreateWithData(
			[]float64{
				1.1, 1.2, 1.3,
				2.1, 2.2, 2.3,
				3.1, 3.0, 3.3,
			},
			[]int{1, 1, 3, 3},
			"NCHW",
		)
		kernelSize := []int{3, 3}
		padding := []int{1, 1}
		stride := []int{1, 1}
		dilation := []int{2, 2}

		output := to.Im2Col(input, kernelSize, padding, stride, dilation)

		Expect(output.Size()).To(Equal([]int{9, 1}))
		Expect(output.Vector()).To(Equal([]float64{
			0.0,
			0.0,
			0.0,
			0.0,
			2.2,
			0.0,
			0.0,
			0.0,
			0.0,
		}))
	})

	It("should do im2col batch & channel", func() {
		input := to.CreateWithData(
			[]float64{
				1111, 1112, 1113, 1121, 1122, 1123, 1131, 1132, 1133,
				1211, 1212, 1213, 1221, 1222, 1223, 1231, 1232, 1233,
				1311, 1312, 1313, 1321, 1322, 1323, 1331, 1332, 1333,
				2111, 2112, 2113, 2121, 2122, 2123, 2131, 2132, 2133,
				2211, 2212, 2213, 2221, 2222, 2223, 2231, 2232, 2233,
				3311, 3312, 3313, 3321, 3322, 3323, 3331, 3332, 3333,
			},
			[]int{2, 3, 3, 3},
			"NCHW",
		)
		kernelSize := []int{3, 3}
		padding := []int{1, 1}
		stride := []int{1, 1}
		dilation := []int{1, 1}

		output := to.Im2Col(input, kernelSize, padding, stride, dilation)

		Expect(output.Size()).To(Equal([]int{27, 18}))
		Expect(output.Vector()).To(Equal([]float64{
			0, 0, 0, 0, 1111, 1112, 0, 1121, 1122, 0, 0, 0, 0, 2111, 2112, 0, 2121, 2122,
			0, 0, 0, 1111, 1112, 1113, 1121, 1122, 1123, 0, 0, 0, 2111, 2112, 2113, 2121, 2122, 2123,
			0, 0, 0, 1112, 1113, 0, 1122, 1123, 0, 0, 0, 0, 2112, 2113, 0, 2122, 2123, 0,
			0, 1111, 1112, 0, 1121, 1122, 0, 1131, 1132, 0, 2111, 2112, 0, 2121, 2122, 0, 2131, 2132,
			1111, 1112, 1113, 1121, 1122, 1123, 1131, 1132, 1133, 2111, 2112, 2113, 2121, 2122, 2123, 2131, 2132, 2133,
			1112, 1113, 0, 1122, 1123, 0, 1132, 1133, 0, 2112, 2113, 0, 2122, 2123, 0, 2132, 2133, 0,
			0, 1121, 1122, 0, 1131, 1132, 0, 0, 0, 0, 2121, 2122, 0, 2131, 2132, 0, 0, 0,
			1121, 1122, 1123, 1131, 1132, 1133, 0, 0, 0, 2121, 2122, 2123, 2131, 2132, 2133, 0, 0, 0,
			1122, 1123, 0, 1132, 1133, 0, 0, 0, 0, 2122, 2123, 0, 2132, 2133, 0, 0, 0, 0,
			0, 0, 0, 0, 1211, 1212, 0, 1221, 1222, 0, 0, 0, 0, 2211, 2212, 0, 2221, 2222,
			0, 0, 0, 1211, 1212, 1213, 1221, 1222, 1223, 0, 0, 0, 2211, 2212, 2213, 2221, 2222, 2223,
			0, 0, 0, 1212, 1213, 0, 1222, 1223, 0, 0, 0, 0, 2212, 2213, 0, 2222, 2223, 0,
			0, 1211, 1212, 0, 1221, 1222, 0, 1231, 1232, 0, 2211, 2212, 0, 2221, 2222, 0, 2231, 2232,
			1211, 1212, 1213, 1221, 1222, 1223, 1231, 1232, 1233, 2211, 2212, 2213, 2221, 2222, 2223, 2231, 2232, 2233,
			1212, 1213, 0, 1222, 1223, 0, 1232, 1233, 0, 2212, 2213, 0, 2222, 2223, 0, 2232, 2233, 0,
			0, 1221, 1222, 0, 1231, 1232, 0, 0, 0, 0, 2221, 2222, 0, 2231, 2232, 0, 0, 0,
			1221, 1222, 1223, 1231, 1232, 1233, 0, 0, 0, 2221, 2222, 2223, 2231, 2232, 2233, 0, 0, 0,
			1222, 1223, 0, 1232, 1233, 0, 0, 0, 0, 2222, 2223, 0, 2232, 2233, 0, 0, 0, 0,
			0, 0, 0, 0, 1311, 1312, 0, 1321, 1322, 0, 0, 0, 0, 3311, 3312, 0, 3321, 3322,
			0, 0, 0, 1311, 1312, 1313, 1321, 1322, 1323, 0, 0, 0, 3311, 3312, 3313, 3321, 3322, 3323,
			0, 0, 0, 1312, 1313, 0, 1322, 1323, 0, 0, 0, 0, 3312, 3313, 0, 3322, 3323, 0,
			0, 1311, 1312, 0, 1321, 1322, 0, 1331, 1332, 0, 3311, 3312, 0, 3321, 3322, 0, 3331, 3332,
			1311, 1312, 1313, 1321, 1322, 1323, 1331, 1332, 1333, 3311, 3312, 3313, 3321, 3322, 3323, 3331, 3332, 3333,
			1312, 1313, 0, 1322, 1323, 0, 1332, 1333, 0, 3312, 3313, 0, 3322, 3323, 0, 3332, 3333, 0,
			0, 1321, 1322, 0, 1331, 1332, 0, 0, 0, 0, 3321, 3322, 0, 3331, 3332, 0, 0, 0,
			1321, 1322, 1323, 1331, 1332, 1333, 0, 0, 0, 3321, 3322, 3323, 3331, 3332, 3333, 0, 0, 0,
			1322, 1323, 0, 1332, 1333, 0, 0, 0, 0, 3322, 3323, 0, 3332, 3333, 0, 0, 0, 0,
		}))
	})

	It("should do maxpooling forward", func() {
		inTensor := to.CreateWithData([]float64{
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
		}, []int{2, 3, 6, 6}, "NCHW")

		out, mask := to.MaxPoolingForward(inTensor,
			[]int{2, 2}, []int{1, 1}, []int{2, 2})

		goldOut := []float64{
			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,

			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,

			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,

			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,

			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,

			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,
			1, 3, 5, 6,
		}

		goldMask := []float64{
			0.000000, 2.000000, 4.000000, 5.000000,
			6.000000, 8.000000, 10.000000, 11.000000,
			18.000000, 20.000000, 22.000000, 23.000000,
			30.000000, 32.000000, 34.000000, 35.000000,

			36.000000, 38.000000, 40.000000, 41.000000,
			42.000000, 44.000000, 46.000000, 47.000000,
			54.000000, 56.000000, 58.000000, 59.000000,
			66.000000, 68.000000, 70.000000, 71.000000,

			72.000000, 74.000000, 76.000000, 77.000000,
			78.000000, 80.000000, 82.000000, 83.000000,
			90.000000, 92.000000, 94.000000, 95.000000,
			102.000000, 104.000000, 106.000000, 107.000000,

			108.000000, 110.000000, 112.000000, 113.000000,
			114.000000, 116.000000, 118.000000, 119.000000,
			126.000000, 128.000000, 130.000000, 131.000000,
			138.000000, 140.000000, 142.000000, 143.000000,

			144.000000, 146.000000, 148.000000, 149.000000,
			150.000000, 152.000000, 154.000000, 155.000000,
			162.000000, 164.000000, 166.000000, 167.000000,
			174.000000, 176.000000, 178.000000, 179.000000,

			180.000000, 182.000000, 184.000000, 185.000000,
			186.000000, 188.000000, 190.000000, 191.000000,
			198.000000, 200.000000, 202.000000, 203.000000,
			210.000000, 212.000000, 214.000000, 215.000000,
		}

		Expect(out.Size()).To(Equal([]int{2, 3, 4, 4}))
		outV := out.Vector()
		for i := range goldOut {
			Expect(outV[i]).To(BeNumerically("~", goldOut[i], 1e-3))
		}

		maskV := mask.Vector()
		for i := range goldMask {
			Expect(int(maskV[i])).To(BeNumerically("~", goldMask[i], 1e-3))
		}
	})

	It("should do maxpooling backward", func() {
		forwardInTensor := to.CreateWithData([]float64{
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
		}, []int{2, 3, 6, 6}, "NCHW")
		backwardInTensor := to.CreateWithData([]float64{
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,

			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,

			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,

			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,

			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,

			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
		}, []int{2, 3, 4, 4}, "NCHW")
		maskTensor := to.CreateWithData(
			[]float64{
				0.000000, 2.000000, 4.000000, 5.000000,
				6.000000, 8.000000, 10.000000, 11.000000,
				18.000000, 20.000000, 22.000000, 23.000000,
				30.000000, 32.000000, 34.000000, 35.000000,

				36.000000, 38.000000, 40.000000, 41.000000,
				42.000000, 44.000000, 46.000000, 47.000000,
				54.000000, 56.000000, 58.000000, 59.000000,
				66.000000, 68.000000, 70.000000, 71.000000,

				72.000000, 74.000000, 76.000000, 77.000000,
				78.000000, 80.000000, 82.000000, 83.000000,
				90.000000, 92.000000, 94.000000, 95.000000,
				102.000000, 104.000000, 106.000000, 107.000000,

				108.000000, 110.000000, 112.000000, 113.000000,
				114.000000, 116.000000, 118.000000, 119.000000,
				126.000000, 128.000000, 130.000000, 131.000000,
				138.000000, 140.000000, 142.000000, 143.000000,

				144.000000, 146.000000, 148.000000, 149.000000,
				150.000000, 152.000000, 154.000000, 155.000000,
				162.000000, 164.000000, 166.000000, 167.000000,
				174.000000, 176.000000, 178.000000, 179.000000,

				180.000000, 182.000000, 184.000000, 185.000000,
				186.000000, 188.000000, 190.000000, 191.000000,
				198.000000, 200.000000, 202.000000, 203.000000,
				210.000000, 212.000000, 214.000000, 215.000000,
			}, []int{2, 3, 4, 4}, "NCHW")

		out := to.MaxPoolingBackward(
			forwardInTensor, backwardInTensor, maskTensor,
			[]int{2, 2}, []int{1, 1}, []int{2, 2})

		goldOut := []float64{
			1, 0, 2, 0, 3, 4,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,

			1, 0, 2, 0, 3, 4,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,

			1, 0, 2, 0, 3, 4,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,

			1, 0, 2, 0, 3, 4,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,

			1, 0, 2, 0, 3, 4,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,

			1, 0, 2, 0, 3, 4,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,
			0, 0, 0, 0, 0, 0,
			1, 0, 2, 0, 3, 4,
		}

		Expect(out.Size()).To(Equal([]int{2, 3, 6, 6}))
		outV := out.Vector()
		for i := range goldOut {
			Expect(outV[i]).To(BeNumerically("~", goldOut[i], 1e-3))
		}
	})

	It("should do avgpooling forward", func() {
		inTensor := to.CreateWithData([]float64{
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
		}, []int{2, 3, 6, 6}, "NCHW")

		out := to.AvgPoolingForward(inTensor,
			[]int{2, 2}, []int{1, 1}, []int{2, 2})

		goldOut := []float64{
			0.25, 1.25, 2.25, 1.50,
			0.50, 2.50, 4.50, 3.00,
			0.50, 2.50, 4.50, 3.00,
			0.25, 1.25, 2.25, 1.50,

			0.25, 1.25, 2.25, 1.50,
			0.50, 2.50, 4.50, 3.00,
			0.50, 2.50, 4.50, 3.00,
			0.25, 1.25, 2.25, 1.50,

			0.25, 1.25, 2.25, 1.50,
			0.50, 2.50, 4.50, 3.00,
			0.50, 2.50, 4.50, 3.00,
			0.25, 1.25, 2.25, 1.50,

			0.25, 1.25, 2.25, 1.50,
			0.50, 2.50, 4.50, 3.00,
			0.50, 2.50, 4.50, 3.00,
			0.25, 1.25, 2.25, 1.50,

			0.25, 1.25, 2.25, 1.50,
			0.50, 2.50, 4.50, 3.00,
			0.50, 2.50, 4.50, 3.00,
			0.25, 1.25, 2.25, 1.50,

			0.25, 1.25, 2.25, 1.50,
			0.50, 2.50, 4.50, 3.00,
			0.50, 2.50, 4.50, 3.00,
			0.25, 1.25, 2.25, 1.50,
		}

		Expect(out.Size()).To(Equal([]int{2, 3, 4, 4}))
		outV := out.Vector()
		for i := range goldOut {
			Expect(outV[i]).To(BeNumerically("~", goldOut[i], 1e-3))
		}
	})

	It("should do avgpooling backward", func() {
		forwardInTensor := to.CreateWithData([]float64{
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,

			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
			1, 2, 3, 4, 5, 6,
		}, []int{2, 3, 6, 6}, "NCHW")
		backwardInTensor := to.CreateWithData([]float64{
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,

			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,

			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,

			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,

			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,

			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
		}, []int{2, 3, 4, 4}, "NCHW")

		out := to.AvgPoolingBackward(
			forwardInTensor, backwardInTensor,
			[]int{2, 2}, []int{1, 1}, []int{2, 2},
		)

		goldOut := []float64{
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,

			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,

			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,

			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,

			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,

			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
			0.25, 0.50, 0.50, 0.75, 0.75, 1.00,
		}

		Expect(out.Size()).To(Equal([]int{2, 3, 6, 6}))
		outV := out.Vector()
		for i := range goldOut {
			Expect(outV[i]).To(BeNumerically("~", goldOut[i], 1e-3))
		}
	})
})
