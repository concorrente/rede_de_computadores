// Pedro Chem & Rafael Almeida

package main

import (
	"fmt"
	"math/rand"
)

const N = 20

var dest = make(chan int, 100)
var rem = make(chan int, 100)
var msg = make(chan string, 100)

func main() {

	fmt.Println("------ Rede de Computadores -------")
	fmt.Println("------	 Inicio do Programa	 -------")
	var token [N + 1]chan bool

	for i := 0; i <= N; i++ { // aloca canais
		token[i] = make(chan bool, 100)
	}

	for i := 0; i < N; i++ {
		go processo(i, token[i], token[i+1], token[0])
	}

	token[0] <- true

	fmt.Scanln()
}

// =====================================================================
func processo(i int, in chan bool, out chan bool, first chan bool) {

	for {

		hasMsg := false
		valor1 := rand.Intn(20)
		valor2 := rand.Intn(20)
		fmt.Println("|", i, "| - Valor1: ", valor1, " Valor2: ", valor2)
		if valor1 == valor2 {
			fmt.Println("|", i, "| - VALORES IGUAIS")
			hasMsg = true
		}
		<-in
		fmt.Println("|", i, "| - TOKEN")

		if hasMsg {
			send(i, valor1, "Teste")
		}
		if i == N-1 {
			first <- true
		} else {
			out <- true
		}

	}
}

func send(i int, destino int, mensagem string) {
	if i == destino {
		fmt.Println("Sou o computador ", i, " e recebi a mensagem |", mensagem, "|")
		return
	} else if i == N-1 {
		send(0, destino, mensagem)
	} else {
		send(i+1, destino, mensagem)
	}
}
