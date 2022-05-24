package main

import (
	"encoding/csv"
	"log"
	"os"
)

/*
	https://stackoverflow.com/questions/24999079/reading-csv-file-in-go
*/

func getCSV(filePath string) [][]string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Não foi possível ler o arquivo", filePath, " | erro:", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Não foi possível parsear os dados "+filePath, err)

	}
	return records
}

func ReadCSV(pathname string) [][]string {
	file := getCSV(pathname)

	return file
}
