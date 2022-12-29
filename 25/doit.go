package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SNAFUtoINT(s string) (int64, error) {
	var res int64
	var base int64
	res = 0
	base = 1
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		switch c {
		case '1':
			res += base
		case '2':
			res += base * 2
		case '-':
			res -= base
		case '=':
			res -= base * 2
		case '0':
		default:
			return 0, fmt.Errorf("Not a SNAFU: %s", s)
		}
		base = base * 5
	}
	return res, nil
}

func INTtoSNAFU(v int64) string {
	base5string := strconv.FormatInt(v, 5)
	base := 1
	passover := byte(0)
	res := []int{}
	for i := len(base5string) - 1; i >= 0; i-- {
		c := base5string[i]
		fv := (c - '0') + passover
		if fv <= 2 {
			res = append(res, int(fv))
			base = base * 5
			passover = 0
			continue
		}
		if fv == 3 {
			res = append(res, -2)
			base = base * 5
			passover = 1
		}
		if fv == 4 {
			res = append(res, -1)
			base = base * 5
			passover = 1
		}
		if fv == 5 {
			res = append(res, 0)
			base = base * 5
			passover = 1
		}
	}
	if passover > 0 {
		res = append(res, int(passover))
	}
	resstring := ""
	for i := len(res) - 1; i >= 0; i-- {
		switch res[i] {
		case 0:
			resstring += "0"
		case 1:
			resstring += "1"
		case 2:
			resstring += "2"
		case -1:
			resstring += "-"
		case -2:
			resstring += "="
		default:
			panic(fmt.Sprintf("Bad element: %d (was converting %d)", res[i], v))
		}
	}
	return resstring
}

func processPart1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic("Load error")
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	var res int64
	for fs.Scan() {
		v, err := SNAFUtoINT(fs.Text())
		if err != nil {
			println(err.Error())
		}
		res += v
	}
	println(INTtoSNAFU(res))
}

func main() {
	processPart1("test.txt")
	processPart1("input.txt")
}
