package viewmodels

import (
	"testing"

	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"

	"github.com/stretchr/testify/assert"
)

func Test_MarketViewModel_ToValueObject(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := MarketViewModel{
			Long:     -200,
			Lat:      -500,
			Registro: "registro",
		}

		vo := sut.ToValueObject()

		assert.Equal(t, sut.Long, vo.Long)
		assert.Equal(t, sut.Lat, vo.Lat)
		assert.Equal(t, sut.Registro, vo.Registro)
	})
}

func Test_NewMarketViewModel(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		vo := valueObjects.MarketValueObjects{
			Long:     -200,
			Lat:      -500,
			Registro: "registro",
		}

		sut := NewMarketViewModel(vo)

		assert.Equal(t, vo.Long, sut.Long)
		assert.Equal(t, vo.Lat, sut.Lat)
		assert.Equal(t, vo.Registro, sut.Registro)
	})
}

func Test_NewSliceOfMarketViewModel(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		vo := []valueObjects.MarketValueObjects{
			{
				Long:     -200,
				Lat:      -500,
				Registro: "registro",
			},
		}

		sut := NewSliceOfMarketViewModel(vo)

		assert.Equal(t, len(sut), len(vo))
	})

	t.Run("should return a empty viewmodel if receive empty valueObject", func(t *testing.T) {
		sut := NewSliceOfMarketViewModel(nil)

		assert.Equal(t, len(sut), 0)
	})
}
