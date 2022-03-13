package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"time"

	"markets/pkg/app/errors"
	"markets/pkg/app/interfaces"
	valueObjects "markets/pkg/domain/value_objects"
	"markets/pkg/infra/database/models"
)

type marketRepository struct {
	logger interfaces.ILogger
	db     *sql.DB
}

var now = time.Now

func (pst marketRepository) Create(ctx context.Context, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, error) {
	sql := `
		INSERT INTO feiras 
			(long, lat, setcens, areap, coddist, distrito, codsubpref, subpref, regiao5, regiao8, nome_feira, registro, logradouro, numero, 
				bairro, referencia, criado_em, atualizado_em)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
		RETURNING *
	`
	prepare, err := pst.db.PrepareContext(ctx, sql)
	if err != nil {
		pst.logger.Error("[MarketRepository::Create] Error in prepare statement")
		return valueObjects.MarketValueObjects{}, errors.NewInternalError("error in prepare statement")
	}

	row := prepare.QueryRowContext(ctx, market.Long, market.Lat, market.Setcens, market.Areap, market.Coddist, market.Distrito, market.Codsubpref,
		market.Subpref, market.Regiao5, market.Regiao8, market.NomeFeira, market.Registro, market.Logradouro, market.Numero, market.Bairro,
		market.Referencia, now(), now())
	if row.Err() != nil {
		pst.logger.Error("[MarketRepository::Create] query execution error")
		return valueObjects.MarketValueObjects{}, errors.NewInternalError("query execution error")
	}

	result, err := pst.scan(row)
	if err != nil {
		pst.logger.Error("[MarketRepository::Create] - scanning the result failure")
		return valueObjects.MarketValueObjects{}, err
	}

	return result, nil
}

func (pst marketRepository) Find(ctx context.Context, market valueObjects.MarketValueObjects) ([]valueObjects.MarketValueObjects, error) {
	sql := `SELECT
								id AS ID,
								long AS Long,
								lat AS Lat,
								setcens AS Setcens,
								areap AS Areap,
								coddist AS Coddist,
								distrito AS Distrito,
								codsubpref AS Codsubpref,
								subpref AS Subpref,
								regiao5 AS Regiao5,
								regiao8 AS Regiao8,
								nome_feira AS NomeFeira,
								registro AS Registro,
								logradouro AS Logradouro,
								numero AS Numero,
								bairro AS Bairro,
								referencia AS Referencia,
								criado_em AS CriadoEm,
								atualizado_em AS AtualizadoEm,
								deletado_em AS DeletadoEm
					FROM feiras
					WHERE deletado_em IS NULL`

	where, fields := buildQuery("AND", "", market)

	sql += where
	prepare, err := pst.db.PrepareContext(ctx, sql)
	if err != nil {
		pst.logger.Error("[MarketRepository::Find] Error in prepare statement")
		return nil, errors.NewInternalError("error in prepare statement")
	}

	rows, err := prepare.QueryContext(ctx, fields...)
	if err != nil {
		pst.logger.Error("[MarketRepository::Find] query execution error")
		return nil, errors.NewInternalError("query execution error")
	}

	var results []valueObjects.MarketValueObjects
	for rows.Next() {
		result, err := pst.scan(rows)
		if err != nil {
			pst.logger.Error("[MarketRepository::Find] - scanning the result failure")
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (pst marketRepository) Update(ctx context.Context, registerCode string, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, error) {
	sql := `UPDATE feiras  SET `

	set, fields := buildQuery("", ",", market)
	fields = append(fields, registerCode)
	set = set[:len(set)-1]
	set += fmt.Sprintf(" WHERE registro = $%v RETURNING feiras.*", len(fields))
	sql += set

	prepare, err := pst.db.PrepareContext(ctx, sql)
	if err != nil {
		pst.logger.Error("[MarketRepository::Update] Error in prepare statement")
		return valueObjects.MarketValueObjects{}, errors.NewInternalError("error in prepare statement")
	}

	row := prepare.QueryRowContext(ctx, fields...)
	if row.Err() != nil {
		pst.logger.Error("[MarketRepository::Update] query execution error")
		return valueObjects.MarketValueObjects{}, errors.NewInternalError("query execution error")
	}

	result, err := pst.scan(row)
	if err != nil {
		pst.logger.Error("[MarketRepository::Update] - scanning the result failure")
		return valueObjects.MarketValueObjects{}, err
	}

	return result, nil
}

func (pst marketRepository) Delete(ctx context.Context, registerCode string) error {
	sql := `UPDATE feiras SET deletado_em = $1 WHERE registro = $2`

	prepare, err := pst.db.PrepareContext(ctx, sql)
	if err != nil {
		pst.logger.Error("[MarketRepository::Delete] Error in prepare statement")
		return errors.NewInternalError("error in prepare statement")
	}

	_, err = prepare.QueryContext(ctx, now(), registerCode)
	if err != nil {
		pst.logger.Error("[MarketRepository::Delete] query execution error")
		return errors.NewInternalError("query execution error")
	}

	return nil
}

func buildQuery(pre, pos string, market valueObjects.MarketValueObjects) (string, []interface{}) {
	var mappingFields = map[string]string{
		"Long": "long", "Lat": "lat", "Setcens": "setcens", "Areap": "areap", "Coddist": "coddist", "Distrito": "distrito", "Codsubpref": "codsubpref",
		"Subpref": "subpref", "Regiao5": "regiao5", "Regiao8": "regiao8", "NomeFeira": "nome_feira", "Registro": "registro", "Logradouro": "logradouro",
		"Numero": "numero", "Bairro": "bairro", "Referencia": "referencia", "CriadoEm": "criado_em", "AtualizadoEm": "atualizado_em",
	}

	vOf := reflect.ValueOf(market)

	where := ""
	var field reflect.Value
	fields := make([]interface{}, 0)
	fieldCount := 1

	for i := 0; i < vOf.NumField(); i++ {
		field = vOf.Field(i)
		fieldName := mappingFields[vOf.Type().Field(i).Name]
		if !field.IsZero() {
			where += fmt.Sprintf(" %s %s = $%v%s", pre, fieldName, fieldCount, pos)
			fields = append(fields, field.Interface())
			fieldCount++
		}
	}

	return where, fields
}

type IRow interface {
	Scan(dest ...interface{}) error
}

func (pst marketRepository) scan(row IRow) (valueObjects.MarketValueObjects, error) {
	model := models.MarketModel{}
	if err := row.Scan(&model.ID, &model.Long, &model.Lat, &model.Setcens, &model.Areap, &model.Coddist, &model.Distrito, &model.Codsubpref,
		&model.Subpref, &model.Regiao5, &model.Regiao8, &model.NomeFeira, &model.Registro, &model.Logradouro, &model.Numero, &model.Bairro,
		&model.Referencia, &model.CriadoEm, &model.AtualizadoEm, &model.DeletadoEm); err != nil {
		return valueObjects.MarketValueObjects{}, errors.NewInternalError("error in scanning the results")
	}
	return model.ToValueObject(), nil
}

func NewMarketRepository(logger interfaces.ILogger, db *sql.DB) interfaces.IMarketRepository {
	return marketRepository{logger, db}
}
