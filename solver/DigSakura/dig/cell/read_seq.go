package cell

import (
	"math/rand"
	"sort"
	"time"
)

func ReadSeq() []int {
	radom := RandInt64(3, 29)
	offset := 0
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		offset += int(radom)
	} else {
		offset -= int(radom)
	}
	originSeq := []int{64, 31, 97, 16, 47, 79, 114, 7, 25, 38, 56, 71, 88, 106, 123, 3, 12, 20, 28, 35, 44, 52, 61, 67, 75, 84, 92, 101, 110, 119, 126, 1, 5, 10, 14, 18, 23, 26, 30, 33, 37, 41, 46, 49, 54, 58, 63, 65, 69, 73, 77, 81, 86, 90, 95, 99, 104, 108, 112, 116, 121, 124, 0, 2, 4, 6, 8, 11, 13, 15, 17, 19, 22, 24, 27, 29, 32, 34, 36, 39, 43, 45, 48, 50, 53, 55, 57, 59, 62, 66, 68, 70, 72, 74, 76, 78, 80, 82, 85, 87, 89, 91, 94, 96, 98, 100, 103, 105, 107, 109, 111, 113, 115, 118, 120, 122, 125, 127, 9, 21, 40, 42, 51, 60, 83, 93, 102, 117}
	targetSeq := []int{}
	for _, elem := range originSeq {
		// 计算偏移量
		tragetNum := (128 + elem + offset) % 128
		targetSeq = append(targetSeq, tragetNum)
	}
	return targetSeq
}

func TestSeq() {
	data := ReadSeq()
	sort.Ints(data)

	for _, elem := range data {
		print(elem, "  ")
	}
}

func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixMilli())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
