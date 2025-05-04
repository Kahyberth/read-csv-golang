package lib

import (
	"fmt"
	"github.com/Kahyberth/read-csv-golang/model"
	"strconv"
)

func ParseRowToUser(row []string) (model.User, error) {
	if len(row) < 16 {
		return model.User{}, fmt.Errorf("fila incompleta: %v", row)
	}

	vals := make([]int, 16)
	for i, s := range row {
		v, err := strconv.Atoi(s)
		if err != nil {
			return model.User{}, fmt.Errorf("error en columna %d: %v", i, err)
		}
		vals[i] = v
	}

	return model.User{
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
	}, nil
}
