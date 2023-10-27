package utilities

import (
	"golang.org/x/crypto/sha3"
)

func HashUser(ID int, Username string) []byte {
	sbyte := []byte(Username)
	shs := sha3.Sum256(sbyte)
	for i := 0; i < ID%5; i++ {
		shs = sha3.Sum256(shs[:])
	}
	return shs[:]
}

func HashKamil(state string) []byte {
	sbyte := []byte(state)
	shs := sha3.Sum256(sbyte)
	for i := 0; i < len(sbyte); i++ {
		shs = sha3.Sum256(shs[:])
	}
	return shs[:]
}
