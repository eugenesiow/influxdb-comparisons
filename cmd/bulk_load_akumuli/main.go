// bulk_load_opentsdb loads an Akumuli daemon with data from stdin.
//
// The caller is responsible for assuring that the database is empty before
// bulk load.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pkg/profile"
)

// Program option vars:
var (
	csvDaemonUrls string
	daemonUrls    []string
	workers       int
	batchSize     int
	backoff       time.Duration
	doLoad        bool
	memprofile    bool
)

// Global vars
var (
	bufPool      sync.Pool
	batchChan    chan *bytes.Buffer
	inputDone    chan struct{}
	workersGroup sync.WaitGroup
)

// Parse args:
func init() {
	flag.StringVar(&csvDaemonUrls, "urls", "127.0.0.1:8282", "OpenTSDB URLs, comma-separated. Will be used in a round-robin fashion.")
	flag.IntVar(&batchSize, "batch-size", 5000, "Batch size (input lines).")
	flag.IntVar(&workers, "workers", 1, "Number of parallel requests to make.")
	//flag.DurationVar(&backoff, "backoff", time.Second, "Time to sleep between requests when server indicates backpressure is needed.")
	flag.BoolVar(&doLoad, "do-load", true, "Whether to write data. Set this flag to false to check input read speed.")
	flag.BoolVar(&memprofile, "memprofile", false, "Whether to write a memprofile (file automatically determined).")

	flag.Parse()

	daemonUrls = strings.Split(csvDaemonUrls, ",")
	if len(daemonUrls) == 0 {
		log.Fatal("missing 'urls' flag")
	}
	fmt.Printf("daemon URLs: %v\n", daemonUrls)
}

func main() {
	if memprofile {
		p := profile.Start(profile.MemProfile)
		defer p.Stop()
	}

	bufPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 4*1024*1024))
		},
	}

	batchChan = make(chan *bytes.Buffer, workers)
	inputDone = make(chan struct{})

	for i := 0; i < workers; i++ {
		daemonUrl := daemonUrls[i%len(daemonUrls)]
		workersGroup.Add(1)
		writer, err := NewTCPWriter(daemonUrl)
		if err != nil {
			log.Fatal(err)
		}
		go processBatches(writer)
	}

	start := time.Now()
	itemsRead := scan(batchSize)

	<-inputDone
	close(batchChan)

	workersGroup.Wait()

	end := time.Now()
	took := end.Sub(start)
	rate := float64(itemsRead) / float64(took.Seconds())

	fmt.Printf("loaded %d items in %fsec with %d workers (mean rate %f/sec)\n", itemsRead, took.Seconds(), workers, rate)
}

// scan reads one line at a time from stdin.
// When the requested number of lines per batch is met, send a batch over batchChan for the workers to write.
func scan(linesPerBatch int) int64 {

	var itemsRead int64

	newline := []byte("\r\n")

	scanner := bufio.NewScanner(bufio.NewReaderSize(os.Stdin, 4*1024*1024))
	for scanner.Scan() {
		itemsRead++

		batch := bufPool.Get().(*bytes.Buffer)
		batch.Reset()

		batch.Write(scanner.Bytes())
		batch.Write(newline)

		batchChan <- batch

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %s", err.Error())
	}

	// Closing inputDone signals to the application that we've read everything and can now shut down.
	close(inputDone)

	return itemsRead
}

// processBatches reads byte buffers from batchChan and writes them to the target server, while tracking stats on the write.
func processBatches(w LineProtocolWriter) {
	for batch := range batchChan {
		// Write the batch: try until backoff is not needed.
		if doLoad {
			var err error
			_, err = w.WriteLineProtocol(batch.Bytes())
			if err != nil {
				log.Fatalf("Error writing: %s\n", err.Error())
			}
		}

		// Return the batch buffer to the pool.
		bufPool.Put(batch)
	}
	workersGroup.Done()
}
