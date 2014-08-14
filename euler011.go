package main

import (
	"fmt"
	"github.com/blevz/euler/common"

	"os"
)

func FileToTwoDArray(path string, x, y int) *common.TDArray {
	toReturn := common.MakeTDArray(x, y)

	dataFile, err := os.Open("euler011_data.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer dataFile.Close()
	var val int
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			fmt.Fscanf(dataFile, "%d", &val)

			toReturn.SetElem(i, j, val)
		}
	}

	return toReturn
}

func main() {
	input := make(chan int)
	tdarray := FileToTwoDArray("euler011_data.txt", 20, 20)
	width, height := tdarray.GetSize()

	go func() {
		//Left/ right
		for y := 0; y < height; y++ {
			for x := 0; x < width-4; x++ {
				product := 1
				for i := 0; i < 4; i++ {

					val := tdarray.GetElemVal(x+i, y)
					product *= val.(int)
				}
				input <- product
			}
		}
		//Up/ Down
		for y := 0; y < height-4; y++ {
			for x := 0; x < width; x++ {
				product := 1
				for i := 0; i < 4; i++ {
					val := tdarray.GetElemVal(x, y+i)
					product *= val.(int)
				}
				input <- product
			}
		}

		//diagonal both add
		for y := 0; y < height-4; y++ {
			for x := 0; x < height-4; x++ {
				product := 1
				for i := 0; i < 4; i++ {
					val := tdarray.GetElemVal(x+i, y+i)
					product *= val.(int)
				}
				input <- product
			}
		}
		//diagonal one add/ one subtracts
		for y := 0; y < height-4; y++ {
			for x := 4; x < height; x++ {
				product := 1
				for i := 0; i < 4; i++ {
					val := tdarray.GetElemVal(x-i, y+i)
					product *= val.(int)
				}
				input <- product
			}
		}
		close(input)
	}()

	max := common.Max_int(input)
	fmt.Println("max is: ", max)

}
