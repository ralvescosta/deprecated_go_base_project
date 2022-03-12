package models

import (
	"time"

	valueObjects "markets/pkg/domain/value_objects"
)

type MarketModel struct {
	ID           int
	Long         int
	Lat          int
	Setcens      string
	Areap        string
	Coddist      int
	Distrito     string
	Codsubpref   int
	Subpref      string
	Regiao5      string
	Regiao8      string
	NomeFeira    string
	Registro     string
	Logradouro   string
	Numero       string
	Bairro       string
	Referencia   string
	CriadoEm     time.Time
	AtualizadoEm time.Time
	DeletadoEm   *time.Time
}

func (pst MarketModel) ToValueObject() valueObjects.MarketValueObjects {
	return valueObjects.MarketValueObjects{
		ID:         pst.ID,
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
