package main

import (
	"flag"

	"github.com/sarchlab/mgpusim/v3/benchmarks/amdappsdk/matrixtranspose"
	"github.com/sarchlab/mgpusim/v3/samples/runner"
)

var dataWidth = flag.Int("width", 256, "The dimension of the square matrix.")

func main() {
	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	// 3gb
	benchmark1 := matrixtranspose.NewBenchmark(runner.Driver())
	benchmark1.Width = 500
	// benchmark.Width = *dataWidth
	runner.AddBenchmark(benchmark1)

	// 2gb
	benchmark := matrixtranspose.NewBenchmark(runner.Driver())
	benchmark.Width = 100
	runner.AddBenchmark(benchmark)

	runner.Run()
	// runner.PauseBenchMark(benchmark1)
}
