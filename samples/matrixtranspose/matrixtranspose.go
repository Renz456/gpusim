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
	benchmark := matrixtranspose.NewBenchmark(runner.Driver())
	benchmark.Width = 100
	// benchmark.Width = *dataWidth
	runner.AddBenchmark(benchmark)

	// 2gb
	// benchmark = matrixtranspose.NewBenchmark(runner.Driver())
	// benchmark.Width = 1500
	// runner.AddBenchmark(benchmark)

	runner.Run()
}
