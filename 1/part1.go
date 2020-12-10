package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	var result []int
	for s.Scan() {
		x, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println(x, err)
			return
		}
		result = append(result, x)
	}
	for _, n := range result {
		for _, m := range result {
			if n+m == 2020 {
				fmt.Println(n * m)
				return
			}
		}
	}

}
