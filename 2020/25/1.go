package main

import "fmt"

func main() {
	cardPublicKey, doorPublicKey := 3418282, 8719412
	cardLoopSize := transformUntil(7, cardPublicKey)
	encryptionKey := transform(doorPublicKey, cardLoopSize)
	fmt.Println(encryptionKey)
}

func transform(subject int, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subject
		value %= 20201227
	}
	return value
}

func transformUntil(subject int, publicKey int) int {
	value := 1
	loopSize := 1
	for {
		value *= subject
		value %= 20201227
		if value == publicKey {
			return loopSize
		}
		loopSize++
	}
}
