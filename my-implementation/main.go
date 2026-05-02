package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")
	parseFile()
}

func parseFile() {
	// First we need to read the file into an object
	readFile, err := os.Open("../measurements.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lines := 1000
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		lines -= 1
		if lines <= 0 {
			break
		}
	}
	readFile.Close()
}
