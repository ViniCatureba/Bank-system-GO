package main

import (
	"bank/contas" // tmb pode dar um apelido como no python "c bank/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteBruno := contas.ContaCorrente{}
	clienteBruno.Depositar(500)
	PagarBoleto(&clienteBruno, 60)
	fmt.Println(clienteBruno.ObterSaldo())

}
