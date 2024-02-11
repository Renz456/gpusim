package main

import (
	"flag"
	"fmt"

	"github.com/sarchlab/mgpusim/v3/benchmarks/heteromark/kmeans"
	"github.com/sarchlab/mgpusim/v3/samples/runner"
)

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

	// benchmark := fir.NewBenchmark(runner.Driver())
	// benchmark.Length = *numData

	// runner.AddBenchmark(benchmark)

	benchmark_kmeans := kmeans.NewBenchmark(runner.Driver())
	benchmark_kmeans.NumPoints = *points
	benchmark_kmeans.NumClusters = *clusters
	benchmark_kmeans.NumFeatures = *features
	benchmark_kmeans.MaxIter = *maxIter

	runner.AddBenchmark(benchmark_kmeans)

	benchmark_kmeans = kmeans.NewBenchmark(runner.Driver())
	benchmark_kmeans.NumPoints = *points
	benchmark_kmeans.NumClusters = *clusters
	benchmark_kmeans.NumFeatures = *features
	benchmark_kmeans.MaxIter = *maxIter

	runner.AddBenchmark(benchmark_kmeans)
	// benchmark = fir.NewBenchmark(runner.Driver())
	// benchmark.Length = *numData

	// runner.AddBenchmark(benchmark)
	fmt.Println("GPUs", runner.GPUIDs)
	runner.Run()
}
