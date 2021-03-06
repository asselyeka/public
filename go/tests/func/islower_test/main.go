package main

import (
	student "student"

	"lib"
	"lib/is"
)

func main() {
	// 15 unvalid strings in the table
	table := lib.MultRandASCII()

	// 15 valid lowercase strings of random size between 1 and 20 letters in the table
	for i := 0; i < 15; i++ {
		size := lib.RandIntBetween(1, 20)
		randLow := lib.RandLower()
		if len(randLow) <= size {
			table = append(table, randLow)
		} else {
			table = append(table, randLow[:size])
		}
	}

	// Special cases added to table
	table = append(table,
		"Hello! How are you?",
		"a",
		"z",
		"!",
		"HelloHowareyou",
		"What's this 4?",
		"Whatsthis4",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!",
		"0123456789",
		"01,02,03",
		"abcdefghijklmnopqrstuvwxyz",
		"hello",
		"hello!",
	)
	for _, arg := range table {
		lib.Challenge("IsLower", student.IsLower, is.Lower, arg)
	}
}
