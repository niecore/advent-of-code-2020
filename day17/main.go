package main

import (
	util "awesomeProject/util"
	"fmt"
)

type Cube struct {
	x, y, z, w int
	active     bool
}

type Cubes struct {
	data       []*Cube
	w, h, d, t int
}

func main() {
	cubes := parseInitialCubesLayout("day17/input.txt")
	cubes2 := cubes

	for i := 0; i < 6; i++ {
		_cubes := predictNewSeatLayout(cubes, false)
		_cubes2 := predictNewSeatLayout(cubes2, true)

		cubes = _cubes
		cubes2 = _cubes2
	}
	println(countActiveCubes(cubes))
	println(countActiveCubes(cubes2))
}

func predictNewSeatLayout(cubes *Cubes, fourthDimension bool) *Cubes {
	var newLayout *Cubes
	if fourthDimension {
		newLayout = New(cubes.w+2, cubes.h+2, cubes.d+2, cubes.t+2)
	} else {
		newLayout = New(cubes.w+2, cubes.h+2, cubes.d+2, 1)
	}

	for w := -1; w < cubes.t+1; w++ {

		if !fourthDimension && w != 0 {
			continue
		}

		for z := -1; z < cubes.d+1; z++ {
			for y := -1; y < cubes.h+1; y++ {
				for x := -1; x < cubes.w+1; x++ {
					cube := cubes.At(x, y, z, w)
					activeNeighbors := activeAdjecentCubes(cubes, x, y, z, w)

					var newState bool
					if cube != nil && cube.active {
						newState = activeNeighbors == 3 || activeNeighbors == 2
					} else {
						newState = activeNeighbors == 3
					}

					if fourthDimension {
						newLayout.Set(x+1, y+1, z+1, w+1, newState)
					} else {
						newLayout.Set(x+1, y+1, z+1, w, newState)
					}
				}
			}
		}
	}

	return newLayout
}

func activeAdjecentCubes(cubes *Cubes, x, y, z, w int) int {
	activeCount := 0
	for _, adjecentCube := range cubes.Adjecent(x, y, z, w) {
		if adjecentCube != nil && adjecentCube.active {
			activeCount++
		}
	}
	return activeCount
}

func countActiveCubes(cubes *Cubes) int {
	activeCount := 0
	for _, cube := range cubes.data {
		if cube.active {
			activeCount++
		}
	}
	return activeCount
}

func New(w, h, d, t int) *Cubes {
	cubes := &Cubes{
		data: make([]*Cube, w*h*d*t),
		w:    w,
		h:    h,
		d:    d,
		t:    t,
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			for z := 0; z < d; z++ {
				for w := 0; w < t; w++ {
					cubes.Init(x, y, z, w, &Cube{x, y, z, w, false})
				}
			}
		}
	}

	return cubes
}

func (t *Cubes) At(x, y, z, w int) *Cube {
	if x >= t.w || y >= t.h || z >= t.d || w >= t.t || x < 0 || y < 0 || z < 0 || w < 0 {
		return nil
	}

	idx := t.t*t.h*t.w*w + t.h*t.w*z + t.w*y
	return t.data[idx+x]
}

func (t *Cubes) Set(x, y, z, w int, active bool) {
	idx := t.t*t.h*t.w*w + t.h*t.w*z + t.w*y
	t.data[idx+x].active = active
}

func (t *Cubes) Init(x, y, z, w int, val *Cube) {
	idx := t.t*t.h*t.w*w + t.h*t.w*z + t.w*y
	t.data[idx+x] = val
}

func (t *Cubes) Adjecent(x, y, z, w int) []*Cube {
	cubes := make([]*Cube, (3*3*3*3)-1)
	idx := 0
	for z_d := -1; z_d <= 1; z_d++ {
		for y_d := -1; y_d <= 1; y_d++ {
			for x_d := -1; x_d <= 1; x_d++ {
				for w_d := -1; w_d <= 1; w_d++ {
					if x_d == 0 && z_d == 0 && y_d == 0 && w_d == 0 {
						continue
					}
					cubes[idx] = t.At(x+x_d, y+y_d, z+z_d, w+w_d)
					idx++
				}
			}
		}
	}

	return cubes
}

func (t *Cubes) printCubes() {
	for w := 0; w < t.t; w++ {
		for z := 0; z < t.d; z++ {
			fmt.Printf("z=%d w=%d\n", z, w)
			for y := 0; y < t.h; y++ {
				for x := 0; x < t.w; x++ {

					if t.At(x, y, z, 0).active {
						print("#")
					} else {
						print(".")
					}
				}
				println()
			}
		}
	}
}

func parseInitialCubesLayout(file string) *Cubes {
	rows := util.ReadInput(file)
	cubes := New(len(rows[0]), len(rows), 1, 1)

	for y, row := range rows {
		for x, column := range row {
			cubes.Set(x, y, 0, 0, column == '#')
		}
	}
	return cubes
}
