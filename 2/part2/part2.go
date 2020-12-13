package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	column := 0
	validPw := 0
	invalidPw := 0
	pos1 := 0
	pos2 := 0
	letterCount := 0
	var letter rune
	var tmp []string
	for s.Scan() {
		if column == 0 {
			// parse number range
			tmp = strings.Split(s.Text(), "-")
			pos1, err = strconv.Atoi(tmp[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			pos2, err = strconv.Atoi(tmp[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			column++
		} else if column == 1 {
			letter, _ = utf8.DecodeRuneInString(s.Text())
			column++
		} else {
			// check password
			for i, l := range s.Text() {
				if i+1 == pos1 || i+1 == pos2 {
					if l == letter {
						letterCount++
					}
				}
			}
			if letterCount == 1 {
				validPw++
			} else {
				invalidPw++
			}
			pos1 = 0
			pos2 = 0
			letterCount = 0
			column = 0
		}
	}
	println("Invalid passwords:", invalidPw)
	println("Valid passwords:", validPw)
}
