package main

import (
	"bufio"
	"os"
)

type Coords struct {
	row    int
	column int
}

type Blister struct {
	coords    Coords
	direction rune
}

type Map struct {
	blisters []*Blister
	freedots []*map[Coords]bool
	columns  int
	rows     int
}

func Init(filename string) *Map {
	res := Map{blisters: make([]*Blister, 0), freedots: make([]*map[Coords]bool, 0)}
	f, err := os.Open(filename)
	if err != nil {
		panic("Load error")
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	row := 0
	for fs.Scan() {
		if row == 0 {
			res.columns = len(fs.Text()) - 2
			row++
			continue
		}
		for i, v := range fs.Text() {
			if v == '>' || v == '<' || v == 'v' || v == '^' {
				res.blisters = append(res.blisters, &Blister{coords: Coords{row: row - 1, column: i - 1}, direction: v})
			}
		}
		row++
	}
	res.rows = row - 2
	return &res
}

func (m *Map) ComputeFreeDots(t int) {
	busy := make([][]bool, m.rows)
	for i := range busy {
		busy[i] = make([]bool, m.columns)
	}
	for _, b := range m.blisters {
		var col int
		var row int
		switch b.direction {
		case '>':
			col = (b.coords.column + t) % m.columns
			row = b.coords.row
		case '<':
			col = (b.coords.column - t) % m.columns
			if col < 0 {
				col += m.columns
			}
			row = b.coords.row
		case 'v':
			col = b.coords.column
			row = (b.coords.row + t) % m.rows
		case '^':
			col = b.coords.column
			row = (b.coords.row - t) % m.rows
			if row < 0 {
				row += m.rows
			}
		}
		busy[row][col] = true
	}
	fd := map[Coords]bool{}
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.columns; col++ {
			if !busy[row][col] {
				c := Coords{row: row, column: col}
				fd[c] = true
			}
		}
	}
	m.freedots = append(m.freedots, &fd)
}

func (m *Map) IsFree(t int, row int, col int) bool {
	if (row == -1 && col == 0) || (row == m.rows && col == m.columns-1) {
		return true
	}
	if row < 0 || row > m.rows-1 || col < 0 || col > m.columns-1 {
		return false
	}
	if t == len(m.freedots) {
		m.ComputeFreeDots(t)
	}
	if t < len(m.freedots) {
		c := Coords{row: row, column: col}
		v, ok := (*m.freedots[t])[c]
		return v && ok
	}
	panic("Dynamic programming error - we should not be here")
}

type Generation struct {
	points map[Coords]bool
	time   int
}

func (g *Generation) ReachedEnd(m *Map) bool {
	for c, _ := range g.points {
		if c.column == m.columns-1 && c.row == m.rows-1 {
			return true
		}
	}
	return false
}

func (g *Generation) ReachedStart(m *Map) bool {
	for c, _ := range g.points {
		if c.column == 0 && c.row == 0 {
			return true
		}
	}
	return false

}

func GetFirstGeneration(m *Map) *Generation {
	res := Generation{points: make(map[Coords]bool), time: 0}
	canStart := false
	t := 0
	for !canStart {
		if m.IsFree(t, 0, 0) {
			res.time = t
			res.points[Coords{row: 0, column: 0}] = true
			canStart = true
		}
		t++
	}
	return &res
}

func (g *Generation) GetNextGeneration(m *Map) *Generation {
	res := Generation{points: make(map[Coords]bool), time: g.time + 1}
	for c, _ := range g.points {
		if m.IsFree(res.time, c.row, c.column) {
			res.points[c] = true
		}
		if m.IsFree(res.time, c.row+1, c.column) {
			c1 := Coords{c.row + 1, c.column}
			res.points[c1] = true
		}
		if m.IsFree(res.time, c.row-1, c.column) {
			c1 := Coords{c.row - 1, c.column}
			res.points[c1] = true
		}
		if m.IsFree(res.time, c.row, c.column+1) {
			c1 := Coords{c.row, c.column + 1}
			res.points[c1] = true
		}
		if m.IsFree(res.time, c.row, c.column-1) {
			c1 := Coords{c.row, c.column - 1}
			res.points[c1] = true
		}
	}
	return &res
}

func processPart1(filename string) {
	m := Init(filename)
	g := GetFirstGeneration(m)
	for !g.ReachedEnd(m) {

		g = g.GetNextGeneration(m)
	}
	println(g.time + 1)
}

func processPart2(filename string) {
	m := Init(filename)
	g := GetFirstGeneration(m)
	for !g.ReachedEnd(m) {
		g = g.GetNextGeneration(m)
	}
	// Turning back
	for c := range g.points {
		delete(g.points, c)
	}
	g.points[Coords{row: m.rows, column: m.columns - 1}] = true
	for !g.ReachedStart(m) {
		g = g.GetNextGeneration(m)
	}
	// println(g.time)
	for c := range g.points {
		delete(g.points, c)
	}
	g.points[Coords{row: -1, column: 0}] = true
	for !g.ReachedEnd(m) {
		g = g.GetNextGeneration(m)
	}
	println(g.time + 1)
}

func main() {
	processPart1("test.txt")
	processPart1("input.txt")
	processPart2("test.txt")
	processPart2("input.txt")
}
