package main

import (
	"fmt"
	"sort"
	"strconv"
)


func main() {

	result := MergeSlices([]float32{1.1, 2.1, 3.1}, []int32{4, 5})
	fmt.Println(result)

	res := GetMapValuesSortedByKey(map[int]string{
				9:  "Сентябрь",
				1:  "Январь",
				2:  "Февраль",
				10: "Октябрь",
				5:  "Май",
				7:  "Июль",
				8:  "Август",
				12: "Декарь",
				3:  "Март",
				6:  "Июнь",
				4:  "Апрель",
				11: "Ноябрь",
			})
	fmt.Println(res)
}

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return float32(1.1)
}

func ReturnIntArray() [3]int {
	array := [3]int{1, 3, 4}
	return array
}

func ReturnIntSlice() []int {
	sl1 := []int{1, 2, 3}
	return sl1
}

func IntSliceToString(intSl []int) string {
	var str string
	for _, val := range intSl {
		str += strconv.Itoa(val)
	}

	return str
}

func MergeSlices(inpFloat []float32, inpInt []int32) []int {
	size := len(inpFloat) + len(inpInt)
	resSl := make([]int, size)
	probSl := make([]int, 0, size)

	for i := 0; i < size; i++ {
		if i < len(inpFloat) {
			resSl[i] = int(inpFloat[i])
			probSl = append(probSl, int(inpFloat[i]))
		}
		if i < len(inpInt) {
			resSl[i+len(inpFloat)] = int(inpInt[i])
		}

	}

	return resSl
}

func GetMapValuesSortedByKey(inputMap map[int]string) []string {
	keySl := make([]int, 0, len(inputMap))
	resSl := make([]string, 0, len(inputMap))

	for key := range inputMap {
		keySl = append(keySl, key)
	}

	sort.Ints(keySl)

	for _, key := range keySl {
		resSl = append(resSl, inputMap[key])
	}

	return resSl
}
