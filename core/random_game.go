package core

import (
	"fmt"
	"math/rand"

	"golang.org/x/exp/slices"
)

func RandomNumber(max int, zeroInclude bool) int {
	if zeroInclude {
		return rand.Intn(max)
	}
	result := 0
	for result == 0 {
		result = rand.Intn(max)
	}
	return result
}

func TicketNumber(numbers int) string {
	randomNumber := ""
	for i := 0; i < numbers; i++ {
		rn := rand.Intn(10)

		randomNumber = fmt.Sprintf("%s%v", randomNumber, rn)
	}
	return randomNumber
}

func MelGenerator(turns int) []int {
	var mel []int

	for i := 0; i < turns; i++ {
		mel = []int{}

		for i := 0; i < 6; i++ {
			rn := RandomNumber(50, false)
			for slices.Contains(mel, rn) {
				// fmt.Printf("%d duplicated ", rn)
				rn = RandomNumber(50, false)
				// fmt.Printf("%d new \n", rn)
			}
			mel = append(mel, rn)
		}
		fmt.Println(mel)
	}

	return mel
}
