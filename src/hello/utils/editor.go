package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Editor(resultado string, numeroLinha int, novoValor string) (string, error) {

	//resultado := "C:\\Users\\USER\\Microsip\\Microsip.txt"

	file, err := os.OpenFile(resultado, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
		return err.Error(), err
	}
	defer file.Close()

	// Lê o conteúdo do arquivo
	conteudo, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler o conteúdo do arquivo: %v", err)
		return err.Error(), err
	}

	//numeroLinha := 2
	//novoValor := "label=7848"

	// Converte o conteúdo para uma string
	conteudoArquivo := string(conteudo)

	linhas := strings.Split(conteudoArquivo, "\n")

	if numeroLinha > 0 && numeroLinha < len(linhas) {
		linhas[numeroLinha-1] = novoValor
	}

	novoConteudoArquivo := strings.Join(linhas, "\n")

	_, err = file.Seek(0, 0)
	if err != nil {
		//log.Fatal("Erro ao mover o ponteiro: %v", err)
		return err.Error(), err
	}

	_, err = file.WriteString(novoConteudoArquivo)
	if err != nil {
		//log.Fatal("Erro ao salvar novo conteudo: %v", err)
		return err.Error(), err
	}

	err = file.Truncate(int64(len(novoConteudoArquivo)))
	if err != nil {
		//log.Fatal("Erro ao truncar: %v", err)
		return err.Error(), err
	}

	//fmt.Println(novoConteudoArquivo)
	//fmt.Println("Alterações salvas com sucesso.")

	return "Alterações salvas com sucesso.", nil

}
