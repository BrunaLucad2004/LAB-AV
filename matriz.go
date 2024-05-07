package main

import (
	"fmt"
	"runtime"
	"time"
)

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
	k := 2.00

	// Multiplicar a matriz pelo escalar
	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {

			fmt.Printf("M.Work started at %s\n", time.Now().Format("15:04:05"))
			fmt.Println("gorotines: ", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)

			matrix[i][j] *= k

			fmt.Printf("Work  finished at %s\n", time.Now().Format("15:04:05"))
		}
	}

	// Exibir a matriz
	fmt.Println("\n", "   Matriz resultante:")
	for i := 0; i < 13; i++ {
		fmt.Println(matrix[i])
	}
}
