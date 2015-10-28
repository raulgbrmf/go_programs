package main

import "fmt"
import "math/big"

func fatorial(N int64) *big.Int {
	result := big.NewInt(1)
	for N > 0 {
		result.Mul(result, big.NewInt(N))
		N--
	}
	return result
}

func main() {
	var N int64
	result := big.NewInt(0)
	fmt.Scanf("%d", &N)
	result = fatorial(N)
	fmt.Print(result)
}
