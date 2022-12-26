package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Monkey struct {
	value    int
	name     string
	known    bool
	goesdown bool
	operator string
	op1      string
	op2      string
	deps     map[string]*Monkey
	group    *Monkeys
}

type Monkeys struct {
	all map[string]*Monkey
}

func Parse(line string) (*Monkey, error) {
	m := &Monkey{}
	m.deps = make(map[string]*Monkey)
	if len(line) < 4 {
		return nil, fmt.Errorf("short line %s", line)
	}

	if line[4] != ':' {
		return nil, fmt.Errorf("Bad fmt: %s", line)
	}
	m.name = line[:4]
	v, err := strconv.Atoi(line[6:])
	if err != nil {
		if len(line) < 17 {
			return nil, fmt.Errorf("Bad fmt: %s", line)
		}
		m.operator = string(line[11])
		m.op1 = line[6:10]
		m.op2 = line[13:17]
		return m, nil
	}
	m.known = true
	m.operator = ":"
	m.value = v
	return m, nil
}

func (ms *Monkeys) FillDeps() error {
	for _, v := range ms.all {
		if !v.known {
			m1, ok1 := ms.all[v.op1]
			m2, ok2 := ms.all[v.op2]
			if !ok1 || !ok2 {
				return fmt.Errorf("Ups: %s ", v.name)
			}
			m1.deps[v.name] = v
			m2.deps[v.name] = v
		}
	}
	return nil
}

func (m *Monkey) Get2() (int, error) {
	//
	// Part2 assumptions:
	//     o
	//    /
	//   o
	//  / \
	// o   o
	// strict tree - no reuse of nodes
	if m.operator == ":" {
		if m.name == "humn" {
			for _, v := range m.deps {
				m1, ok := m.group.all[v.name]
				if ok {
					m1.goesdown = true
					return m1.Get2()
				}
			}
		}
		return m.value, nil
	}
	m1, ok1 := m.group.all[m.op1]
	if !ok1 {
		return 0, fmt.Errorf("%s unknown", m.op1)
	}
	m2, ok2 := m.group.all[m.op2]
	if !ok2 {
		return 0, fmt.Errorf("%s unknown", m.op2)
	}

	if m.name == "root" {
		if m1.goesdown {
			// m2 goes up
			v, err := m2.Get()
			if err != nil {
				return 0, fmt.Errorf("Get2 root problem")
			}
			m.value = v
			m1.value = v
			m1.known = true
			return v, nil
		}
		if m2.goesdown {
			v, err := m1.Get()
			if err != nil {
				return 0, fmt.Errorf("Get2 - 1 root problem")
			}
			m.value = v
			m2.value = v
			m2.known = true
			return v, nil
		}
	}
	if m.goesdown {
		if m.known {
			return m.value, nil
		}
		for _, v := range m.deps {
			mup, ok := m.group.all[v.name]
			if ok {
				mup.goesdown = true
				required_result, err := mup.Get2()
				if err != nil {
					return 0, fmt.Errorf("Problem at %s: %s", m.name, err.Error())
				}
				if m1.goesdown {
					v2, err := m2.Get()
					if err != nil {
						return 0, fmt.Errorf("Problem at %s v2=%s: %s", m.name, m2.name, err.Error())
					}
					m.known = true
					m.value = required_result
					var v1 int
					if m.operator == "+" {
						v1 = required_result - v2
					}
					if m.operator == "-" {
						v1 = required_result + v2
					}
					if m.operator == "*" {
						v1 = required_result / v2
					}
					if m.operator == "/" {
						v1 = required_result * v2
					}
					m1.known = true
					m1.value = v1
					return v1, nil
				}
				if m2.goesdown {
					v1, err := m1.Get()
					if err != nil {
						return 0, fmt.Errorf("Problem at %s v1=%s: %s", m.name, m1.name, err.Error())
					}
					m.known = true
					m.value = required_result
					var v2 int
					if m.operator == "+" {
						v2 = required_result - v1
					}
					if m.operator == "-" {
						v2 = v1 - required_result
					}
					if m.operator == "*" {
						v2 = required_result / v1
					}
					if m.operator == "/" {
						v2 = v1 / required_result
					}

					m2.known = true
					m2.value = v2
					return v2, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("Wrong call on %s", m.name)
}

func (m *Monkey) Get() (int, error) {
	if m.known {
		return m.value, nil
	}
	m1, ok1 := m.group.all[m.op1]
	if !ok1 {
		return 0, fmt.Errorf("%s unknown", m.op1)
	}
	m2, ok2 := m.group.all[m.op2]
	if !ok2 {
		return 0, fmt.Errorf("%s unknown", m.op2)
	}
	v1, e1 := m1.Get()
	if e1 != nil {
		return 0, e1
	}
	v2, e2 := m2.Get()
	if e2 != nil {
		return 0, e2
	}
	if m.operator == "+" {
		m.value = v1 + v2
		m.known = true
		return m.value, nil
	}

	if m.operator == "-" {
		m.value = v1 - v2
		m.known = true
		return m.value, nil
	}
	if m.operator == "*" {
		m.value = v1 * v2
		m.known = true
		return m.value, nil
	}
	if m.operator == "/" {
		if v2 != 0 {
			m.value = v1 / v2
			m.known = true
			return m.value, nil
		}
		panic(fmt.Sprintf("Division by 0 came from %s", m2.name))
	}
	panic(fmt.Sprintf("Unknown operator: %s", m.operator))
}

func Load(filename string) Monkeys {
	var res Monkeys
	res.all = make(map[string]*Monkey)
	f, err := os.Open(filename)
	if err != nil {
		panic("Load error")
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	for fs.Scan() {
		m, err := Parse(fs.Text())
		if err == nil {
			m.group = &res
			res.all[m.name] = m
		}
	}
	return res
}

func processPart1(filename string) {
	fmt.Printf("Part1 %s\n", filename)
	ms := Load(filename)
	m, ok := ms.all["root"]
	if ok {
		v, err := m.Get()
		if err == nil {
			fmt.Printf("root: %d\n", v)
		} else {
			fmt.Print(err)
		}
	}
}

func processPart2(filename string) {
	fmt.Printf("Part2 %s\n", filename)
	ms := Load(filename)
	err := ms.FillDeps()
	if err != nil {
		fmt.Printf("Error filling deps: %s\n", err.Error())
		return
	}
	m, ok := ms.all["humn"]
	if !ok {
		fmt.Printf("No humn")
		return
	}
	m.goesdown = true
	m.known = false
	v, err := m.Get2()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Part2: %d\n", v)
}

func main() {
	processPart1("test.txt")
	processPart1("input.txt")
	processPart2("test.txt")
	processPart2("input.txt")
}
