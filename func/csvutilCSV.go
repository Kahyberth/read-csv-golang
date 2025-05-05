package _func

import (
	"encoding/csv"
	"fmt"
	"github.com/Kahyberth/read-csv-golang/model"
	"github.com/jszwec/csvutil"
	"io"
	"os"
	"time"
)

func ReadUtils() ([]model.User, time.Duration, error) {
	start := time.Now()

	originalFile, err := os.Open("clientes.csv")
	if err != nil {
		return nil, 0, fmt.Errorf("error al abrir el archivo original: %w", err)
	}
	defer func(originalFile *os.File) {
		err := originalFile.Close()
		if err != nil {

		}
	}(originalFile)

	tempFile, err := os.CreateTemp("", "users_with_header_*.csv")
	if err != nil {
		return nil, 0, fmt.Errorf("error al crear archivo temporal: %w", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(tempFile.Name())

	header := "id;gender;age;education_level;socioeconomic_stratum;city_of_residence;children_count;salary_multiplier;is_retired;card_type;intent_to_buy_card;articles_count;article_type;most_purchased_month;purchase_in_first_half;most_wanted_article_type\n"
	if _, err := tempFile.WriteString(header); err != nil {
		return nil, 0, fmt.Errorf("error al escribir encabezado: %w", err)
	}

	if _, err := io.Copy(tempFile, originalFile); err != nil {
		return nil, 0, fmt.Errorf("error al copiar datos: %w", err)
	}

	if _, err := tempFile.Seek(0, 0); err != nil {
		return nil, 0, fmt.Errorf("error al posicionarse en el inicio: %w", err)
	}

	reader := csv.NewReader(tempFile)
	reader.Comma = ';'

	decoder, err := csvutil.NewDecoder(reader)
	if err != nil {
		return nil, 0, fmt.Errorf("error al crear el decoder: %w", err)
	}

	var users []model.User
	for {
		var user model.User
		if err := decoder.Decode(&user); err != nil {
			if err == io.EOF {
				break
			}
			return nil, 0, fmt.Errorf("error al decodificar: %w", err)
		}
		users = append(users, user)
	}
	elapsed := time.Since(start)

	return users, elapsed, nil
}
