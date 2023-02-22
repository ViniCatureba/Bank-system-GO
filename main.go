package main

import (
	"bank/clientes"
	"bank/contas" // tmb pode dar um apelido como no python "c bank/contas"
	"fmt"
)

func main() {
	clienteBruno := contas.ContaCorrente{
		Titular:       clientes.Titular{},
		NumeroAgencia: 0,
		NumeroConta:   0,
	}
	clienteBruno.Depositar(500)
	fmt.Println(clienteBruno.ObterSaldo())

}
