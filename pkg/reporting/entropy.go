package reporting

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/danel-alan/entropy/pkg/entropy"
)

type (
	EntropyReporter struct {
		DefaultBlockSize uint64
		HighEntropy      float64
		LowEntropy       float64
	}
	EntropyReport struct {
		EntropyDetail []float64      `json:"entropy_detail"`
		Summary       EntropySummary `json:"summary"`
	}
	EntropySummary struct {
		LowEntropyBlocks  int `json:"low_entropy_blocks"`
		HighEntropyBlocks int `json:"high_entropy_blocks"`
	}
)

func (er *EntropyReporter) Report(r io.Reader, size uint64) (*EntropyReport, error) {
	if size == 0 {
		size = er.DefaultBlockSize
	}
	blocks, err := blocks(r, size)
	if err != nil {
		return nil, fmt.Errorf("reporting: error reading past %v blocks: %v", len(blocks), err)
	}
	entropies := entropy.ShannonAllBatch(blocks)
	var summary EntropySummary
	for _, entropy := range entropies {
		if entropy > er.HighEntropy {
			summary.HighEntropyBlocks++
		}
		if entropy < er.LowEntropy {
			summary.LowEntropyBlocks++
		}
	}
	return &EntropyReport{
		EntropyDetail: entropies,
		Summary:       summary,
	}, nil
}

// blocks divides the bytes from a reader into blocks of bytes of size n.
func blocks(r io.Reader, size uint64) ([][]byte, error) {
	blocks := make([][]byte, 0)
loop:
	for {
		block := make([]byte, size)
		if bReads, err := io.ReadFull(r, block); err != nil {
			switch err {
			case io.EOF:
				break loop
			case io.ErrUnexpectedEOF:
				block = block[:bReads]
			default:
				return blocks, err
			}
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

func (r *EntropyReport) String() string {
	res, _ := json.MarshalIndent(r, "", "\t")
	return string(res)
}
