package main

import (
	"fmt"
	"runtime"
	"time"
)

func MultiplyRow(row []float64, scalar float64) {

	//As funcoes para tempo e ver as goroutines foram retiradas dos exemplos do professor

	fmt.Printf("M.Work started at %s\n", time.Now().Format("15:04:05")) //inicia a contagem do tempo
	fmt.Println("gorotines: ", runtime.NumGoroutine())                  //informa o numero de gorotines sendo executadas
	time.Sleep(1 * time.Second)                                         //tempo para cada processo

	//multplica o escalar por cada elemento da linha
	for i := 0; i < len(row); i++ {
		row[i] *= scalar
	}

	fmt.Printf("Work  finished at %s\n", time.Now().Format("15:04:05")) //finaliza a contagem do tempo

}

func main() {

	// Definir uma matriz 13x13
	matrix := [13][13]float64{
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
		{3, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 8},
	}

	// Escalar
	k := 28.03

	// Multiplicar a matriz pelo escalar
	for i := 0; i < 13; i++ {
		MultiplyRow(matrix[i][:], k)
	}

	// Exibir a matriz
	fmt.Println("Matrix:")
	for i := 0; i < 13; i++ {
		fmt.Println(matrix[i])
	}
}
