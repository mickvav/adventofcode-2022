package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coords struct {
	row    int
	column int
}

type Elve struct {
	coords    Coords
	nextpoint *Point
}
type Point struct {
	coords Coords
	elves  []*Elve
}

type Elves struct {
	elves  map[Coords]*Elve
	points map[Coords]*Point
}

func Init(filename string) *Elves {
	res := Elves{points: make(map[Coords]*Point), elves: make(map[Coords]*Elve)}
	f, err := os.Open(filename)
	if err != nil {
		panic("Load error")
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	row := 0
	for fs.Scan() {
		for i, v := range fs.Text() {
			if v == '#' {
				c := Coords{row: row, column: i}
				e := Elve{coords: c, nextpoint: nil}
				res.elves[c] = &e
			}
		}
		row++
	}
	return &res
}

func Directions(step int) string {
	switch step % 4 {
	case 0:
		return "NSWE"
	case 1:
		return "SWEN"
	case 2:
		return "WENS"
	case 3:
		return "ENSW"
	}
	return "nonesense"
}

func (es *Elves) HasNeighbours(e *Elve) bool {
	var p Coords
	p = Coords{row: e.coords.row + 1, column: e.coords.column}
	_, ok := es.elves[p]
	if ok {
		return true
	}

	p = Coords{row: e.coords.row + 1, column: e.coords.column + 1}
	_, ok = es.elves[p]
	if ok {
		return true
	}
	p = Coords{row: e.coords.row + 1, column: e.coords.column - 1}
	_, ok = es.elves[p]
	if ok {
		return true
	}
	p = Coords{row: e.coords.row, column: e.coords.column + 1}
	_, ok = es.elves[p]
	if ok {
		return true
	}
	p = Coords{row: e.coords.row, column: e.coords.column - 1}
	_, ok = es.elves[p]
	if ok {
		return true
	}
	p = Coords{row: e.coords.row - 1, column: e.coords.column}
	_, ok = es.elves[p]
	if ok {
		return true
	}
	p = Coords{row: e.coords.row - 1, column: e.coords.column + 1}
	_, ok = es.elves[p]
	if ok {
		return true
	}
	p = Coords{row: e.coords.row - 1, column: e.coords.column - 1}
	_, ok = es.elves[p]
	if ok {
		return true
	}
	return false
}

func (es *Elves) isElveAt(row int, column int) bool {
	p := Coords{row: row, column: column}
	_, ok := es.elves[p]
	return ok
}

func (es *Elves) HasNeighboursAt(e *Elve, d rune) bool {
	switch d {
	case 'N':
		return es.isElveAt(e.coords.row-1, e.coords.column) || es.isElveAt(e.coords.row-1, e.coords.column+1) || es.isElveAt(e.coords.row-1, e.coords.column-1)
	case 'S':
		return es.isElveAt(e.coords.row+1, e.coords.column) || es.isElveAt(e.coords.row+1, e.coords.column+1) || es.isElveAt(e.coords.row+1, e.coords.column-1)
	case 'E':
		return es.isElveAt(e.coords.row+1, e.coords.column+1) || es.isElveAt(e.coords.row, e.coords.column+1) || es.isElveAt(e.coords.row-1, e.coords.column+1)
	case 'W':
		return es.isElveAt(e.coords.row+1, e.coords.column-1) || es.isElveAt(e.coords.row, e.coords.column-1) || es.isElveAt(e.coords.row-1, e.coords.column-1)
	}
	panic(fmt.Sprintf("Wrong direction: %c", d))
}

func (es *Elves) Propose(e *Elve, d rune) {
	c := Coords{row: e.coords.row, column: e.coords.column}
	switch d {
	case 'N':
		c.row -= 1
	case 'S':
		c.row += 1
	case 'E':
		c.column += 1
	case 'W':
		c.column -= 1
	}
	p, ok := es.points[c]
	if !ok {
		// Create new point
		p = &Point{coords: c, elves: []*Elve{e}}
		e.nextpoint = p
		es.points[c] = p
	} else {
		p.elves = append(p.elves, e)
		e.nextpoint = p
	}
}

func (es *Elves) FirstPart(step int) {
	d := Directions(step)
	for _, e := range es.elves {
		e.nextpoint = nil
		if !es.HasNeighbours(e) {
			continue
		}
		for _, dir := range d {
			if !es.HasNeighboursAt(e, dir) {
				es.Propose(e, dir)
				break
			}
		}
	}
}

func (es *Elves) move(e *Elve) {
	if e.nextpoint != nil {
		delete(es.elves, e.coords)
		e.coords.column = e.nextpoint.coords.column
		e.coords.row = e.nextpoint.coords.row
		es.elves[e.coords] = e
		e.nextpoint = nil
	}
}

func (es *Elves) SecondPart() int {
	moves := 0
	for _, e := range es.elves {
		if e.nextpoint == nil {
			continue
		}
		if len(e.nextpoint.elves) == 1 {
			es.move(e)
			moves += 1
		}
	}
	for c := range es.points {
		delete(es.points, c)
	}
	return moves
}

func (es *Elves) BBox() (int, int, int, int) {
	mincolumn := 0
	maxcolumn := 0
	minrow := 0
	maxrow := 0
	for c, _ := range es.elves {
		mincolumn = c.column
		maxcolumn = c.column
		minrow = c.row
		maxrow = c.row
		break
	}
	for c, _ := range es.elves {
		if c.column > maxcolumn {
			maxcolumn = c.column
		}
		if c.row > maxrow {
			maxrow = c.row
		}
		if c.column < mincolumn {
			mincolumn = c.column
		}
		if c.row < minrow {
			minrow = c.row
		}
	}
	return minrow, maxrow, mincolumn, maxcolumn
}

func (es *Elves) Print() {
	minrow, maxrow, mincolumn, maxcolumn := es.BBox()
	dots := 0
	for r := minrow; r <= maxrow; r++ {
		line := ""
		for c := mincolumn; c <= maxcolumn; c++ {
			if es.isElveAt(r, c) {
				line += "#"
			} else {
				line += "."
				dots++
			}
		}
		println(line)
	}
	println("Dots: ", dots)
}

func processPart1(filename string) {
	es := Init(filename)
	for s := 0; s < 10; s++ {
		es.FirstPart(s)
		es.SecondPart()
	}
	es.Print()
}

func processPart2(filename string) {
	es := Init(filename)
	s := 0
	moves := 1
	for moves > 0 {
		es.FirstPart(s)
		moves = es.SecondPart()
		//es.Print()
		s++
	}
	print("Part 2:", filename, s)
}

func main() {
	processPart1("test.txt")
	processPart1("input.txt")
	processPart2("test.txt")
	processPart2("input.txt")
}
