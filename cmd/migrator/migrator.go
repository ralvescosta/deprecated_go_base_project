package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ralvescosta/base/pkg/app/interfaces"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
	"github.com/ralvescosta/base/pkg/infra/database"
	"github.com/ralvescosta/base/pkg/infra/environments"
	"github.com/ralvescosta/base/pkg/infra/logger"
	"github.com/ralvescosta/base/pkg/infra/repositories"
)

func readCsvFile(logger interfaces.ILogger, filePath string) []valueObjects.MarketValueObjects {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	var records []valueObjects.MarketValueObjects
	bypassFirstLine := 0
	for {
		rec, err := csvReader.Read()
		if bypassFirstLine == 0 {
			bypassFirstLine++
			continue
		}
		if err != nil && err != io.EOF {
			logger.Error(fmt.Sprintf("csv line unformatted - %s", err.Error()))
			continue
		}
		if err != nil && err == io.EOF {
			break
		}

		id, _ := strconv.Atoi(rec[0])
		long, _ := strconv.Atoi(rec[1])
		lat, _ := strconv.Atoi(rec[2])
		coddist, _ := strconv.Atoi(rec[5])
		codsubpref, _ := strconv.Atoi(rec[7])

		vo := valueObjects.MarketValueObjects{
			ID:         id,
			Long:       long,
			Lat:        lat,
			Setcens:    rec[3],
			Areap:      rec[4],
			Coddist:    coddist,
			Distrito:   rec[6],
			Codsubpref: codsubpref,
			Subpref:    rec[8],
			Regiao5:    rec[9],
			Regiao8:    rec[10],
			NomeFeira:  rec[11],
			Registro:   rec[12],
			Logradouro: rec[13],
			Numero:     rec[14],
			Bairro:     rec[15],
			Referencia: rec[16],
		}
		records = append(records, vo)
	}

	return records
}

func exec() {
	if err := environments.NewEnvironment().Configure(); err != nil {
		log.Fatal(err)
	}
	os.Setenv("LOG_FILE", "./logs/seeder.log")

	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("[Seeder] - Starting...")
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fileDir := currentDir + "/cmd/seeders/DEINFO_AB_FEIRASLIVRES_2014.csv"

	logger.Info("[Seeder] - Reading the CSV file...")
	records := readCsvFile(logger, fileDir)
	logger.Info("[Seeder] - CSV File read")

	logger.Info("[Seeder] - Connection to the database...")
	db, err := database.Connect(logger, make(chan bool))
	if err != nil {
		log.Fatal(err)
	}
	marketRepository := repositories.NewMarketRepository(logger, db)
	logger.Info("[Seeder] - Database connected")

	row := db.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM feiras")
	var amount int
	row.Scan(&amount)
	if amount >= len(records) {
		logger.Info("[Seeder] - Seeder has already been run")
		return
	}

	logger.Info("[Seeder] - Register records in database...")
	for _, r := range records {
		marketRepository.Create(context.Background(), r)
	}
	logger.Info("[Seeder] finished successfully")
}

func ListTables(ctx context.Context, logger interfaces.ILogger, db *sql.DB) ([]string, error) {
	rawTables, err := db.QueryContext(ctx, "SHOW tables")
	if err != nil {
		return []string{}, err
	}
	defer rawTables.Close()

	var tables []string
	for rawTables.Next() {
		var table string
		err = rawTables.Scan(&table)
		if err != nil {
			return []string{}, err
		}

		tables = append(tables, table)
	}

	return tables, nil
}

func CreateMigrateTable(ctx context.Context, logger interfaces.ILogger, db *sql.DB) error {
	row := db.QueryRowContext(ctx, "CREATE TABLE migrations (name varchar not null, created_at TIMESTAMPTZ);")
	err := row.Err()
	if err != nil {
		return err
	}

	return nil
}

func ListMigrations() ([]string, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(workDir + "/migrate")
	if err != nil {
		return nil, err
	}

	migrates := []string{}
	for _, f := range files {
		migrates = append(migrates, f.Name())
	}

	return migrates, nil
}

func ExecuteMigrateUp(ctx context.Context, logger interfaces.ILogger, db *sql.DB, sqlFile string) error {
	if !strings.Contains(sqlFile, "up") {
		return nil
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dat, err := os.ReadFile(cwd + "migration/" + sqlFile)
	if err != nil {
		return err
	}

	row := db.QueryRowContext(ctx, string(dat))
	err = row.Err()
	if err != nil {
		return err
	}

	row = db.QueryRowContext(ctx, "INSERT INTO migrations (name, created_at) values ($1, $2)", sqlFile, time.Now())
	err = row.Err()
	if err != nil {
		return err
	}

	return nil
}

func Migrate() {
	if err := environments.NewEnvironment().Configure(); err != nil {
		log.Fatal(err)
	}
	os.Setenv("LOG_FILE", "./logs/seeder.log")

	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("[Migrator] - Connection to the database...")
	db, err := database.Connect(logger, make(chan bool))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	tablesCreated, err := ListTables(ctx, logger, db)
	if err != nil {
		log.Fatal(err)
	}

	if !contains(tablesCreated, "migrations") {
		if err := CreateMigrateTable(ctx, logger, db); err != nil {
			log.Fatal(err)
		}
	}

	tablesToCreate, err := ListMigrations()
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range tablesToCreate {
		if !contains(tablesCreated, t) {
			if err := ExecuteMigrateUp(ctx, logger, db, t); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func contains(slice []string, pattern string) bool {
	for _, v := range slice {
		if v == pattern {
			return true
		}
	}

	return false
}
