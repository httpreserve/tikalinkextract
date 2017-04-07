package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	noprotocol   bool
	file         string
	vers         bool
	fileThrottle = 8
	totalFiles   int
	quoteCells   = false
)

func init() {
	flag.StringVar(&file, "file", "false", "File to extract information from.")
	flag.BoolVar(&vers, "version", false, "[Optional] Output version of the tool.")
	flag.BoolVar(&noprotocol, "noprotocol", false, "[Optional] For www. links (without a protocol, don't prepend http://.")
	flag.BoolVar(&quoteCells, "quote", false, "[Optional] Some URLS may contain commas, quote cells for your CSV parser.")
}

func outputList(linkpool []string) {
	defer wg.Done()
	if len(linkpool) > 0 {
		for _, x := range linkpool {
			fmt.Fprintf(os.Stdout, "%s\n", x)
		}
		linkpool = linkpool[:0]
	}
}

var wg sync.WaitGroup

func processall(file string) {
	//check the services are up and running
	findOpenConnections()

	//make a listing of all the files we're going to process
	//efficient enough with memory?
	err := filepath.Walk(file, readFile)
	if err != nil {
		logStringError("%v", err)
		os.Exit(1)
	}

	//time how long it takes to prcess files and extract entities
	start := time.Now()

	//read each file into each server and collect results

	if len(allfiles) <= 0 {
		fmt.Fprintf(os.Stderr, "No files to process.\n")
		os.Exit(1)
	}

	totalFiles = len(allfiles)

	for x := 0; x < len(allfiles); x += fileThrottle {
		remain := min(x+fileThrottle, len(allfiles))
		filepool := allfiles[x:remain]
		b := true
		for b {
			b, err = extractAndAnalyse(filepool)
			if err != nil {
				logStringError("%v", err)
				os.Exit(1)
			}
		}
		linkpool := make([]string, len(linklist))
		copy(linkpool, linklist)
		linklist = linklist[:0]

		// process output in background while we handle other filed
		wg.Add(1)
		go outputList(linkpool)

		//release waitgroup, exit...i believe this will prevent race
		//conditions when working between the two lists in this loop.
		wg.Wait()
	}

	//output that time...
	elapsed := time.Since(start)
	fmt.Fprintf(os.Stderr, "\nTika extract took %s\n", elapsed)

}

func main() {
	flag.Parse()
	if flag.NFlag() <= 0 { // can access args w/ len(os.Args[1:]) too
		fmt.Fprintln(os.Stderr, "Usage:  links [-file ...]")
		fmt.Fprintln(os.Stderr, "Usage:        [Optional -noprotocol]")
		fmt.Fprintln(os.Stderr, "Usage:        [Optional -quote]")
		fmt.Fprintln(os.Stderr, "              [Optional -version]")
		fmt.Fprintln(os.Stderr, "Output: [CSV] {filename}, {link}")
		fmt.Fprintln(os.Stderr, "Output: [CSV] \"{filename}\", \"{link}\"")
		flag.Usage()
		os.Exit(0)
	}
	if vers {
		fmt.Fprintln(os.Stdout, getVersion())
		os.Exit(1)
	}

	processall(file)
}

//math.Min uses float64, so let's not cast
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
