package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	if arg == "init" {
		for d := range 24 {
			day := d + 1
			err := os.MkdirAll(fmt.Sprintf("day%d", day), 0755)
			if err != nil {
				panic(err)
			}
			fileData := fmt.Sprintf("package day%d\n\nfunc run() {\n\n}\n", day)
			file, err := os.Create(fmt.Sprintf("day%d/solution.go", day))
			if err != nil {
				panic(err)
			}
			file.Write([]byte(fileData))
		}
		return
	}
	day, err := strconv.Atoi(arg)

	if err != nil {
		panic(err)
	}
	fmt.Printf("AoC 2024 - Running day %d\n", day)
}
