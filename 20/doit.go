package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	value    int
	is_moved bool
}

type Numbers struct {
	initial_order *[]*Number
	N             int
	actual_order  *[]*Number
}

func Print(n *Numbers) {
	println(String(n))
}

func String(n *Numbers) string {
	res := ""
	for _, p := range *n.actual_order {
		res += " "
		res += strconv.Itoa(p.value)
	}
	return res
}

func Load(filename string) *Numbers {
	var res Numbers
	f, err := os.Open(filename)
	if err != nil {
		panic("Load error")
	}
	res.initial_order = &[]*Number{}
	res.actual_order = &[]*Number{}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	for fs.Scan() {
		value, err := strconv.Atoi(fs.Text())
		if err != nil {
			continue
		}
		n := Number{value: value, is_moved: false}
		*(res.initial_order) = append(*(res.initial_order), &n)
		*(res.actual_order) = append(*(res.actual_order), &n)
		res.N++
	}
	return &res
}

func Mix(res *Numbers) {
	for i, _ := range *res.initial_order {
		//		Print(res)
		Move(res, i)
	}
}

func Mult(res *Numbers) {
	for _, v := range *res.initial_order {
		v.value = v.value * 811589153
	}
}
func GetCoords(res *Numbers) string {
	zero_location := -1
	for i, v := range *res.actual_order {
		if v.value == 0 {
			zero_location = i
			break
		}
	}
	v1 := (*(res.actual_order))[(zero_location+1000)%res.N].value
	v2 := (*(res.actual_order))[(zero_location+2000)%res.N].value
	v3 := (*(res.actual_order))[(zero_location+3000)%res.N].value
	return strconv.Itoa(v1 + v2 + v3)
}

func Move(r *Numbers, i int) {
	n := (*r.initial_order)[i]
	new_pos := -1
	found_pos := -1
	for actual_pos, nn := range *r.actual_order {
		if n == nn {
			found_pos = actual_pos
			new_pos = (actual_pos + n.value)
			if new_pos > 0 {
				new_pos = new_pos % (r.N - 1)
			} else {
				new_pos = r.N + new_pos%(r.N-1) - 1
			}
			break
		}
	}
	if new_pos >= 0 {
		ao := (*r.actual_order)[:found_pos]
		ao = append(ao, (*r.actual_order)[found_pos+1:]...)
		bo := []*Number{}
		bo = append(bo, ao[:new_pos]...)
		bo = append(bo, n)
		bo = append(bo, ao[new_pos:]...)
		r.actual_order = &bo
	}
}

func main() {
	//	N := Load("input.txt")
	N := Load("test.txt")
	Mix(N)
	Print(N)
	fmt.Printf(GetCoords(N))
	fmt.Printf("--Part 2---\n")
	N = Load("input.txt")
	Mult(N)
	Mix(N)
	Mix(N)
	Mix(N)
	Mix(N)
	Mix(N)
	Mix(N)
	Mix(N)
	Mix(N)
	Mix(N)
	Mix(N)
	fmt.Printf(GetCoords(N))
}
