package main

import (
	"br-atacadao.corp/etiquetas/parse"
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

func main() {
	// ponteiro para slice de aquivos
	arquivos := new([]parse.Arquivo)

	parse.ObterEtiquetas("./", "sas06", arquivos)
}
