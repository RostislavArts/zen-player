package utils

import (
	"math/rand"
)

func ShuffleSlice(slice *[]string) {
	rand.Shuffle(len(*slice), func(i, j int) {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	})
}

