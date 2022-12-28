package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	row    int
	column int
	wall   bool
	trace  int
	right  *Point
	left   *Point
	top    *Point
	bottom *Point
}

type Row struct {
	points []*Point
	start  int
	stop   int
}

type Map struct {
	rows    []Row
	program string
}

type RotatorKey struct {
	row    int
	column int
	facing int
}

type Walker struct {
	point    *Point
	facing   int                // Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^).
	rotators map[RotatorKey]int // Map of new facings
}

func (w *Walker) Step() {
	if w.facing == 0 {
		if !w.point.right.wall {
			w.point = w.point.right
		}
	}
	if w.facing == 1 {
		if !w.point.bottom.wall {
			w.point = w.point.bottom
		}
	}
	if w.facing == 2 {
		if !w.point.left.wall {
			w.point = w.point.left
		}
	}
	if w.facing == 3 {
		if !w.point.top.wall {
			w.point = w.point.top
		}
	}
}

func (w *Walker) Step2() {
	rk := RotatorKey{column: w.point.column, row: w.point.row, facing: w.facing}
	newfacing, ok := w.rotators[rk]
	if !ok {
		newfacing = w.facing
	}
	w.point.trace = 4 + w.facing
	n := w.point.Next(w.facing)
	if n.wall {
		return
	}
	w.point = n
	w.facing = newfacing
	w.point.trace = 4 + newfacing
}

func (p *Point) Next(facing int) *Point {
	switch facing {
	case 0:
		return p.right
	case 1:
		return p.bottom
	case 2:
		return p.left
	case 3:
		return p.top
	}
	panic(fmt.Sprintf("Wrong facing: %d", facing))
}

func (p *Point) SetNext(facing int, p1 *Point) {
	switch facing {
	case 0:
		p.right = p1
	case 1:
		p.bottom = p1
	case 2:
		p.left = p1
	case 3:
		p.top = p1
	default:
		panic(fmt.Sprintf("Wrong facing: %d", facing))
	}
}

func (w *Walker) stichEdges(m *Map, s1 RotatorKey, s2 RotatorKey, f1 int, f2 int, len int) error {
	r1 := m.rows[s1.row-1]
	if s1.column < r1.start || s1.column > r1.stop {
		return fmt.Errorf("No such point (s1): %d %d", s1.row, s1.column)
	}
	r2 := m.rows[s2.row-1]
	if s2.column < r2.start || s2.column > r2.stop {
		return fmt.Errorf("No such point (s2): %d %d", s2.row, s2.column)
	}
	p1 := r1.points[s1.column-r1.start]
	p2 := r2.points[s2.column-r2.start]
	for i := 0; i < len; i++ {
		rk1 := RotatorKey{column: p1.column, row: p1.row, facing: s1.facing}
		rk2 := RotatorKey{column: p2.column, row: p2.row, facing: s2.facing}
		p1.SetNext(s1.facing, p2)
		p2.SetNext(s2.facing, p1)
		w.rotators[rk1] = (s2.facing + 2) % 4
		w.rotators[rk2] = (s1.facing + 2) % 4
		p1 = p1.Next(f1)
		p2 = p2.Next(f2)
	}
	return nil
}

func (w *Walker) TurnRight() {
	w.facing = (w.facing + 1) % 4
}

func (w *Walker) TurnLeft() {
	w.facing = (4 + w.facing - 1) % 4
}

func (r Row) String() string {
	facingstrings := []string{">", "v", "<", "^"}
	res := ""
	for i := 0; i < r.start; i++ {
		res += " "
	}
	for i := r.start; i <= r.stop; i++ {
		if r.points[i-r.start].wall {
			res += "#"
		} else {
			t := r.points[i-r.start].trace
			if t > 0 {
				res += facingstrings[t%4]
			} else {
				res += "."
			}
		}
	}
	return res
}
func (m *Map) Print() {
	for i, r := range m.rows {
		print(r.String(), " =====", i, "\n")
	}
}

func (m *Map) GetPoint(row int, column int) (*Point, error) {
	rowi := row - 1
	if row > len(m.rows) {
		rowi = (row - 1) % len(m.rows)
	}
	if row < 1 {
		rowi = len(m.rows) + (row-1)%len(m.rows)
	}
	r := m.rows[rowi]
	if column < r.start {
		ci := column - r.start
		ci = len(r.points) + ci%len(r.points)
		return r.points[ci], nil
	}
	if column > r.stop {
		ci := column - r.start
		ci = ci % len(r.points)
		return r.points[ci], nil
	}
	ci := column - r.start
	return r.points[ci], nil
}

func ParseMapLine(row int, line string) (Row, error) {
	// "...#.......#"
	// "   ....#"
	res := Row{}
	for i, v := range line {
		if v == ' ' {
			continue
		}
		if res.start == 0 {
			res.start = i + 1
		}
		res.stop = i + 1
		P := Point{row: row, column: i + 1, top: nil, bottom: nil}
		if v == '#' {
			P.wall = true
		} else {
			P.wall = false
		}
		res.points = append(res.points, &P)
	}
	for i, P := range res.points {
		if i == 0 {
			P.left = res.points[len(res.points)-1]
		} else {
			P.left = res.points[i-1]
		}
		if i == len(res.points)-1 {
			P.right = res.points[0]
		} else {
			P.right = res.points[i+1]
		}
	}
	return res, nil
}

func LinkRows(m Map) {
	maxstop := 0
	for _, r := range m.rows {
		if maxstop < r.stop {
			maxstop = r.stop
		}
	}
	for ci := 1; ci <= maxstop; ci++ {
		var firstrow, lastrow int
		firstrow = -1
		lastrow = -1
		for ri, r := range m.rows {
			if ci >= r.start && ci <= r.stop {
				if firstrow == -1 {
					firstrow = ri
				}
				if lastrow != -1 {
					cip1 := ci - m.rows[lastrow].start
					p1 := m.rows[lastrow].points[cip1]
					p2 := r.points[ci-r.start]
					p1.bottom = p2
					p2.top = p1
				}
				lastrow = ri
			}
		}
		cip1 := ci - m.rows[firstrow].start
		cip2 := ci - m.rows[lastrow].start
		p1 := m.rows[firstrow].points[cip1]
		p2 := m.rows[lastrow].points[cip2]
		p1.top = p2
		p2.bottom = p1
	}
}

func ParseMap(filename string) Map {
	var res Map
	f, err := os.Open(filename)
	if err != nil {
		panic("Load error")
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	scanningMap := true
	row := 1
	for fs.Scan() {
		if fs.Text() != "" && scanningMap {
			line, _ := ParseMapLine(row, fs.Text())
			res.rows = append(res.rows, line)
			row += 1
		}
		if fs.Text() == "" {
			LinkRows(res)
			scanningMap = false
			continue
		}
		res.program = fs.Text()
	}
	return res
}

func walk(m Map, w Walker) (int, int, int) {
	r := regexp.MustCompile("[0-9]+|R|L")
	matches := r.FindAllString(m.program, -1)
	for _, op := range matches {
		if op == "R" {
			w.TurnRight()
		} else if op == "L" {
			w.TurnLeft()
		} else {
			distance, _ := strconv.Atoi(op)
			for s := 0; s < distance; s++ {
				w.Step2()
			}
		}
		// print(w.point.row, " ", w.point.column, "\n")
		//		m.Print()
	}
	return w.point.row, w.point.column, w.facing
}

func processPart1(filename string) {
	fmt.Printf("Part1 %s\n", filename)
	m := ParseMap(filename)
	m.Print()
	w := Walker{point: m.rows[0].points[0], facing: 0, rotators: map[RotatorKey]int{}}
	row, col, facing := walk(m, w)
	print(1000*row + 4*col + facing)
}

type stichOpts struct {
	s1  RotatorKey
	s2  RotatorKey
	f1  int
	f2  int
	len int
}

func processPart2(filename string) {
	stitchmap := map[string][]stichOpts{
		//    Bo----oC
		//    /|   /|
		//   / |  / |
		// Ao----oD |
		//  | Fo-|--oG
		//  | /  | /
		//  |/   |/
		// Eo----oH
		//
		//
		//          B .- C
		//          .    .
		//C .- B -. A -- D
		//.    |    |    .
		//G .- F .- E -- H .- D
		//          .    |    .
		//          F .- G .- C
		"test.txt": {
			stichOpts{ // B -> A
				s1:  RotatorKey{row: 1, column: 9, facing: 2},
				f1:  1,
				s2:  RotatorKey{row: 5, column: 5, facing: 3},
				f2:  0,
				len: 4,
			},
			stichOpts{ // B -> C
				s1:  RotatorKey{row: 1, column: 9, facing: 3},
				f1:  0,
				s2:  RotatorKey{row: 5, column: 4, facing: 3},
				f2:  2,
				len: 4,
			},
			stichOpts{ // H -> D
				s1:  RotatorKey{row: 8, column: 12, facing: 0},
				f1:  3,
				s2:  RotatorKey{row: 9, column: 13, facing: 3},
				f2:  0,
				len: 4,
			},
			stichOpts{ // D -> C
				s1:  RotatorKey{row: 4, column: 12, facing: 0},
				f1:  3,
				s2:  RotatorKey{row: 9, column: 16, facing: 0},
				f2:  1,
				len: 4,
			},
			stichOpts{ // G -> C
				s1:  RotatorKey{row: 8, column: 1, facing: 2},
				f1:  3,
				s2:  RotatorKey{row: 12, column: 13, facing: 1},
				f2:  0,
				len: 4,
			},
			stichOpts{ // G -> F
				s1:  RotatorKey{row: 8, column: 1, facing: 1},
				f1:  0,
				s2:  RotatorKey{row: 12, column: 12, facing: 1},
				f2:  2,
				len: 4,
			},
			stichOpts{ // F -> E
				s1:  RotatorKey{row: 8, column: 5, facing: 1},
				f1:  0,
				s2:  RotatorKey{row: 12, column: 9, facing: 2},
				f2:  3,
				len: 4,
			},
		},
		"input.txt": {
			//
			//    Bo----oC
			//    /|   /|
			//   / |  / |
			// Ao----oD |
			//  | Fo-|--oG
			//  | /  | /
			//  |/   |/
			// Eo----oH
			//             150|
			//        100|101
			// 1  50|51
			//      B -- C -- G      1
			//      |    |    |
			//      A -- D -- H      50/51
			//      |    |
			// A -- E -- H           100/101
			// |    |    |
			// B -- F -- G           150/151
			// |    |
			// C -- G                200
			//
			stichOpts{ // B -> C
				s1:  RotatorKey{row: 1, column: 51, facing: 3},
				f1:  0,
				s2:  RotatorKey{row: 151, column: 1, facing: 2},
				f2:  1,
				len: 50,
			},
			stichOpts{ // B -> A
				s1:  RotatorKey{row: 1, column: 51, facing: 2},
				f1:  1,
				s2:  RotatorKey{row: 150, column: 1, facing: 2},
				f2:  3,
				len: 50,
			},
			stichOpts{ // A -> E
				s1:  RotatorKey{row: 51, column: 51, facing: 2},
				f1:  1,
				s2:  RotatorKey{row: 101, column: 1, facing: 3},
				f2:  0,
				len: 50,
			},
			stichOpts{ // C -> G
				s1:  RotatorKey{row: 200, column: 1, facing: 1},
				f1:  0,
				s2:  RotatorKey{row: 1, column: 101, facing: 3},
				f2:  0,
				len: 50,
			},
			stichOpts{ // G -> F
				s1:  RotatorKey{row: 200, column: 50, facing: 0},
				f1:  3,
				s2:  RotatorKey{row: 150, column: 100, facing: 1},
				f2:  2,
				len: 50,
			},
			stichOpts{ // G -> H
				s1:  RotatorKey{row: 150, column: 100, facing: 0},
				f1:  3,
				s2:  RotatorKey{row: 1, column: 150, facing: 0},
				f2:  1,
				len: 50,
			},
			stichOpts{ // H -> D
				s1:  RotatorKey{row: 100, column: 100, facing: 0},
				f1:  3,
				s2:  RotatorKey{row: 50, column: 150, facing: 1},
				f2:  2,
				len: 50,
			},
		},
	}
	fmt.Printf("Part2 %s\n", filename)
	fmt.Printf("Part2: %d\n", 0)
	m := ParseMap(filename)
	w := Walker{point: m.rows[0].points[0], facing: 0, rotators: map[RotatorKey]int{}}

	for _, so := range stitchmap[filename] {
		err := w.stichEdges(&m, so.s1, so.s2, so.f1, so.f2, so.len)
		if err != nil {
			print(err)
		}
	}
	row, col, facing := walk(m, w)
	m.Print()
	print(1000*row + 4*col + facing)
}

func main() {
	//	processPart1("test.txt")
	//	processPart1("input.txt")
	processPart2("test.txt")
	processPart2("input.txt")
}
