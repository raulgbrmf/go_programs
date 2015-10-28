package main

import "fmt"

func staircase(height int) []string {
	hashtags := make([]string, (height+1)*height)
	generalcounter := 0
	var floorup int
	var floordown int
	floorup = height - 1
	floordown = 0
	for i := 0; i < height; i++ {
		for j := 0; j < floorup; j++ {
			hashtags[generalcounter] = " "
			generalcounter++
			// fmt.Print(" ")
		}
		for k := 0; k < floordown; k++ {
			hashtags[generalcounter] = "#"
			generalcounter++
			// fmt.Print("#")
		}
		hashtags[generalcounter] = "#"
		// fmt.Println("#")
		generalcounter++
		hashtags[generalcounter] = "\n"
		generalcounter++
		floorup--
		floordown++
	}
	return hashtags
}

func main() {
	var height int
	fmt.Scanf("%d", &height)
	hashtags := staircase(height)
	for i := 0; i < ((height + 1) * height); i++ {
		fmt.Printf("%s", hashtags[i])
	}
}
