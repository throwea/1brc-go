package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// TODO:
// - Read in the file line by line -> DONE
// - Cities appear multiple times
// - Implement the brute force solution first then track the time it takes

func main() {
	start := time.Now()
	fmt.Println("Hello World")
	parseFile()
	fmt.Printf("Time taken: %2f", time.Since(start).Seconds())
}

func parseFile() {
	// First we need to read the file into an object
	readFile, err := os.Open("../measurements.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// lines := 1000
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		// lines -= 1
		// if lines <= 0 {
		// 	break
		// }
	}
	readFile.Close()
}
