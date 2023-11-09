package utilities

import (
	"math/rand"
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
