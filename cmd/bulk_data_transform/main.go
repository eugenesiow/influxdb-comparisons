// bulk_data_gen generates time series data from pre-specified use cases.
//
// Supported formats:
// InfluxDB bulk load format
// ElasticSearch bulk load format
// Cassandra query format
// Mongo custom format
// OpenTSDB bulk HTTP format
//
// Supported use cases:
// Devops: scale_var is the number of hosts to simulate, with log messages
//         every 10 seconds.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"path/filepath"
)

// Output data format choices:
var formatChoices = []string{"akumuli", "influx-bulk", "es-bulk", "cassandra", "mongo", "opentsdb"}

// Use case choices:
var useCaseChoices = []string{"shelburne","green_taxi","srbench"}

// Program option vars:
var (
	daemonUrl string
	dbName    string

	format  string
	useCase string

	inputFile string

	interleavedGenerationGroupID uint
	interleavedGenerationGroups  uint

	debug int
)

// Parse args:
func init() {
	flag.StringVar(&format, "format", formatChoices[0], fmt.Sprintf("Format to emit. (choices: %s)", strings.Join(formatChoices, ", ")))

	flag.StringVar(&useCase, "use-case", useCaseChoices[0], "Use case to model. (choices: shelburne, green_taxi, srbench)")

	flag.StringVar(&inputFile, "input", "data.csv", "Input file.")

	flag.IntVar(&debug, "debug", 0, "Debug printing (choices: 0, 1, 2) (default 0).")

	flag.UintVar(&interleavedGenerationGroupID, "interleaved-generation-group-id", 0, "Group (0-indexed) to perform round-robin serialization within. Use this to scale up data generation to multiple processes.")
	flag.UintVar(&interleavedGenerationGroups, "interleaved-generation-groups", 1, "The number of round-robin serialization groups. Use this to scale up data generation to multiple processes.")

	flag.Parse()

	if !(interleavedGenerationGroupID < interleavedGenerationGroups) {
		log.Fatal("incorrect interleaved groups configuration")
	}

	validFormat := false
	for _, s := range formatChoices {
		if s == format {
			validFormat = true
			break
		}
	}
	if !validFormat {
		log.Fatal("invalid format specifier")
	}

}

func main() {
	switch useCase {
	case "srbench":
		filepath.Walk(inputFile, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				//println(info.Name())
				runGenerator(path)
			}
			return nil
		})
	default:
		runGenerator(inputFile)
	}
}

func runGenerator(inputPath string) {
	out := bufio.NewWriterSize(os.Stdout, 4<<20)
	defer out.Flush()

	var sim Simulator

	cfg := &DevopsSimulatorConfig{
		filePath: inputPath,
		useCase: useCase,
		format: format,
	}
	sim = cfg.ToSimulator()

	var serializer func(*Point, io.Writer) error
	switch format {
	case "akumuli":
		serializer = (*Point).SerializeAkumuliBulk
	case "influx-bulk":
		serializer = (*Point).SerializeInfluxBulk
	case "es-bulk":
		serializer = (*Point).SerializeESBulk
	case "cassandra":
		serializer = (*Point).SerializeCassandra
	case "mongo":
		serializer = (*Point).SerializeMongo
	case "opentsdb":
		serializer = (*Point).SerializeOpenTSDBBulk
	default:
		panic("unreachable")
	}

	var currentInterleavedGroup uint = 0

	point := MakeUsablePoint()
	for !sim.Finished() {
		sim.Next(point)

		// in the default case this is always true
		if currentInterleavedGroup == interleavedGenerationGroupID {
			//println("printing")
			if point.Timestamp==nil {
				break
			}
			err := serializer(point, out)
			if err != nil {
				log.Fatal(err)
			}

		}
		point.Reset()

		currentInterleavedGroup++
		if currentInterleavedGroup == interleavedGenerationGroups {
			currentInterleavedGroup = 0
		}
	}

	err := out.Flush()
	if err != nil {
		log.Fatal(err.Error())
	}
}
