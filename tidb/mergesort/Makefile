.PHONY: all

all: test bench

test:
	go test

bench:
	go test -bench BenchmarkMergeSort -run xx -count 5 -benchmem 

prof:
	go test -bench BenchmarkMergeSort -run xx -count 5 -benchmem -cpuprofile cpu.prof

trace:
	go test -bench BenchmarkMergeSort -run xx -trace trace.out

