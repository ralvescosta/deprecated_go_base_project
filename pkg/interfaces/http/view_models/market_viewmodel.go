package viewmodels

import valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"

type MarketViewModel struct {
	Long       int    `json:"long" validate:"required"`
	Lat        int    `json:"lat" validate:"required"`
	Setcens    string `json:"setcens" validate:"required"`
	Areap      string `json:"areap" validate:"required"`
	Coddist    int    `json:"coddist" validate:"required"`
	Distrito   string `json:"distrito" validate:"required"`
	Codsubpref int    `json:"codsubpref" validate:"required"`
	Subpref    string `json:"subpref" validate:"required"`
	Regiao5    string `json:"regiao5" validate:"required"`
	Regiao8    string `json:"regiao8" validate:"required"`
	NomeFeira  string `json:"nome_feira" validate:"required"`
	Registro   string `json:"registro" validate:"required"`
	Logradouro string `json:"logradouro" validate:"required"`
	Numero     string `json:"numero" validate:"required"`
	Bairro     string `json:"bairro" validate:"required"`
	Referencia string `json:"referencia" validate:"required"`
}

func (pst MarketViewModel) ToValueObject() valueObjects.MarketValueObjects {
	return valueObjects.MarketValueObjects{
		Long:       pst.Long,
		Lat:        pst.Lat,
		Setcens:    pst.Setcens,
		Areap:      pst.Areap,
		Coddist:    pst.Coddist,
		Distrito:   pst.Distrito,
		Codsubpref: pst.Codsubpref,
		Subpref:    pst.Subpref,
		Regiao5:    pst.Regiao5,
		Regiao8:    pst.Regiao8,
		NomeFeira:  pst.NomeFeira,
		Registro:   pst.Registro,
		Logradouro: pst.Logradouro,
		Numero:     pst.Numero,
		Bairro:     pst.Bairro,
		Referencia: pst.Referencia,
	}
}

func NewSliceOfMarketViewModel(vo []valueObjects.MarketValueObjects) []MarketViewModel {
	if len(vo) == 0 {
		return []MarketViewModel{}
	}

	var result []MarketViewModel
	for _, v := range vo {
		result = append(result, NewMarketViewModel(v))
	}

	return result
}

func NewMarketViewModel(vo valueObjects.MarketValueObjects) MarketViewModel {
	return MarketViewModel{
		Long:       vo.Long,
		Lat:        vo.Lat,
		Setcens:    vo.Setcens,
		Areap:      vo.Areap,
		Coddist:    vo.Coddist,
		Distrito:   vo.Distrito,
		Codsubpref: vo.Codsubpref,
		Subpref:    vo.Subpref,
		Regiao5:    vo.Regiao5,
		Regiao8:    vo.Regiao8,
		NomeFeira:  vo.NomeFeira,
		Registro:   vo.Registro,
		Logradouro: vo.Logradouro,
		Numero:     vo.Numero,
		Bairro:     vo.Bairro,
		Referencia: vo.Referencia,
	}
}
