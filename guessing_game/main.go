package guessing_game

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func Main() {
	fmt.Println("Jogo da adivinhaçao")
	fmt.Println("um numero aleatorio sera sorteado. tente acertar o numero inteiro de 0 a 100")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("qual e o seu chute? ")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("o seu chute nao e um numero inteiro valido")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("voce errou, o numero sorteado e maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("voce errou, o numero sorteado e menor que", chuteInt)
		case chuteInt == x:
			fmt.Printf(
				"voce Acertou! o numero sorteado foi %d\n"+
					"Voce acertou em %d tentativas."+
					"\nessas foram as suas tentativas: %v",
				x, i+1, chutes[:i],
			)
			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf(
		"voce perdeu! o numero sorteado foi %d\n"+
			"Voce teve 10 tentativas."+
			"\nessas foram as suas tentativas: %v",
		x, chutes,
	)

}
