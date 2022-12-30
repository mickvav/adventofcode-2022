package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"time"
)

type Blueprint struct {
	oreRobotCostOre        int
	clayRobotCostOre       int
	obsidianRobotCostOre   int
	obsidianRobotCostClay  int
	geodeRobotCostObsidian int
	geodeRobotCostOre      int
}

type Map struct {
	blueprints map[int]Blueprint
}

func Init(filename string) *Map {
	res := Map{blueprints: make(map[int]Blueprint)}

	f, err := os.Open(filename)
	if err != nil {
		panic("Load error")
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	re := regexp.MustCompile(`Blueprint ([0-9]+): Each ore robot costs ([0-9]+) ore. Each clay robot costs ([0-9]+) ore. Each obsidian robot costs ([0-9]+) ore and ([0-9]+) clay. Each geode robot costs ([0-9]+) ore and ([0-9]+) obsidian.`)
	for fs.Scan() {
		m := re.FindStringSubmatch(string(fs.Bytes()))
		if m != nil {
			b := Blueprint{}
			b.oreRobotCostOre, err = strconv.Atoi(m[2])
			if err != nil {
				panic("Cant parse int")
			}
			b.clayRobotCostOre, err = strconv.Atoi(m[3])
			if err != nil {
				panic("Cant parse int")
			}
			b.obsidianRobotCostOre, err = strconv.Atoi(m[4])
			if err != nil {
				panic("Cant parse int")
			}
			b.obsidianRobotCostClay, err = strconv.Atoi(m[5])
			if err != nil {
				panic("Cant parse int")
			}
			b.geodeRobotCostOre, err = strconv.Atoi(m[6])
			if err != nil {
				panic("Cant parse int")
			}
			b.geodeRobotCostObsidian, err = strconv.Atoi(m[7])
			if err != nil {
				panic("Cant parse int")
			}
			id, err := strconv.Atoi(m[1])
			if err != nil {
				panic("Cant parse int")
			}
			res.blueprints[id] = b
		} else {
			println("Unable to parse: ", string(fs.Bytes()))
		}
	}
	return &res
}

type State struct {
	ore            int
	clay           int
	obsidian       int
	geodes         int
	orerobots      int
	clayrobots     int
	obsidianrobots int
	geoderobots    int
}

type Generation struct {
	states  map[State]bool
	time    int
	deleted int
}

func (g *Generation) ReachedEnd(m *Map) bool {
	return g.time >= 24
}

func (g *Generation) ReachedEnd2(m *Map) bool {
	return g.time >= 32
}

func GetFirstGeneration() *Generation {
	res := Generation{states: make(map[State]bool), time: 0}
	s := State{orerobots: 1}
	res.states[s] = true
	return &res
}

func (s *State) CanCreateOreRobot(bp Blueprint) bool {
	return s.ore >= bp.oreRobotCostOre
}

func (s *State) CanCreateClayRobot(bp Blueprint) bool {
	return s.ore >= bp.clayRobotCostOre
}

func (s *State) CanCreateObsidianRobot(bp Blueprint) bool {
	return s.ore >= bp.obsidianRobotCostOre && s.clay >= bp.obsidianRobotCostClay
}
func (s *State) CanCreateGeodesRobot(bp Blueprint) bool {
	return s.ore >= bp.geodeRobotCostOre && s.obsidian >= bp.geodeRobotCostObsidian
}

func (s *State) Produce() State {
	return State{
		ore:            s.ore + s.orerobots,
		clay:           s.clay + s.clayrobots,
		obsidian:       s.obsidian + s.obsidianrobots,
		geodes:         s.geodes + s.geoderobots,
		orerobots:      s.orerobots,
		clayrobots:     s.clayrobots,
		obsidianrobots: s.obsidianrobots,
		geoderobots:    s.geoderobots,
	}
}

func (s *State) CreateGeodesRobot(bp Blueprint) State {
	s1 := s.Produce()
	s1.obsidian -= bp.geodeRobotCostObsidian
	s1.ore -= bp.geodeRobotCostOre
	s1.geoderobots += 1
	return s1
}

func (s *State) CreateObsidianRobot(bp Blueprint) State {
	s1 := s.Produce()
	s1.clay -= bp.obsidianRobotCostClay
	s1.ore -= bp.obsidianRobotCostOre
	s1.obsidianrobots += 1
	return s1
}

func (s *State) CreateClayRobot(bp Blueprint) State {
	s1 := s.Produce()
	s1.ore -= bp.clayRobotCostOre
	s1.clayrobots += 1
	return s1
}

func (s *State) CreateOreRobot(bp Blueprint) State {
	s1 := s.Produce()
	s1.ore -= bp.oreRobotCostOre
	s1.orerobots += 1
	return s1
}

func (g *Generation) GetPessimisticGeodesNumber() int {
	remainingtime := 32 - g.time
	res := 0
	for s := range g.states {
		g := s.geodes + s.geoderobots*remainingtime
		if g > res {
			res = g
		}
	}
	return res
}

func (s *State) GetOptimisticGeodeNumber(g *Generation) int {
	remainingtime := 32 - g.time
	return s.geodes + s.geoderobots*remainingtime + remainingtime*(remainingtime-1)/2
}

func (g *Generation) FindDeleteCandidates(limit int64) (*map[State]bool, bool) {
	start := time.Now()
	candidates := map[State]bool{}

	needsatleast := g.GetPessimisticGeodesNumber()
	for s := range g.states {
		if s.GetOptimisticGeodeNumber(g) < needsatleast {
			candidates[s] = true
		}
	}
	print("Poor guys: ", len(candidates))
	for s, _ := range g.states {
		if _, ok := candidates[s]; ok {
			continue
		}
		for s1, _ := range g.states {
			if _, ok := candidates[s]; ok {
				continue
			}
			if s != s1 {
				if s.ore >= s1.ore && s.clay >= s1.clay && s.obsidian >= s1.obsidian && s.geodes >= s1.geodes &&
					s.orerobots >= s1.orerobots && s.clayrobots >= s1.clayrobots && s.obsidianrobots >= s1.obsidianrobots && s.geoderobots >= s1.geoderobots {
					candidates[s1] = true
				}
				if s1.ore >= s.ore && s1.clay >= s.clay && s1.obsidian >= s.obsidian && s1.geodes >= s.geodes &&
					s1.orerobots >= s.orerobots && s1.clayrobots >= s.clayrobots && s1.obsidianrobots >= s.obsidianrobots && s1.geoderobots >= s.geoderobots {
					candidates[s] = true
					break
				}
			}
		}
		if len(candidates)*3 > len(g.states) {
			break
		}
		if time.Since(start).Microseconds() > limit {
			break
		}
	}
	dt := time.Since(start).Microseconds()
	println("[DC]", dt, float32(dt)/float32((len(g.states)*len(g.states))))
	return &candidates, len(candidates) > 0
}

func (g *Generation) FindDeleteCandidate() (State, bool) {

	for s, _ := range g.states {
		for s1, _ := range g.states {
			if s != s1 {
				if s.ore >= s1.ore && s.clay >= s1.clay && s.obsidian >= s1.obsidian && s.geodes >= s1.geodes &&
					s.orerobots >= s1.orerobots && s.clayrobots >= s1.clayrobots && s.obsidianrobots >= s1.obsidianrobots && s.geoderobots >= s1.geoderobots {
					return s1, true
				}
				if s1.ore >= s.ore && s1.clay >= s.clay && s1.obsidian >= s.obsidian && s1.geodes >= s.geodes &&
					s1.orerobots >= s.orerobots && s1.clayrobots >= s.clayrobots && s1.obsidianrobots >= s.obsidianrobots && s1.geoderobots >= s.geoderobots {
					return s, true
				}
			}
		}
	}
	return State{}, false
}

func (g *Generation) GetNextGeneration(bp Blueprint) *Generation {
	start := time.Now()
	res := Generation{states: make(map[State]bool), time: g.time + 1}
	for s, _ := range g.states {
		if s.CanCreateOreRobot(bp) {
			res.states[s.CreateOreRobot(bp)] = true
		}
		if s.CanCreateClayRobot(bp) {
			res.states[s.CreateClayRobot(bp)] = true
		}
		if s.CanCreateObsidianRobot(bp) {
			res.states[s.CreateObsidianRobot(bp)] = true
		}
		if s.CanCreateGeodesRobot(bp) {
			res.states[s.CreateGeodesRobot(bp)] = true
		}
		res.states[s.Produce()] = true
	}
	deleted := 0
	//	if len(res.states) < 100000 {
	//		found := true
	//		for found {
	//			var s State
	//			s, found = res.FindDeleteCandidate()
	//			delete(res.states, s)
	//			deleted++
	//		}
	//	}
	candidates, _ := res.FindDeleteCandidates(10 + time.Since(start).Microseconds()/2)
	for s := range *candidates {
		delete(res.states, s)
		deleted++
	}
	res.deleted = deleted
	PrintMemUsage()
	runtime.GC()
	return &res
}

func (g *Generation) GetValue() int {
	vmax := 0
	for s, _ := range g.states {
		if s.geodes > vmax {
			vmax = s.geodes
		}
	}
	return vmax
}

func processPart1(filename string) {
	m := Init(filename)
	totalQuality := 0
	for i, b := range m.blueprints {
		g := GetFirstGeneration()
		for !g.ReachedEnd(m) {
			g = g.GetNextGeneration(b)
			println(filename, " ", i, " ", g.time, " ", len(g.states), " ", g.deleted, " ", g.GetValue())
		}
		v := g.GetValue()
		totalQuality += v * i
	}
	print(filename, " total quality: ", totalQuality)
}

func processPart2(filename string) {
	m := Init(filename)
	total := 1
	starttime := time.Now()
	for i, b := range m.blueprints {
		g := GetFirstGeneration()
		for !g.ReachedEnd2(m) {
			g = g.GetNextGeneration(b)
			println(filename, " ", i, " ", g.time, " ", len(g.states), " ", g.deleted, " ", g.GetValue())
			println(time.Since(starttime).Milliseconds())
		}
		v := g.GetValue()
		total *= v
	}
	print(filename, " total: ", total)
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	//	processPart1("test.txt")
	//	processPart1("input.txt")
	processPart2("test2.txt")
	processPart2("input2.txt")
}
