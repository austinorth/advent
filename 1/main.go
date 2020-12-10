package main

func main() {
	f, err := os.Open("input.txt")
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	var result = []int
	for s.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}

}
