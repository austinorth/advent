package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	column := 0
	validPw := 0
	invalidPw := 0
	min := 0
	max := 0
	letterCount := 0
	var tmp []string
	var letter string
	for s.Scan() {
		if column == 0 {
			// parse number range
			tmp = strings.Split(s.Text(), "-")
			min, err = strconv.Atoi(tmp[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			max, err = strconv.Atoi(tmp[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			column++
		} else if column == 1 {
			letter = string((s.Text())
			letter = letter[0]
			column++
		} else {
			// check password
			for _, l := range s.Text() {
				if l == letter {
					letterCount++
				}
			}
			if letter < min || letter > max {
				invalidPw++
			} else {
				validPw++
			}
			min = 0
			max = 0
			letterCount = 0
			column = 0
		}
	}
}
