package main

import (
	"br-atacadao.corp/etiquetas/parse"
)

func main() {
	// ponteiro para slice de aquivos
	arquivos := new([]parse.Arquivo)

	parse.ObterEtiquetas("./", "sas06", arquivos)
}
