package parse

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
	Hash string
	body string
}

type Etiqueta struct {
	Identificador string
	Categoria     string
	Hash          string
	body          string
}
