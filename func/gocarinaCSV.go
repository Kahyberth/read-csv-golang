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

func ReadCSVOcarina() ([]model.User, error) {
	start := time.Now()

	originalFile, err := os.ReadFile("clientes.csv")
	if err != nil {
		return nil, fmt.Errorf("error al leer archivo original: %v", err)
	}

	header := "id;gender;age;education_level;socioeconomic_stratum;city_of_residence;children_count;salary_multiplier;is_retired;card_type;intent_to_buy_card;articles_count;article_type;most_purchased_month;purchase_in_first_half;most_wanted_article_type\n"

	tempFilePath := "clientes_with_header.csv"
	err = os.WriteFile(tempFilePath, []byte(header+string(originalFile)), 0644)
	if err != nil {
		return nil, fmt.Errorf("error al escribir archivo temporal: %v", err)
	}

	file, err := os.Open(tempFilePath)
	if err != nil {
		return nil, fmt.Errorf("error al abrir archivo temporal: %v", err)
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
		return nil, fmt.Errorf("error al parsear CSV: %v", err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Tiempo de lectura (Gocarina): %.2f s\n", elapsed.Seconds())

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(tempFilePath)

	return users, nil
}
