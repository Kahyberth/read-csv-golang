package _func

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/Kahyberth/read-csv-golang/lib"
	"github.com/Kahyberth/read-csv-golang/model"
)

func ReadCSVOne() ([]model.User, time.Duration, error) {
	start := time.Now()
	file, err := os.OpenFile("clientes.csv", os.O_RDONLY|os.O_CREATE, os.ModePerm)

	if err != nil {
		log.Printf("Error al abrir el archivo: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	r := csv.NewReader(file)
	r.Comma = ';'

	rows, err := r.ReadAll()
	if err != nil {
		log.Printf("Error al leer el archivo CSV: %v\n", err)
	}

	var users []model.User

	for _, row := range rows {
		user, err := lib.ParseRowToUser(row)
		if err != nil {
			log.Println("Error al procesar fila:", err)
			continue
		}
		users = append(users, user)
	}

	elapsed := time.Since(start)

	return users, elapsed, nil
}
