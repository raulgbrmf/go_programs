// nao fiz teste para esse pois se eu transformasse em um metodo, para nao usar scanf eu teria que armazenar
//Tudo de uma vez em um vetor talvez. Isso ficaria igual ao metodo anterior...
package main

import "fmt"

func main() {
	var number int
	var numbers, sum int64
	numbers = 0
	sum = 0
	fmt.Scanf("%d", &number)
	for i := 0; i < number; i++ {
		fmt.Scanf("%d", &numbers)
		sum += numbers
	}
	fmt.Print(sum)
}
