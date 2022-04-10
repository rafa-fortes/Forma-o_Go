package main

import "fmt"

func main() {
	fmt.Println("Olá mundo Go!")

	n1 := 5
	n2 := 3

	fmt.Println(soma2numeros(n1, n2))

	qualEhOMaiorNumero(n1, n2)

	numerosDivisiveisPor5Entre0e50()

}

func soma2numeros(n1, n2 int) int {
	return n1 + n2
}

// E se precisarmos comparar ambos os números, imprimindo quem é o maior?
// Teremos que comparar os números. Para isso, podemos usar instruções if/else e fazer as comparações:

func qualEhOMaiorNumero(n1, n2 int) {
	var mensagem string
	if n1 < n2 {
		mensagem = "O número 1 é menor que o número 2"
	} else if n1 == n2 {
		mensagem = "Os numeros 1 e 2 são iguais"
	} else {
		mensagem = "O número 1 é maior que o número 2"
	}
	fmt.Println(mensagem)
}

//Mas, e se precisarmos verificar todos os números divisíveis por 5 entre 0 e 50, por exemplo?
//Para isso, podemos usar um looping, juntamente com o if:

func numerosDivisiveisPor5Entre0e50() {
	for i := 0; i <= 50; i++ {
		if i%5 == 0 {
			fmt.Println(i, " é um número divisível por 5")
		}
	}
}
