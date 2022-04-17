package model

import (
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
)

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
