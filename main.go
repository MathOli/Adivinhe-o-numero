package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var (
		chute   int
		chances int
	)

	max := 100
	min := 1

	na := numeroAleatorio(max, min)

	dificuldade()
	chances = tentativas()

	fmt.Println("-------------------")

	restantes := chances

	for {

		str := fmt.Sprintf("Adivinhe o Numero(entre %d e %d)", min, max)
		fmt.Println(str)

		tenta := fmt.Sprintf("Você possue %d/%d", restantes, chances)
		fmt.Println(tenta)

		fmt.Println("Qual é seu chute")
		fmt.Scan(&chute)

		winOrLose(chute, na, restantes)

		time.Sleep(2 * time.Second)
		fmt.Println("-------------------")

		restantes -= 1
	}

}

func numeroAleatorio(max int, min int) int {

	rand.Seed(time.Now().UnixNano())
	dado := rand.Intn(100-1) + 1

	return dado
}

func dificuldade() {
	fmt.Println(" Escolha a dificuldade")
	fmt.Println(" 1 - Facil")
	fmt.Println(" 2 - Medio")
	fmt.Println(" 3 - Dificil")
}

func tentativas() int {
	fmt.Print("Escolha:")
	var dificuldade int
	fmt.Scan(&dificuldade)

	switch dificuldade {
	case 1:
		return 10
	case 2:
		return 5
	case 3:
		return 3
	default:
		fmt.Println("Dificuldade invalida")
		os.Exit(0)
	}

	return 0
}

func winOrLose(chute int, na int, restantes int) {
	if na == chute {
		fmt.Println("Parabens acertou")
		os.Exit(0)
	} else if restantes > 1 {
		if na > chute {
			fmt.Println("seu chute foi menor")
		} else {
			fmt.Println("seu chute foi maior")
		}
	} else {
		fmt.Println("Você perdeu todas as suas chances")
		os.Exit(0)
	}
}
