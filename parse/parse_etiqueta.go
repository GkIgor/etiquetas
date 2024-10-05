package parse

import (
	"bufio"
	"fmt"
	"os"

	"br-atacadao.corp/etiquetas/util"
)

func TentarFazerParseDaEtiqueta(dirNome string, arquivos *[]Arquivo) ([]string, error) {
	arquivo, erro := os.Open(dirNome)

	if erro != nil {
		panic(erro)
	}

	defer arquivo.Close()

	*arquivos = append(*arquivos, Arquivo{Hash: util.GerarSha256(dirNome), Nome: dirNome})

	// Um token é uma etiqueta zpl completa!
	arquivoTokens := make([]string, 0)

	scanner := bufio.NewScanner(arquivo)

	for scanner.Scan() {
		line := scanner.Text()
		arquivoTokens = append(arquivoTokens, line)
	}

	if erro := scanner.Err(); erro != nil {
		descricaoDoErro := fmt.Sprintf("Erro ao ler o arquivo! File: %s \n Erro: %s", dirNome, erro)
		panic(descricaoDoErro)
	}

	return arquivoTokens, nil
}
