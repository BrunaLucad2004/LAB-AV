package main

import (
	"fmt"
	//"sync"
)

// cria uma constante pra limitar a capacidade do canal
const capacity = 100

// canal para verificar se algo foi feito
var done = make(chan bool)

// canal que recebe as "caixas" do deposito, com uma capacidade total de 100 itens
var deposit = make(chan int, capacity)

// cria a variavel para sincronizar
//var wg sync.WaitGroup

// funcao principal
func main() {
	go producer()

	//wg.Add(1)

	go consumer()

	//espera a função do consumidor terminar
	//wg.Wait()

	<-done
}

// funcao produtor
func producer() {
	for box := 0; box < 101; box++ {
		select {
		case deposit <- box:
			fmt.Println("Caixa enviada:", box)
		default:
			fmt.Println("Caixa não enviada, o canal cheio")
		}
	}
	done <- true
	//close(deposit)
}

// funcao consumidor
func consumer() {
	// Sinaliza que a goroutine terminou
	//defer wg.Done()

	for {
		select {
		case pack, ok := <-deposit:
			if !ok {
				fmt.Println("Canal está vazio, não há itens para retirar.")
				return
			}

			fmt.Println("Recebeu caixa:", pack)
		default:

		}
	}
}

//func clean()
//func fill()
