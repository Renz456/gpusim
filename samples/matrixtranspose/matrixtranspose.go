package main

import (
	"flag"

	"github.com/sarchlab/mgpusim/v3/benchmarks/amdappsdk/matrixtranspose"
	"github.com/sarchlab/mgpusim/v3/samples/runner"
)

var dataWidth = flag.Int("width", 256, "The dimension of the square matrix.")

var numData = flag.Int("length", 4096, "The number of samples to filter.")

var points = flag.Int("points", 1024, "The number of points.")
var clusters = flag.Int("clusters", 5, "The number of clusters.")
var features = flag.Int("features", 32,
	"The number of features for each point.")
var maxIter = flag.Int("max-iter", 5,
	"The maximum number of iterations to run")

func main() {
	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	// 3gb

	// benchmark_kmeans := kmeans.NewBenchmark(runner.Driver())
	// benchmark_kmeans.NumPoints = *points
	// benchmark_kmeans.NumClusters = *clusters
	// benchmark_kmeans.NumFeatures = *features
	// benchmark_kmeans.MaxIter = *maxIter

	// runner.AddBenchmark(benchmark_kmeans)

	benchmark1 := matrixtranspose.NewBenchmark(runner.Driver())
	benchmark1.Width = 80
	// benchmark.Width = *dataWidth
	runner.AddBenchmark(benchmark1)
	// 2gb
	benchmark := matrixtranspose.NewBenchmark(runner.Driver())
	benchmark.Width = 100
	runner.AddBenchmark(benchmark)

	runner.Run()
	// runner.PauseBenchMark(benchmark_kmeans)
}
