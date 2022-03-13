package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"markets/pkg/app/interfaces"
	valueObjects "markets/pkg/domain/value_objects"
	"markets/pkg/infra/database"
	"markets/pkg/infra/environments"
	"markets/pkg/infra/logger"
	"markets/pkg/infra/repositories"
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

func main() {
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
	fileDir := currentDir + "/db/DEINFO_AB_FEIRASLIVRES_2014.csv"

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
