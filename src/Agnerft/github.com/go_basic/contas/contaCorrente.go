package contas

import (
	"github.com/go_basic/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Cliente
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	//contaValida := valorConta.numeroConta == c.numeroConta

	if valor > 0 {
		c.saldo += valor
		return "saldo acrescido com sucesso, seu saldo atual é = ", c.saldo
	} else {
		//teste := strconv.FormatInt(int64(c.numeroConta), 10)
		return "Valor informado não é válido", valor
	}
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque <= c.saldo && valorDoSaque >= 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com Sucesso"
	} else {
		return "Operação negada, saldo insuficiente ou valor inadequado"
	}

}

func (c *ContaCorrente) Transferir(contaParaTransferir *ContaCorrente, valor float64) bool {

	if contaParaTransferir.saldo <= c.saldo && valor > 0 {
		c.saldo -= valor
		contaParaTransferir.Depositar(valor)

		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Obtersaldo() float64 {
	saldoObtido := c.saldo

	return saldoObtido
}
