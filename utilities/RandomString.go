package utilities

import (
	"math/rand"
	"slices"
)

func CreateRandomString(length int) string {
	ran_str := make([]byte, length)
	for i := 0; i < length; i++ {
		if rand.Intn(length)%2 == 0 {
			ran_str[i] = byte(65 + rand.Intn(25))
		} else {
			ran_str[i] = byte(97 + rand.Intn(25))
		}

	}

	return string(ran_str)
}

func CreateRandomIntArray(i, base int) []int {
	if i == 0 {
		return []int{}
	}
	arr := CreateRandomIntArray(i-1, base)
	tmp := rand.Intn(base) + 1
	for slices.Contains(arr, tmp) {
		tmp = rand.Intn(base) + 1
	}
	return append(arr, tmp)
}
