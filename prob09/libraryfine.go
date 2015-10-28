package main

import "fmt"

func libraryfine(yearRecv int, monthRecv int, dayRecv int, yearShould int, monthShould int, dayShould int) int {
	var hackos int
	if yearShould > yearRecv {
		hackos = 0
	} else if yearRecv > yearShould {
		hackos = 10000
	} else if monthShould > monthRecv {
		hackos = 0
	} else if monthRecv > monthShould {
		hackos = 500 * (monthRecv - monthShould)
	} else if dayRecv > dayShould {
		hackos = 15 * (dayRecv - dayShould)
	} else {
		hackos = 0
	}
	return hackos
}

func main() {
	var yearRecv int
	var monthRecv int
	var dayRecv int
	var yearShould int
	var monthShould int
	var dayShould int
	var hackos int
	fmt.Scanf("%d", &dayRecv)
	fmt.Scanf("%d", &monthRecv)
	fmt.Scanf("%d", &yearRecv)
	fmt.Scanf("%d", &dayShould)
	fmt.Scanf("%d", &monthShould)
	fmt.Scanf("%d", &yearShould)
	hackos = libraryfine(yearRecv, monthRecv, dayRecv, yearShould, monthShould, dayShould)
	fmt.Println(hackos)
}
