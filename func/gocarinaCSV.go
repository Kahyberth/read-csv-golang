package _func

import (
	"encoding/csv"
	"fmt"
	"github.com/Kahyberth/read-csv-golang/model"
	"github.com/gocarina/gocsv"
	"io"
	"os"
	"time"
)

func ReadCSVGocarina() ([]model.User, error) {

	start := time.Now()

	file, err := os.Open("clientes.csv")
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		reader := csv.NewReader(in)
		reader.Comma = ';'
		return reader
	})

	var users []model.User

	if err := gocsv.UnmarshalFile(file, &users); err != nil {
		return nil, fmt.Errorf("error al leer el archivo CSV: %v", err)
	}
	fmt.Print(users)

	elapsed := time.Since(start)
	fmt.Printf("Tiempo de lectura (Gocarina): %.2f s\n", elapsed.Seconds())

	return users, nil
}
