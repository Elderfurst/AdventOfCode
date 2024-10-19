package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	input := "yzbqklnj"

	counter := 1
	firstValue := 0
	secondValue := 0
	foundFirst := false
	foundSecond := false

	for !foundFirst || !foundSecond {
		test := fmt.Sprintf("%s%d", input, counter)

		hash := md5.Sum([]byte(test))
		result := hex.EncodeToString(hash[:])

		if strings.HasPrefix(result, "00000") && !foundFirst {
			firstValue = counter
			foundFirst = true
		}

		if strings.HasPrefix(result, "000000") {
			secondValue = counter
			foundSecond = true
		}

		counter++
	}

	fmt.Println(firstValue)
	fmt.Println(secondValue)
}
