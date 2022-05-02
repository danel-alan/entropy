package entropy

import (
	"math"
	"runtime"
)

// Shannon returns the entropy of a slice of bytes by the Shannon formula.
func Shannon(s []byte) float64 {
	dictionary := make(map[byte]uint64)
	for _, b := range s {
		dictionary[b]++
	}
	var entropy float64
	for _, count := range dictionary {
		f := float64(count) / float64(len(s))
		entropy += -f * math.Log2(f)
	}
	return math.Round(entropy*100)/100
}

// ShannonAll returns a slice of entropies from of a matrix of bytes using the Shannon formula.
func ShannonAll(blocks [][]byte) []float64 {
	entropies := make([]float64, 0, len(blocks))
	for _, block := range blocks {
		entropies = append(entropies, Shannon(block))
	}
	return entropies
}

var numCPU = float64(runtime.NumCPU())

type future struct {
	res []float64
	idx int
}

// ShannonAllBatch returns the ShannonAll result using all the cpus available or all the bytes arrs
func ShannonAllBatch(blocks [][]byte) []float64 {
	blocksSize := len(blocks)
	if len(blocks) <= 1 {
		return ShannonAll(blocks)
	}
	byteBaches := batches(blocks, int(math.Min(float64(blocksSize), numCPU)))
	batchSize := len(byteBaches)
	futuresChan := make(chan future, batchSize)
	for i, b := range byteBaches { // Scatter batches
		go func(idx int, batch [][]byte) {
			futuresChan <- future{
				res: ShannonAll(batch),
				idx: idx,
			}
		}(i, b)
	}
	results := make([][]float64, batchSize)
	for i := 0; i < batchSize; i++ { // Gather batches by order
		result := <-futuresChan
		results[result.idx] = result.res
	}
	entropies := make([]float64, 0, blocksSize)
	for _, e := range results {
		entropies = append(entropies, e...)
	}
	return entropies
}

// batches divide a slice of bytes into many slices by batchSize.
// 	helper func from https://github.com/golang/go/wiki/SliceTricks#batching-with-minimal-allocation
func batches(s [][]byte, batchSize int) [][][]byte {
	batches := make([][][]byte, 0, (len(s)+batchSize-1)/batchSize)
	for batchSize < len(s) {
		s, batches = s[batchSize:], append(batches, s[0:batchSize:batchSize])
	}
	return append(batches, s)
}
