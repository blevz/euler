package common

import (
	"encoding/csv"
	"fmt"
	"os"
)

func genPerms(q []byte, s []byte, output chan<- string) {
	size := len(q)
	if size == 0 {
		output <- string(s)
		return
	}
	for k := 0; k < size; k++ {
		s = append(s, q[0])
		q = q[1:len(q)]
		genPerms(q, s, output)
		q = append(q, s[len(s)-1])
		s = s[0 : len(s)-1]
	}

}

func ProducePermutations(starting string, output chan<- string) {
	q := []byte(starting)
	s := []byte("")
	genPerms(q, s, output)
	close(output)
}

func GetAllCSVStrings(pathname string) ([]string, error) {
	file, err := os.Open(pathname)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}
	reader := csv.NewReader(file)
	strArr, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}
	return strArr, nil
}

func GetStringVal(str string) int {
	sum := 0
	for _, r := range str {
		sum += int(r - 'A' + 1)
	}
	return sum
}
