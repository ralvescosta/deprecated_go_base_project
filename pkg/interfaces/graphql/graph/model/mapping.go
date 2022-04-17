package model

import (
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
)

func CreateMarketToValueObject(c CreateMarket) valueObjects.MarketValueObjects {
	return valueObjects.MarketValueObjects{
		Long:       c.Long,
		Lat:        c.Lat,
		Setcens:    c.Setcens,
		Areap:      c.Areap,
		Coddist:    c.Coddist,
		Distrito:   c.Distrito,
		Codsubpref: c.Codsubpref,
		Subpref:    c.Subpref,
		Regiao5:    c.Regiao5,
		Regiao8:    c.Regiao8,
		NomeFeira:  c.NomeFeira,
		Registro:   c.Registro,
		Logradouro: c.Logradouro,
		Numero:     c.Numero,
		Bairro:     c.Bairro,
		Referencia: c.Referencia,
	}
}

func MarketFilterToValueObject(f MarketFilters) valueObjects.MarketValueObjects {
	return valueObjects.MarketValueObjects{
		Long:       safeInt(f.Long),
		Lat:        safeInt(f.Lat),
		Setcens:    safeString(f.Setcens),
		Areap:      safeString(f.Areap),
		Coddist:    safeInt(f.Coddist),
		Distrito:   safeString(f.Distrito),
		Codsubpref: safeInt(f.Codsubpref),
		Subpref:    safeString(f.Subpref),
		Regiao5:    safeString(f.Regiao5),
		Regiao8:    safeString(f.Regiao8),
		NomeFeira:  safeString(f.NomeFeira),
		Registro:   safeString(f.Registro),
		Logradouro: safeString(f.Logradouro),
		Numero:     safeString(f.Numero),
		Bairro:     safeString(f.Bairro),
		Referencia: safeString(f.Referencia),
	}
}

func UpdateMarketToValueObject(c MarketToUpdate) valueObjects.MarketValueObjects {
	return valueObjects.MarketValueObjects{
		Long:       safeInt(c.Long),
		Lat:        safeInt(c.Lat),
		Setcens:    safeString(c.Setcens),
		Areap:      safeString(c.Areap),
		Coddist:    safeInt(c.Coddist),
		Distrito:   safeString(c.Distrito),
		Codsubpref: safeInt(c.Codsubpref),
		Subpref:    safeString(c.Subpref),
		Regiao5:    safeString(c.Regiao5),
		Regiao8:    safeString(c.Regiao8),
		NomeFeira:  safeString(c.NomeFeira),
		Registro:   c.Registro,
		Logradouro: safeString(c.Logradouro),
		Numero:     safeString(c.Numero),
		Bairro:     safeString(c.Bairro),
		Referencia: safeString(c.Referencia),
	}
}

func safeInt(i *int) int {
	if i == nil {
		return 0
	}

	return *i
}

func safeString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func ValueObjectToMarket(vo valueObjects.MarketValueObjects) *Market {
	return &Market{
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

func ValueObjectSliceToMarketSlice(voSlice []valueObjects.MarketValueObjects) []*Market {
	slice := []*Market{}

	for _, v := range voSlice {
		slice = append(slice,
			&Market{
				Long:       v.Long,
				Lat:        v.Lat,
				Setcens:    v.Setcens,
				Areap:      v.Areap,
				Coddist:    v.Coddist,
				Distrito:   v.Distrito,
				Codsubpref: v.Codsubpref,
				Subpref:    v.Subpref,
				Regiao5:    v.Regiao5,
				Regiao8:    v.Regiao8,
				NomeFeira:  v.NomeFeira,
				Registro:   v.Registro,
				Logradouro: v.Logradouro,
				Numero:     v.Numero,
				Bairro:     v.Bairro,
				Referencia: v.Referencia,
			})
	}

	return slice
}
