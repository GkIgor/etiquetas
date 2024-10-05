package parse

type Arquivo struct {
	Nome string
	Hash string
}

type Etiqueta struct {
	Identificador string
	Categoria     string
	Hash          string
	body          string
}
