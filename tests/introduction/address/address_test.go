package address

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{
			"Rua ABC",
			"Rua",
		},
		{
			"Avenida Paulista",
			"Avenida",
		},
		{
			"Rodovia dos imigrantes",
			"Rodovia",
		},
		{
			"Praça das rosas",
			"Tipo Inválido",
		},
		{
			"Estrada 123",
			"Estrada",
		},
		{
			"RUA DOS BOBOS",
			"Rua",
		},
		{
			"",
			"Tipo Inválido",
		},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
