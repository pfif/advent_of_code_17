package main

import (
	"fmt"
	"os"

	"florent/adventofcode/2022/day7/explorer"
	"florent/adventofcode/2022/day7/parser"
)

func main() {
	directory, err := parser.ParseDirectories()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	println(explorer.SumSmallDirectoriesSizes(AdapterDirectories{source: directory}))
}

type AdapterDirectories struct {
	source parser.Directory
}

func (d AdapterDirectories) FileSizes() []int {
	sizes := []int{}
	for _, file := range d.source.Files() {
		sizes = append(sizes, file.Size())
	}
	return sizes
}

func (d AdapterDirectories) SubDirectories() []explorer.Directory {
	result := []explorer.Directory{}
	for _, subdir := range d.source.Subdirectories() {
		result = append(result, &AdapterDirectories{source: subdir})
	}

	return result
}
