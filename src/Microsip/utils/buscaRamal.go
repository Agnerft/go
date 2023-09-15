package utils

import (
	"github.com.br/agnerft/models"
)

//var clientConfig []models.ClienteConfig

func BuscaRamal(ramal int, c []models.ClienteConfig) string {
	
	for i := range c[0].QuantRamais[0].Ramal {
		// 	ramalDesejado, _ := strconv.Atoi(clienteConfig[0].Ramal)

		// 	if clienteConfig[0].QuantRamaisOpen[i].Ramal == ramalDesejado {
		// 		clienteConfig[0].QuantRamaisOpen[i].INUSE = true

		// 		break // Parar o loop após encontrar o ramal desejado
	}

	return "TESTE"
}
