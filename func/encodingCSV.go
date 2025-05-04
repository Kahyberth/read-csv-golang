package _func

import (
	"encoding/csv"
	"github.com/Kahyberth/read-csv-golang/model"
	"log"
	"os"
	"strconv"
)

func ReadCSVOne() ([]model.User, error) {
	file, err := os.OpenFile("clientes.csv", os.O_RDONLY|os.O_CREATE, os.ModePerm)

	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error al leer el archivo CSV: %v", err)
	}

	var users []model.User

	for _, row := range rows {
		if len(row) < 16 {
			log.Println("Fila con columnas insuficientes, se ignora:", row)
			continue
		}

		vals := make([]int, 16)
		for i, s := range row {
			vals[i], err = strconv.Atoi(s)
			if err != nil {
				log.Printf("Error al convertir valor '%s' a int en la columna %d: %v\n", s, i, err)
				continue
			}
		}

		user := model.User{
			ID:                    vals[0],
			Gender:                vals[1],
			Age:                   vals[2],
			EducationLevel:        vals[3],
			SocioeconomicStratum:  vals[4],
			CityOfResidence:       vals[5],
			ChildrenCount:         vals[6],
			SalaryMultiplier:      vals[7],
			IsRetired:             vals[8],
			CardType:              vals[9],
			IntentToBuyCard:       vals[10],
			ArticlesCount:         vals[11],
			ArticleType:           vals[12],
			MostPurchasedMonth:    vals[13],
			PurchaseInFirstHalf:   vals[14],
			MostWantedArticleType: vals[15],
		}

		users = append(users, user)
	}

	return users, nil
}
