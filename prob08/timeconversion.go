package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeconversion(palavra string) []string {
	palavras := make([]string, 100)
	palavras2 := make([]string, 100)
	horario := make([]string, 5)
	palavras = strings.Split(palavra, ":")
	if len(palavras[2]) == 4 && palavras[2][2] == 80 {
		hora, _ := strconv.Atoi(palavras[0])
		if hora == 12 {
			palavras2 = strings.Split(palavras[2], "P")
			minuto := palavras[1]
			segundo := palavras2[0]
			horario[0] = strconv.Itoa(hora)
			horario[1] = ":"
			horario[2] = minuto
			horario[3] = ":"
			horario[4] = segundo
			// fmt.Printf("%d:%s:%s", hora, minuto, segundo)
			// fmt.Printf("\n")
		} else {
			palavras2 = strings.Split(palavras[2], "P")
			hora += 12
			minuto := palavras[1]
			segundo := palavras2[0]
			horario[0] = strconv.Itoa(hora)
			horario[1] = ":"
			horario[2] = minuto
			horario[3] = ":"
			horario[4] = segundo
			// fmt.Printf("%d:%s:%s", hora, minuto, segundo)
			// fmt.Printf("\n")
		}
	} else if len(palavras[2]) == 4 {
		horateste, _ := strconv.Atoi(palavras[0])
		if horateste == 12 {
			palavras2 = strings.Split(palavras[2], "A")
			horateste := "00"
			minuto := palavras[1]
			segundo := palavras2[0]
			horario[0] = horateste
			horario[1] = ":"
			horario[2] = minuto
			horario[3] = ":"
			horario[4] = segundo
			// fmt.Printf("%s:%s:%s", horateste, minuto, segundo)
			// fmt.Printf("\n")
		} else {
			palavras2 = strings.Split(palavras[2], "A")
			hora := palavras[0]
			minuto := palavras[1]
			segundo := palavras2[0]
			horario[0] = hora
			horario[1] = ":"
			horario[2] = minuto
			horario[3] = ":"
			horario[4] = segundo
			// fmt.Printf("%s:%s:%s", hora, minuto, segundo)
			// fmt.Printf("\n")
		}
	}
	return horario
}

func main() {
	var palavra string
	horario := make([]string, 5)
	fmt.Scanf("%s", &palavra)
	horario = timeconversion(palavra)
	fmt.Print(horario)
}
