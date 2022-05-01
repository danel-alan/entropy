package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/danel-alan/entropy/pkg/http/rest"
	"github.com/danel-alan/entropy/pkg/reporting"
)

var (
	defaultBlockSize = flag.Uint64("def_size", 1024, "default size for a block")
	high             = flag.Float64("high", 7, "threshold for counting blocks as high entropy")
	low              = flag.Float64("low", 2, "threshold for counting blocks as low entropy")
	port             = flag.Int("port", 8080, "port of the application")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: entropy-rest [flags]\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	reporter := &reporting.EntropyReporter{
		DefaultBlockSize: *defaultBlockSize,
		HighEntropy:      *high,
		LowEntropy:       *low,
	}
	router := gin.Default()
	router.POST("api/entropy", rest.ReportFileEntropy(reporter))
	if err := router.Run(fmt.Sprintf(":%v", *port)); err != nil {
		log.Fatal(err)
	}
}
