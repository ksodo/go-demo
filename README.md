Demo go testing and benchmarking.

Build. Always rebuild after changing source code.

    $ go build

The app is trivial. It only prints response for btcusdt ticker.

    $ ./go-demo
    &{0 Operation successful {1581022816 {9307.7 9308.22 9305.41 9339.49 9309.59 74.351373}}}


Test
----

Use `_test` suffix to name test file. Define test functions with `Test` prefix.
Run test to see example failure messages.

    $ go test -v
    === RUN   TestGetTickersFail
    --- FAIL: TestGetTickersFail (0.27s)
            demo_test.go:13: 0 != -2
            demo_test.go:18: Operation successful != {"error":{"code":1001,"message":"market does not have a valid value"}}
    === RUN   TestGetTickersFailData
    --- FAIL: TestGetTickersFailData (0.10s)
            demo_test.go:33: expected: 0.000000 > 0
    === RUN   TestGetTickersPass
    --- PASS: TestGetTickersPass (0.06s)
    FAIL
    exit status 1
    FAIL    _/home/karma/dev/go-demo        0.434s


Benchmark
---------

Run benchmark. `-run` option skips tests since benchmark won't run with failing
tests. Second column is number of loop iterations.

    $ go test -run=X -bench .
		goos: linux
		goarch: amd64
		BenchmarkTickersAlloc-6         30000000               148 ns/op
		BenchmarkGetTickers-6                 20          58059016 ns/op
		PASS
		ok      _/home/karma/dev/go-demo        6.198s

Use `-cpu` to schedule on an idle cpu to get better results.

		$ go test -run=X -bench . -cpu 3
		goos: linux
		goarch: amd64
		BenchmarkTickersAlloc-3         30000000                43.6 ns/op
		BenchmarkGetTickers-3                 20          56690830 ns/op
		PASS
		ok      _/home/karma/dev/go-demo        3.060s

Cpu profiling is built in.

		$ go test -run=X -bench . -cpu 3 -cpuprofile cpu.out

Text report of cpu profile. Graphical reports are supported. Below it is
expected that runtime.mallocgc dominates cpu since BenchmarkTickersAlloc is
measuring allocation.

		$ go tool pprof cpu.out
		File: go-demo.test
		Type: cpu
		Time: Feb 6, 2020 at 4:36pm (EST)
		Duration: 3.49s, Total samples = 1.83s (52.48%)
		Entering interactive mode (type "help" for commands, "o" for options)
		(pprof) top
		Showing nodes accounting for 1310ms, 71.58% of 1830ms total
		Showing top 10 nodes out of 147
					flat  flat%   sum%        cum   cum%
				 410ms 22.40% 22.40%      980ms 53.55%  runtime.mallocgc
				 280ms 15.30% 37.70%      280ms 15.30%  runtime.heapBitsSetType
				 210ms 11.48% 49.18%      280ms 15.30%  runtime.scanobject
