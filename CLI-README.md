  -akitartm-port int
        Custom port to host AkitaRTM. A 4-digit or 5-digit port number is required. If 
        this number is not given or a invalid number is given number, a random port 
        will be used.
  -analyzer-Name string
        The name of the analyzer to use.
  -analyzer-period float
        The period to dump the analyzer results.
  -buffer-level-trace-dir string
        The directory to dump the buffer level traces.
  -buffer-level-trace-period float
        The period to dump the buffer level trace.
  -clusters int
        The number of clusters. (default 5)
  -debug-isa
        Generate the ISA debugging file.
  -features int
        The number of features for each point. (default 32)
  -gpus string
        The GPUs to use, use a format like 1,2,3,4. By default, GPU 1 is used.
  -length int
        The number of samples to filter. (default 4096)
  -magic-memory-copy
        Copy data from CPU directly to global memory
  -max-inst uint
        Terminate the simulation after the given number of instructions is retired.
  -max-iter int
        The maximum number of iterations to run (default 5)
  -metric-file-name string
        Modify the name of the output csv file. (default "metrics")
  -parallel
        Run the simulation in parallel.
  -points int
        The number of points. (default 1024)
  -report-all
        Report all metrics to .csv file.
  -report-busy-time
        Report SIMD Unit's busy time
  -report-cache-hit-rate
        Report the cache hit rate of each cache.
  -report-cache-latency
        Report the average cache latency.
  -report-cpi-stack
        Report CPI stack
  -report-dram-transaction-count
        Report the number of transactions accessing the DRAMs.
  -report-inst-count
        Report the number of instructions executed in each compute unit.
  -report-rdma-transaction-count
        Report the number of transactions going through the RDMA engines.
  -report-tlb-hit-rate
        Report the TLB hit rate of each TLB.
  -timing
        Run detailed timing simulation.
  -trace-mem
        Generate memory trace
  -trace-vis
        Generate trace for visualization purposes.
  -trace-vis-db string
        The database to store the visualization trace. Possible values are sqlite, mysql, and csv. (default "sqlite")
  -trace-vis-db-file string
        The file name of the database to store the visualization trace. Extension names are not required. If not specified, a random file name will be used. This flag does not work with Mysql db. When MySQL is used, the database name is always randomly generated.
  -trace-vis-end float
        The end time of collecting visualization traces. A negative numbermeans that the trace will be collected to the end of the simulation. (default -1)
  -trace-vis-start float
        The starting time to collect visualization traces. A negative number represents starting from the beginning. (default -1)
  -unified-gpus string
        Run multi-GPU benchmark in a unified mode.
        Use a format like 1,2,3,4. Cannot coexist with -gpus.
  -use-unified-memory
        Run benchmark with Unified Memory or not
  -verify
        Verify the emulation result.
  -width int
        The dimension of the square matrix. (default 256)