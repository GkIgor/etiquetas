package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

const (
	TYPE_06  = 1
	TYPE_06C = 3
	TYPE_16  = 4
	TYPE_16C = 5

	START         = "^XA"
	POSITION_1    = "~DG000"
	PARTIAL_END_1 = ":::^XA"
	END           = "^XZ"
)

type Arquivo struct {
	Nome string
	hash string
}

type Etiqueta struct {
	Identificador string
	Categoria     string
	Hash          string
}

func main() {
	// ponteiro para slice de aquivos
	arquivos := new([]Arquivo)

	obterEtiquetas("./", "sas06", arquivos)
}

// etq name start With sas06/sas16 and sas06c/sas16c
func obterEtiquetas(path string, prefixo string, arquivos *[]Arquivo) {
	dir, erro := os.Open(path)

	if erro != nil {
		panic(erro)
	}

	defer dir.Close()
	nomeDasEtiquetas := make([]string, 0)
	entradas, erro := dir.Readdir(-1)

	if erro != nil {
		panic(erro)
	}

	for _, entrada := range entradas {
		if strings.HasPrefix(entrada.Name(), prefixo) {
			nomeDasEtiquetas = append(nomeDasEtiquetas, entrada.Name())
		}
	}

	etqParsed, erro := tentarFazerParseDaEtiqueta(path+nomeDasEtiquetas[0], arquivos)

	if erro != nil {
		etiqueta := new(Etiqueta)

		for _, linha := range etqParsed {
			gerarEtiqueta(linha, etiqueta)
		}
	}
}

func gerarSha256(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func tentarFazerParseDaEtiqueta(dirNome string, arquivos *[]Arquivo) ([]string, error) {
	arquivo, erro := os.Open(dirNome)

	if erro != nil {
		panic(erro)
	}

	defer arquivo.Close()

	*arquivos = append(*arquivos, Arquivo{hash: gerarSha256(dirNome), Nome: dirNome})

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

func gerarEtiqueta(linha string, etiquetas *Etiqueta) {
	// PARAMOS AQUI!!!!!
	// TODO: PARA CADA LINHA, COMEÇAR A GERAR UMA NOVA ETIQUETA
	//TODO LEMBRAR DE FAZER O HASH DA ETIQUETA
	// TODO LEMBRAR QUE NO FINAL DA ETIQUETA DEVE LIMPAR O PONTEIRO DA ETIQUETA
	// NÃO ESQUECER DE ANEXAR O PONTEIRO DE ETIQUETA AO PONTEIRO DE ARQUIVO
}
