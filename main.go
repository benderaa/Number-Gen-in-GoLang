// So up here ^^^^^^^ are your prefixes

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Global var and functions

var digits = []rune("0123456789")

func randSeq(mobile int) string {

	b := make([]rune, mobile)

	for i := range b {

		b[i] = digits[rand.Intn(len(digits))]
	}

	return string(b)
}

// Main function

func main() {

	rand.Seed(time.Now().UnixNano())
	// Choose how many numbers you want eg 6000
	for mob := 1; mob < 600; mob++ {

		var code = ("64")

		// enter prefix you wanna generate here
		var prefix = ("27")

		var mob = (code + prefix + randSeq(7))

		time.Sleep(500 * time.Millisecond)

		fmt.Println(mob)

	}

}
