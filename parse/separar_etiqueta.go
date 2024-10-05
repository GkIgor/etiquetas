package parse

import "strings"

func SepararEtiqueta(linha string, etiqueta *Etiqueta, tipo int) {
	// PARAMOS AQUI!!!!!
	// TODO: PARA CADA LINHA, COMEÇAR A GERAR UMA NOVA ETIQUETA
	//TODO LEMBRAR DE FAZER O HASH DA ETIQUETA
	// TODO LEMBRAR QUE NO FINAL DA ETIQUETA DEVE LIMPAR O PONTEIRO DA ETIQUETA
	// NÃO ESQUECER DE ANEXAR O PONTEIRO DE ETIQUETA AO PONTEIRO DE ARQUIVO

	switch tipo {
	case TYPE_06:
		separarTipo06(linha, etiqueta)

	}
}

func separarTipo06(linha string, etiqueta *Etiqueta) {
	if etiqueta.Identificador == "" {
		etiqueta.Identificador = "sas06"
	}

	if strings.HasPrefix(linha, START) {
		etiqueta.body += linha + "\n"
	}

	if strings.HasPrefix(linha, POSITION_1) && strings.HasPrefix(etiqueta.body, START) {
		etiqueta.body += linha + "\n"
	}

	// TODO: PARAMOS AQUI!
	// LEMBRAR DE TERMINAR DE ANEXAR AS LINHAS NO CORPO DA ETIQUETA
	// LEMBRAR DE VERIFICAR O INICIO E FINAL DA ETIQUETA ANTES DE FAZER UM APPEND

}
