package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var src, dst []byte
	var start, end time.Time

	sizePtr := flag.String("s", "64M", "Size of memory to copy in bytes")
	iterPtr := flag.String("i", "1k", "Number of iterations")
	flag.Parse()

	log.Println("Allocating two", *sizePtr, "byte memory blocks")

	memSize := strUnitToInt64(*sizePtr)
	iterations := int( strUnitToInt64(*iterPtr) )

	src = make([]byte, memSize)
	dst = make([]byte, memSize)
	log.Println("Inserting random data")
	rand.Read(src)

	log.Println("Copying data, iterations:", *iterPtr)
	start = time.Now()
	for i := 0; i < iterations; i++ {
		copy(dst, src)
	}
	end = time.Now()

	log.Println("Copy completed in", end.Sub(start))
	log.Printf("%.2f MB/sec\n", float64(memSize)*float64(iterations)/float64(end.Sub(start))*1000)
}

func strUnitToInt64(input string) int64 {
	value, err := strconv.ParseInt(input, 10, 64)
	if unit := input[len(input)-1:]; err != nil {
		value, err = strconv.ParseInt(input[:len(input)-1], 10, 64)
		if err != nil {
			log.Println("Bad input", input)
			os.Exit(1)
		}

		switch unit {
		case "k":
			value = value * 1024
		case "M":
			value = value * 1024 * 1024
		case "G":
			value = value * 1024 * 1024 * 1024
		default:
			log.Println("Bad value", input)
			os.Exit(1)
		}
	}
	return value
}
