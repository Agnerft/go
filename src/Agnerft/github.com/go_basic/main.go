package main

import (
	"fmt"

	"github.com/go_basic/clientes"
	"github.com/go_basic/contas"
)

func PagarBoleto(conta verificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {

	clienteCleito := clientes.Cliente{Nome: "Cleito", CPF: "02900274036", Profissao: "Barmen"}
	contaDoCleito := contas.ContaPoupanca{Titular: clienteCleito, NumeroAgencia: 001, NumeroConta: 123456, Operação: 002}

	contaDoCleito.Depositar(500)

	PagarBoleto(&contaDoCleito, 420)

	fmt.Println(contaDoCleito.Obtersaldo())
}
