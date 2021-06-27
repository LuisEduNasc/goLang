package formas

import (
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Triangle", func(t *testing.T) {
		ret := Triangle{10, 12}
		areaEsperada := float64(60)
		areaRecebida := ret.GetArea()

		if areaEsperada != areaRecebida {
			t.Fatalf("Área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Square", func(t *testing.T) {
		sq := Square{10}

		areaEsperada := float64(100)
		areaRecebida := sq.GetArea()

		if areaEsperada != areaRecebida {
			t.Fatalf("Área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
