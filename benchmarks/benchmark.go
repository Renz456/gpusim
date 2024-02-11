// Package benchmarks defines Benchmark interface.
package benchmarks

import "github.com/sarchlab/mgpusim/v3/driver"

// A Benchmark is a GPU program that can run on the GCN3 simulator
type Benchmark interface {
	SelectGPU(gpuIDs []int)
	Run()
	Verify()
	SetUnifiedMemory()
	GetQueue() *driver.CommandQueue
	GetContext() *driver.Context
}
