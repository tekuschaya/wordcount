package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str := readInput()
	fmt.Println(len(str))
}

func readInput() []string {
	var str []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		str = append(str, word)
	}
	return str
}