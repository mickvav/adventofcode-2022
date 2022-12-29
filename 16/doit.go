package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Map struct {
	valves       map[string]int
	valvenames   []string
	valvenumbers map[string]int
	connections  map[string]map[string]bool
}

func Init(filename string) *Map {
	res := Map{valves: make(map[string]int), connections: make(map[string]map[string]bool), valvenames: make([]string, 0), valvenumbers: make(map[string]int)}

	f, err := os.Open(filename)
	if err != nil {
		panic("Load error")
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	re := regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=([0-9]+); tunnels? leads? to valves? (.*)`)
	sep := regexp.MustCompile(`, `)
	valvenumber := 0
	for fs.Scan() {
		m := re.FindSubmatch(fs.Bytes())
		if m != nil {
			vn := string(m[1])
			res.valvenames = append(res.valvenames, vn)
			res.valvenumbers[vn] = valvenumber
			rate, err := strconv.Atoi(string(m[2]))
			if err != nil {
				println("Atoi error:", err.Error())
				continue
			}
			res.valves[string(m[1])] = rate
			peers := sep.Split(string(m[3]), -1)
			for _, p := range peers {
				res.Connect(string(m[1]), p)
			}
			valvenumber++
		} else {
			println("Unable to parse: ", string(fs.Bytes()))
		}
	}
	return &res
}

func (m *Map) Connect(v1 string, v2 string) {
	p, ok := m.connections[v1]
	if ok {
		p[v2] = true
	} else {
		p = map[string]bool{v2: true}
		m.connections[v1] = p
	}
	p, ok = m.connections[v2]
	if ok {
		p[v1] = true
	} else {
		p = map[string]bool{v1: true}
		m.connections[v2] = p
	}
}

type State struct {
	valvestatus uint64 // We have only 59 valves
	actorpos    string
}

func (s *State) IsOpen(v string, m *Map) bool {
	vn := m.valvenumbers[v]
	var b uint64
	b = 1 << uint64(vn)
	return (s.valvestatus & b) > 0
}

func (s *State) Open(v string, m *Map) {
	vn := m.valvenumbers[v]
	var b uint64
	b = 1 << uint64(vn)
	s.valvestatus = s.valvestatus | b
}

func (s *State) Income(m *Map) int {
	var b uint64
	res := 0
	for i, v := range m.valvenames {
		b = 1 << uint64(i)
		if (s.valvestatus & b) > 0 {
			res += m.valves[v]
		}
	}
	return res
}

type Generation struct {
	states     map[State]int
	statetrace map[State]string
	time       int
}

func (g *Generation) ReachedEnd(m *Map) bool {
	return g.time >= 30
}

func GetFirstGeneration(m *Map) *Generation {
	res := Generation{states: make(map[State]int), statetrace: make(map[State]string), time: 0}
	s := State{valvestatus: 0, actorpos: "AA"}
	res.states[s] = 0
	res.statetrace[s] = ""
	return &res
}

func (g *Generation) GetNextGeneration(m *Map) *Generation {
	res := Generation{states: make(map[State]int), statetrace: make(map[State]string), time: g.time + 1}
	for s, v := range g.states {
		// 1. Move to neighbours
		for p, _ := range m.connections[s.actorpos] {
			vnew := v + s.Income(m)
			s1 := State{valvestatus: s.valvestatus, actorpos: p}
			existingvnew, ok := res.states[s1]
			if (ok && existingvnew < vnew) || (!ok) {
				res.states[s1] = vnew
				res.statetrace[s1] = g.statetrace[s] + p
			}
		}
		// 2. Open valve
		if !s.IsOpen(s.actorpos, m) && m.valves[s.actorpos] > 0 {
			s1 := State{valvestatus: s.valvestatus, actorpos: s.actorpos}
			vnew := v + s1.Income(m)
			s1.Open(s.actorpos, m)
			existingvnew, ok := res.states[s1]
			if (ok && existingvnew < vnew) || (!ok) {
				res.states[s1] = vnew
				res.statetrace[s1] = g.statetrace[s] + "[]"
			}
		}
	}
	return &res
}

func (g *Generation) GetValue() (int, string) {
	vmax := 0
	sp := State{actorpos: ""}
	for s, v := range g.states {
		if v > vmax {
			vmax = v
			sp.actorpos = s.actorpos
			sp.valvestatus = s.valvestatus
		}
	}
	if sp.actorpos == "" {
		return vmax, ""
	}
	return vmax, g.statetrace[sp]
}

func processPart1(filename string) {
	m := Init(filename)
	g := GetFirstGeneration(m)
	for !g.ReachedEnd(m) {
		g = g.GetNextGeneration(m)
		v, trace := g.GetValue()
		println(g.time, v, trace, len(g.states))
	}
	println(g.GetValue())
}

func main() {
	processPart1("test.txt")
	processPart1("input.txt")
	// processPart2("test.txt")
	// processPart2("input.txt")
}
