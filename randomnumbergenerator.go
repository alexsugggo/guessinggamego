package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	for count := 0; count < 10; count++ {
		year := rand.Intn(2022) + 2000
		leapyear := year%4 == 0 || (year == 0 && year%100 != 0)
		month := rand.Intn(12) + 1
		daysInMonth := 31
		switch month {
		case 2:
			daysInMonth = 28
			if leapyear {
				daysInMonth = 29
			}

		case 4, 6, 9, 11:
			daysInMonth = 30
		}
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}


