package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "O valor do deposito menor que 0", c.saldo
	}

}

func (c *ContaCorrente) Tranferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 500}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 200}
	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDoGustavo.saldo)

	status := contaDaSilvia.Tranferir(300, &contaDoGustavo)

	fmt.Println(status)

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDoGustavo.saldo)
	contaDaSilvia.Depositar(500)
	fmt.Println(contaDaSilvia.saldo)
}
