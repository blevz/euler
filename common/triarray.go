package common

import (
	"fmt"
	"os"
)

type TriangleArray struct {
	data [][]interface{}
}

func MakeTriangleArray() *TriangleArray {
	var toReturn TriangleArray
	toReturn.data = make([][]interface{}, 0)
	return &toReturn
}

func MakeTriangleArrayWithHeight(h int) *TriangleArray {
	var toReturn TriangleArray
	toReturn.data = make([][]interface{}, h)
	for x := 0; x < h; x++ {
		toReturn.data[x] = make([]interface{}, x+1)
		for y := 0; y <= x; y++ {
			toReturn.data[x] = append(toReturn.data[x], nil)
		}
	}
	return &toReturn
}

func FileToTriangleArray(filepath string) *TriangleArray {
	toReturn := MakeTriangleArray()
	dataFile, err := os.Open(filepath)
	if err != nil {
		panic("open problem")
	}
	defer dataFile.Close()
	var val int
	err = nil
	h := 0
	x := 0
	for {
		_, err = fmt.Fscanf(dataFile, "%d", &val)
		if err != nil {
			break
		}
		toReturn.SetVal(h, x, val)
		//fmt.Println(h, x, val)
		x++
		if x > h {
			h++
			x = 0
		}
	}

	return toReturn
}

func (tri *TriangleArray) growToDepth(depth int) {
	for h := len(tri.data); h < depth+1; h++ {
		tri.data = append(tri.data, make([]interface{}, h))
		for i := 0; i < h+1; i++ {
			tri.data[h] = append(tri.data[h], nil)
		}
	}
}

func (tri *TriangleArray) SetVal(depth, x int, val interface{}) {
	if len(tri.data) < depth+1 {
		tri.growToDepth(depth)
	}
	tri.data[depth][x] = val
}

func (tri TriangleArray) GetVal(depth, x int) interface{} {
	return tri.data[depth][x]
}

func (tri TriangleArray) GetDepth() int {
	return len(tri.data)
}
