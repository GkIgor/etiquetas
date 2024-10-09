package parse

import (
	"fmt"
	"os"
	"strings"

	"br-atacadao.corp/etiquetas/util"
)

// etq name start With sas06/sas16 and sas06c/sas16c
func ObterEtiquetas(path string, prefixo string, arquivos *[]Arquivo) {
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

	etqParsed, erro := TentarFazerParseDaEtiqueta(path+nomeDasEtiquetas[0], arquivos)

	if erro != nil {
		etiqueta := new(Etiqueta)

		var tipo int

		switch prefixo {
		case "sas06":
			tipo = TYPE_06
		case "sas16":
			tipo = TYPE_16
		case "sas06c":
			tipo = TYPE_06C
		case "sas16c":
			tipo = TYPE_16C
		default:
			panic("Tipo de etiqueta desconhecido")
		}

		controlador := 0

		for index, linha := range etqParsed {
			if index == 0 && strings.HasPrefix(linha, START) {
				controlador = 1
				SepararEtiqueta(linha, etiqueta, tipo)
			} else {
				controlador = 0
			}

			if index == 1 && strings.HasPrefix(linha, POSITION_1) && controlador == 1 {
				controlador = 1
				SepararEtiqueta(linha, etiqueta, tipo)
			} else {
				controlador = 0
			}

			if controlador == 1 {
				SepararEtiqueta(linha, etiqueta, tipo)
			} else {
				panic("Etiqueta inválida")
			}

		}

		etiqueta.Hash = util.GerarSha256(path + nomeDasEtiquetas[0])

		*arquivos = append(*arquivos, Arquivo{Nome: path + nomeDasEtiquetas[0], Hash: etiqueta.Hash, body: etiqueta.body})

		// imprimir no terminal o arquivo
		fmt.Println(arquivos)
	}
}
