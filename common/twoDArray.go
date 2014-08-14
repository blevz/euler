package common

type TDArray struct {
	data          []interface{}
	height, width int
}

func MakeTDArray(h, w int) *TDArray {
	var t TDArray
	t.data = make([]interface{}, h*w)
	t.height = h
	t.width = w
	return &t
}

func (t TDArray) GetSize() (int, int) {
	return t.height, t.width
}

func (t TDArray) checkBounds(x, y int) bool {
	//Returns true if in bounds
	if (x >= 0 && x < t.width) && (y >= 0 && y < t.height) {
		return true
	} else {
		return false
	}
}

func (t TDArray) GetElemVal(x, y int) interface{} {
	if t.checkBounds(x, y) {
		return t.data[y*t.width+x]
	} else {
		panic("out of bounds")
	}
}

func (t *TDArray) SetElem(x, y int, val interface{}) {
	if t.checkBounds(x, y) {
		t.data[y*t.width+x] = val
	}
}

func (t *TDArray) GetElemRef(x, y int) *interface{} {
	if t.checkBounds(x, y) {
		return &t.data[y*t.width+x]
	} else {
		return nil
	}
}
