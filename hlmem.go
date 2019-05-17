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
	iterPtr := flag.Int("i", 1, "Number of iterations")
	flag.Parse()

	log.Println("Allocating two", *sizePtr, "byte memory blocks")

	memSize, err := strconv.ParseInt(*sizePtr, 10, 64)
	if unit := (*sizePtr)[len(*sizePtr)-1:]; err != nil {
		memSize, err = strconv.ParseInt((*sizePtr)[:len(*sizePtr)-1], 10, 64)
		if err != nil {
			log.Println("Bad memory size", *sizePtr)
			os.Exit(1)
		}
		log.Println("do stuff with", *sizePtr)
		switch unit {
		case "k":
			memSize = memSize * 1024
		case "M":
			memSize = memSize * 1024 * 1024
		case "G":
			memSize = memSize * 1024 * 1024 * 1024
		default:
			log.Println("Bad memory size", *sizePtr)
			os.Exit(1)
		}
	}

	src = make([]byte, memSize)
	dst = make([]byte, memSize)
	log.Println("Inserting random data")
	rand.Read(src)

	log.Println("Copying data, itertions:", *iterPtr)
	start = time.Now()
	for i := 0; i < *iterPtr; i++ {
		copy(dst, src)
	}
	end = time.Now()

	log.Println("Copy completed in", end.Sub(start))
	log.Printf("%.2f MB/sec\n", float64(memSize)*float64(*iterPtr)/float64(end.Sub(start))*1000)
}
