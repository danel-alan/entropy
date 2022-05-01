package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/danel-alan/entropy/pkg/reporting"
)

var (
	blockSize = flag.Uint64("size", 1024, "block size to analize")
	high      = flag.Float64("high", 7, "threshold for counting blocks as high entropy")
	low       = flag.Float64("low", 2, "threshold for counting blocks as low entropy")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: entropy-cli [flags] [path ...]\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: no file path found")
		return
	}

	for _, arg := range args {
		switch info, err := os.Stat(arg); {
		case err != nil:
			fmt.Fprintln(os.Stderr, fmt.Errorf("file error: %v", err))
			return
		case info.IsDir():
			fmt.Fprintln(os.Stderr, fmt.Errorf("%v is a directory", info.Name()))
		default:
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("file error: %v", err))
				return
			}
			r := reporting.EntropyReporter{
				HighEntropy: *high,
				LowEntropy:  *low,
			}
			report, err := r.Report(f, *blockSize)
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("calculation error: %v", err))
				return
			}
			result, _ := json.MarshalIndent(report, "", "\t")
			fmt.Fprintln(os.Stdout, string(result))
		}
	}
}
