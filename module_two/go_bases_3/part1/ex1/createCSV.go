package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

/*
	Referências:
	Para manipular uma struct e extrair insumos: https://stackoverflow.com/questions/24337145/get-name-of-struct-field-using-reflection
	Para criação de arquivos CSV e como manipulá-los: https://webdamn.com/write-data-to-csv-file-using-golang/
*/

func writeData(items []Product, csvFile io.Writer) {
	var fields []string

	baseStruct := reflect.Indirect(reflect.ValueOf(items[0]))
	var numFields = baseStruct.NumField()

	for i := 0; i < numFields; i++ {
		fields = append(fields, baseStruct.Type().Field(i).Name)
	}

	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Write(fields)

	for _, currentEmptyRow := range items {
		row := reflect.Indirect(reflect.ValueOf(currentEmptyRow))

		var currentLine []string

		// Adiciona os dados específicos de uma linha para o slice de strings chamado currentLine
		for _, field := range fields {
			currentCSVContentLine := fmt.Sprint(row.FieldByName(field))
			currentLine = append(currentLine, currentCSVContentLine)
		}

		csvWriter.Write(currentLine)
	}
	csvWriter.Flush()

}

func CreateCSV(produtos []Product, filename string) {
	filename = filename + ".csv"
	csvFile, err := os.Create(filename)

	if err != nil {
		log.Fatalf("Falha ao criar o arquivo %s. Ocorreu o seguinte erro: %v", filename, err.Error())
	}

	writeData(produtos, csvFile)

	csvFile.Close()
}
