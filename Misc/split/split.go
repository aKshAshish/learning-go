package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func fileErr(err error) {
	if err != nil {
		log.Fatalf("error reading the file %v", err)
	}
}

const (
	CHUNK_SIZE = 64 // 64 Bytes
)

func main() {
	// Check if the file name has been passed or not through system args.
	if len(os.Args) < 2 {
		log.Fatalf("no file name provided")
	}
	// File name received.
	name := os.Args[1]

	// Get the full file path for the file name provided,
	// i.e. to be used to access the file.
	cwd, err := getCwd()
	name = cwd + string(os.PathSeparator) + name

	// Open the file and handle the error if any.
	fp, err := os.Open(name)
	fileErr(err)
	defer fp.Close()

	// Get File size using file stats.
	info, err := fp.Stat()
	fileErr(err)
	size := info.Size()

	partitions := getPartitions(int(size))

	// Initialize buffered reader
	b := make([]byte, CHUNK_SIZE)
	chunks := make([]string, partitions)

	for i := 0; i < partitions; i++ {
		n, err := fp.Read(b)
		if err != nil {
			log.Fatalf("error occured while reading chunk %d; error %v", i, err)
		}
		chunk, err := writeChunk(os.Args[1], b[:n], i+1)
		chunks[i] = chunk
	}
	fmt.Printf("%s", chunks)
}

func writeChunk(name string, b []byte, num int) (string, error) {
	name = strings.Split(name, ".")[0]
	cwd, err := getCwd()
	if err != nil {
		return "", err
	}
	chunkName := fmt.Sprintf("%s%s%s_chunk_%d", cwd, string(os.PathSeparator), name, num)

	err = os.WriteFile(chunkName, b, 0644)
	if err != nil {
		return "", err
	}
	return chunkName, nil
}

func getCwd() (string, error) {
	return filepath.Abs(filepath.Dir("."))
}

// getPartitions: calculates the number of partitions that can be created of size CHUNK_SIZE.
func getPartitions(size int) int {
	partitions := size / CHUNK_SIZE
	if size%CHUNK_SIZE != 0 {
		partitions += 1
	}
	return partitions
}
