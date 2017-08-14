package random_crockford

import (
	"math/rand"
	"strings"
	"time"
)

const sep = ""

var dictionary = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h",
	/* "i", */ // No i
	"j", "k",
	/* "l", */ // No l
	"m", "n",
	/* "o", */ // No o
	"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var dictionarySize = len(dictionary)

type RandomCrockford struct {
	radix int
}

const empty = ""

// Next returns next random string
func (r *RandomCrockford) Next() string {
	var rolls []string

	if r.radix == 0 {
		return empty
	}

	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	rolls = make([]string, r.radix)

	for i := 0; i < r.radix; i++ {
		rolls[i] = dictionary[r1.Intn(dictionarySize)]
	}

	return strings.Join(rolls, sep)
}

func NewRandom(radix int) *RandomCrockford {

	if radix < 0 {
		radix = 0
	}

	return &RandomCrockford{radix: radix}
}
