package repositories

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ralvescosta/base/pkg/app/interfaces"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
	"github.com/ralvescosta/base/pkg/infra/database/models"
	"github.com/ralvescosta/base/pkg/infra/logger"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func Test_MarketRepo_Create(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		sut.sqlMockForCreateSuccessfully()

		result, err := sut.repo.Create(context.Background(), sut.marketMocked)

		assert.NoError(t, err)
		assert.Equal(t, sut.marketMocked, result)
	})

	t.Run("should return err when prepare statement failure", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		sut.logger.On("Error", "[MarketRepository::Create] Error in prepare statement", []zapcore.Field(nil))

		_, err := sut.repo.Create(context.Background(), sut.marketMocked)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err if query failure", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		prepare := sut.sqlMock.ExpectPrepare("")
		prepare.ExpectQuery().WithArgs()
		sut.logger.On("Error", "[MarketRepository::Create] query execution error", []zapcore.Field(nil))

		_, err := sut.repo.Create(context.Background(), sut.marketMocked)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err when scanning failure", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		prepare := sut.sqlMock.ExpectPrepare("")
		row := sut.sqlMock.NewRows([]string{""})
		prepare.ExpectQuery().WithArgs().WillReturnRows(row)
		sut.logger.On("Error", "[MarketRepository::Create] - scanning the result failure", []zapcore.Field(nil))

		_, err := sut.repo.Create(context.Background(), sut.marketMocked)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})
}

func Test_MarketRepo_Find(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		sut.sqlMockForFindSuccessfully()

		result, err := sut.repo.Find(context.Background(), valueObjects.MarketValueObjects{Long: sut.marketMocked.Long})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err when prepare statement failure", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		sut.logger.On("Error", "[MarketRepository::Find] Error in prepare statement", []zapcore.Field(nil))
		result, err := sut.repo.Find(context.Background(), valueObjects.MarketValueObjects{Long: sut.marketMocked.Long})

		assert.Error(t, err)
		assert.Nil(t, result)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err if query failure", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		prepare := sut.sqlMock.ExpectPrepare("")
		prepare.ExpectQuery().WithArgs()
		sut.logger.On("Error", "[MarketRepository::Find] query execution error", []zapcore.Field(nil))

		result, err := sut.repo.Find(context.Background(), valueObjects.MarketValueObjects{Long: sut.marketMocked.Long})

		assert.Error(t, err)
		assert.Nil(t, result)
		sut.logger.AssertExpectations(t)
	})
}

func Test_MarketRepo_Update(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		sut.sqlMockForUpdateSuccessfully()
		sut.marketMocked.Registro = ""
		sut.marketMocked.ID = 0

		_, err := sut.repo.Update(context.Background(), "registro", sut.marketMocked)

		assert.NoError(t, err)
	})

	t.Run("should return err when prepare statement failure", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		sut.logger.On("Error", "[MarketRepository::Update] Error in prepare statement", []zapcore.Field(nil))

		_, err := sut.repo.Update(context.Background(), "registro", sut.marketMocked)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err if query failure", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		prepare := sut.sqlMock.ExpectPrepare("")
		prepare.ExpectQuery().WithArgs()
		sut.logger.On("Error", "[MarketRepository::Update] query execution error", []zapcore.Field(nil))

		_, err := sut.repo.Update(context.Background(), "registro", sut.marketMocked)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err when scanning failure", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		prepare := sut.sqlMock.ExpectPrepare("")
		row := sut.sqlMock.NewRows([]string{""})
		prepare.ExpectQuery().WithArgs().WillReturnRows(row)
		sut.logger.On("Error", "[MarketRepository::Update] - scanning the result failure", []zapcore.Field(nil))

		_, err := sut.repo.Update(context.Background(), "registro", sut.marketMocked)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})
}

func Test_MarketRepo_Delete(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		sut.sqlMockForDeleteSuccessfully()

		err := sut.repo.Delete(context.Background(), sut.marketMocked.Registro)

		assert.NoError(t, err)
	})

	t.Run("should return err when prepare statement failure", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		sut.logger.On("Error", "[MarketRepository::Delete] Error in prepare statement", []zapcore.Field(nil))

		err := sut.repo.Delete(context.Background(), sut.marketMocked.Registro)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err if query failure", func(t *testing.T) {
		sut := makeMarketRepositorySut()

		prepare := sut.sqlMock.ExpectPrepare("")
		prepare.ExpectQuery().WithArgs()
		sut.logger.On("Error", "[MarketRepository::Delete] query execution error", []zapcore.Field(nil))

		err := sut.repo.Delete(context.Background(), sut.marketMocked.Registro)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})
}

type marketRepositorySutRtn struct {
	logger       *logger.LoggerSpy
	db           *sql.DB
	sqlMock      sqlmock.Sqlmock
	repo         interfaces.IMarketRepository
	marketMocked valueObjects.MarketValueObjects
	modelMocked  models.MarketModel
}

func (pst marketRepositorySutRtn) sqlMockForCreateSuccessfully() {
	query :=
		"INSERT INTO feiras \\(long, lat, setcens, areap, coddist, distrito, codsubpref, subpref, regiao5, regiao8, nome_feira, registro, logradouro, numero, bairro, referencia, criado_em, atualizado_em\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7, \\$8, \\$9, \\$10, \\$11, \\$12, \\$13, \\$14, \\$15, \\$16, \\$17, \\$18\\) RETURNING \\*"
	rows := pst.sqlMock.NewRows(
		[]string{"id", "long", "lat", "setcens", "areap", "coddist", "distrito", "codsubpref", "subpref", "regiao5", "regiao8", "nome_feira", "registro",
			"logradouro", "numero", "bairro", "referencia", "criado_em", "atualizado_em", "deletado_em"},
	).AddRow(
		pst.modelMocked.ID,
		pst.modelMocked.Long,
		pst.modelMocked.Lat,
		pst.modelMocked.Setcens,
		pst.modelMocked.Areap,
		pst.modelMocked.Coddist,
		pst.modelMocked.Distrito,
		pst.modelMocked.Codsubpref,
		pst.modelMocked.Subpref,
		pst.modelMocked.Regiao5,
		pst.modelMocked.Regiao8,
		pst.modelMocked.NomeFeira,
		pst.modelMocked.Registro,
		pst.modelMocked.Logradouro,
		pst.modelMocked.Numero,
		pst.modelMocked.Bairro,
		pst.modelMocked.Referencia,
		pst.modelMocked.CriadoEm,
		pst.modelMocked.AtualizadoEm,
		pst.modelMocked.DeletadoEm,
	)

	prepare := pst.sqlMock.ExpectPrepare(query)

	prepare.ExpectQuery().WithArgs(
		pst.modelMocked.Long,
		pst.modelMocked.Lat,
		pst.modelMocked.Setcens,
		pst.modelMocked.Areap,
		pst.modelMocked.Coddist,
		pst.modelMocked.Distrito,
		pst.modelMocked.Codsubpref,
		pst.modelMocked.Subpref,
		pst.modelMocked.Regiao5,
		pst.modelMocked.Regiao8,
		pst.modelMocked.NomeFeira,
		pst.modelMocked.Registro,
		pst.modelMocked.Logradouro,
		pst.modelMocked.Numero,
		pst.modelMocked.Bairro,
		pst.modelMocked.Referencia,
		pst.modelMocked.CriadoEm,
		pst.modelMocked.AtualizadoEm,
	).WillReturnRows(rows)
}

func (pst marketRepositorySutRtn) sqlMockForFindSuccessfully() {
	query := "SELECT id AS ID, long AS Long, lat AS Lat, setcens AS Setcens, areap AS Areap, coddist AS Coddist, distrito AS Distrito, codsubpref AS Codsubpref, subpref AS Subpref, regiao5 AS Regiao5, regiao8 AS Regiao8, nome_feira AS NomeFeira, registro AS Registro, logradouro AS Logradouro, numero AS Numero, bairro AS Bairro, referencia AS Referencia, criado_em AS CriadoEm, atualizado_em AS AtualizadoEm, deletado_em AS DeletadoEm FROM feiras WHERE deletado_em IS NULL AND long = \\$1"
	rows := pst.sqlMock.NewRows(
		[]string{"id", "long", "lat", "setcens", "areap", "coddist", "distrito", "codsubpref", "subpref", "regiao5", "regiao8", "nome_feira", "registro",
			"logradouro", "numero", "bairro", "referencia", "criado_em", "atualizado_em", "deletado_em"},
	).AddRow(
		pst.modelMocked.ID,
		pst.modelMocked.Long,
		pst.modelMocked.Lat,
		pst.modelMocked.Setcens,
		pst.modelMocked.Areap,
		pst.modelMocked.Coddist,
		pst.modelMocked.Distrito,
		pst.modelMocked.Codsubpref,
		pst.modelMocked.Subpref,
		pst.modelMocked.Regiao5,
		pst.modelMocked.Regiao8,
		pst.modelMocked.NomeFeira,
		pst.modelMocked.Registro,
		pst.modelMocked.Logradouro,
		pst.modelMocked.Numero,
		pst.modelMocked.Bairro,
		pst.modelMocked.Referencia,
		pst.modelMocked.CriadoEm,
		pst.modelMocked.AtualizadoEm,
		pst.modelMocked.DeletadoEm,
	)

	prepare := pst.sqlMock.ExpectPrepare(query)

	prepare.ExpectQuery().WithArgs(
		pst.modelMocked.Long,
	).WillReturnRows(rows)
}

func (pst marketRepositorySutRtn) sqlMockForUpdateSuccessfully() {
	query :=
		"UPDATE feiras  SET   long = \\$1,  lat = \\$2,  setcens = \\$3,  areap = \\$4,  coddist = \\$5,  distrito = \\$6,  codsubpref = \\$7,  subpref = \\$8,  regiao5 = \\$9,  regiao8 = \\$10,  nome_feira = \\$11,  logradouro = \\$12,  numero = \\$13,  bairro = \\$14,  referencia = \\$15 WHERE registro = \\$16 RETURNING feiras.\\*"
	rows := pst.sqlMock.NewRows(
		[]string{"id", "long", "lat", "setcens", "areap", "coddist", "distrito", "codsubpref", "subpref", "regiao5", "regiao8", "nome_feira", "registro",
			"logradouro", "numero", "bairro", "referencia", "criado_em", "atualizado_em", "deletado_em"},
	).AddRow(
		pst.modelMocked.ID,
		pst.modelMocked.Long,
		pst.modelMocked.Lat,
		pst.modelMocked.Setcens,
		pst.modelMocked.Areap,
		pst.modelMocked.Coddist,
		pst.modelMocked.Distrito,
		pst.modelMocked.Codsubpref,
		pst.modelMocked.Subpref,
		pst.modelMocked.Regiao5,
		pst.modelMocked.Regiao8,
		pst.modelMocked.NomeFeira,
		pst.modelMocked.Registro,
		pst.modelMocked.Logradouro,
		pst.modelMocked.Numero,
		pst.modelMocked.Bairro,
		pst.modelMocked.Referencia,
		pst.modelMocked.CriadoEm,
		pst.modelMocked.AtualizadoEm,
		pst.modelMocked.DeletadoEm,
	)

	prepare := pst.sqlMock.ExpectPrepare(query)

	prepare.ExpectQuery().WithArgs(
		pst.modelMocked.Long,
		pst.modelMocked.Lat,
		pst.modelMocked.Setcens,
		pst.modelMocked.Areap,
		pst.modelMocked.Coddist,
		pst.modelMocked.Distrito,
		pst.modelMocked.Codsubpref,
		pst.modelMocked.Subpref,
		pst.modelMocked.Regiao5,
		pst.modelMocked.Regiao8,
		pst.modelMocked.NomeFeira,
		pst.modelMocked.Logradouro,
		pst.modelMocked.Numero,
		pst.modelMocked.Bairro,
		pst.modelMocked.Referencia,
		pst.modelMocked.Registro,
	).WillReturnRows(rows)
}

func (pst marketRepositorySutRtn) sqlMockForDeleteSuccessfully() {
	query := "UPDATE feiras SET deletado_em = \\$1 WHERE registro = \\$2"
	rows := pst.sqlMock.NewRows([]string{})

	prepare := pst.sqlMock.ExpectPrepare(query)

	prepare.ExpectQuery().WithArgs(
		pst.modelMocked.CriadoEm,
		pst.modelMocked.Registro,
	).WillReturnRows(rows)
}

func makeMarketRepositorySut() marketRepositorySutRtn {
	logger := logger.NewLoggerSpy()
	db, mock, _ := sqlmock.New()
	repo := NewMarketRepository(logger, db)

	marketMocked := valueObjects.MarketValueObjects{
		ID:         1,
		Long:       -100,
		Lat:        -100,
		Setcens:    "setcens",
		Areap:      "areap",
		Coddist:    10,
		Distrito:   "distrito",
		Codsubpref: 10,
		Subpref:    "subpref",
		Regiao5:    "regiao5",
		Regiao8:    "regiao8",
		NomeFeira:  "nomefeira",
		Registro:   "registro",
		Logradouro: "logradouro",
		Numero:     "numero",
		Bairro:     "bairro",
		Referencia: "referencia",
	}

	t := time.Now()
	now = func() time.Time { return t }

	modelMocked := models.MarketModel{
		ID:           1,
		Long:         -100,
		Lat:          -100,
		Setcens:      "setcens",
		Areap:        "areap",
		Coddist:      10,
		Distrito:     "distrito",
		Codsubpref:   10,
		Subpref:      "subpref",
		Regiao5:      "regiao5",
		Regiao8:      "regiao8",
		NomeFeira:    "nomefeira",
		Registro:     "registro",
		Logradouro:   "logradouro",
		Numero:       "numero",
		Bairro:       "bairro",
		Referencia:   "referencia",
		CriadoEm:     t,
		AtualizadoEm: t,
		DeletadoEm:   nil,
	}
	return marketRepositorySutRtn{logger, db, mock, repo, marketMocked, modelMocked}
}
