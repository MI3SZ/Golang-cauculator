package main

import "fmt"

var (
	a     float64
	b     float64
	total float64
	op    string
)

// Pergunta os numeros pra operacao
func pedirnumero() {
	fmt.Println("digite um numero:")
	_, _ = fmt.Scanln(&a)
	fmt.Println("digite outro numero:")
	_, _ = fmt.Scanln(&b)
}

// Pergunta a operacao
func pedirOperacao() {
	fmt.Println("digite uma operacao (+ - * /):")
	_, _ = fmt.Scanln(&op)
}

// Funcoes pra cada operacao
func soma() {
	total = a + b
}
func subtracao() {
	total = a - b
}
func multiplicacao() {
	total = a * b
}
func divisao() {
	if b == 0 {
		fmt.Println("divisao com zero nao existe")
		return
	}

	total = a / b
}

// Funcao pra detectar a operacao e chamar a funcao certa
func calc() {
	switch op {
	case "+":
		soma()
	case "-":
		subtracao()
	case "*":
		multiplicacao()
	case "/":
		divisao()
	}
	fmt.Println(a, op, b, "=", total)
}

// Chama o pedir numero, operacao, caucula e finaliza a execucao
func main() {
	pedirnumero()
	pedirOperacao()
	calc()
	angola()
}

// Agradece pela prefencia e encerra o programa
func angola() {
	fmt.Println("Fim do programa, obrigado pela preferencia !")
}
