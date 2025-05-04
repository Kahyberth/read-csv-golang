package lib

import (
	"context"
	"fmt"
	"github.com/Kahyberth/read-csv-golang/model"
	"github.com/jackc/pgx/v5"
	"log"
)

func UploadUsersToDB(conn *pgx.Conn, usuarios []model.User, batchSize int) error {
	// Limpia la tabla antes de insertar nuevos datos
	_, err := conn.Exec(context.Background(), "DROP TABLE IF EXISTS users")
	if err != nil {
		return fmt.Errorf("error al truncar la tabla: %v", err)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {

		}
	}(tx, context.Background())

	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            gender INT,
            age INT,
            education_level INT,
            socioeconomic_stratum INT,
            city_of_residence INT,
            children_count INT,
            salary_multiplier INT,
            is_retired INT,
            card_type INT,
            intent_to_buy_card INT,
            articles_count INT,
            article_type INT,
            most_purchased_month INT,
            purchase_in_first_half INT,
            most_wanted_article_type INT
        );
    `

	_, err = conn.Exec(context.Background(), createTableSQL)
	if err != nil {
		log.Fatalf("unable to create table: %v\n", err)
	}
	//TODO: BatchSize = 1000
	for i := 0; i < len(usuarios); i += batchSize {
		end := i + batchSize
		if end > len(usuarios) {
			end = len(usuarios)
		}
		batch := &pgx.Batch{}

		for _, u := range usuarios[i:end] {
			batch.Queue(`INSERT INTO users (id, gender, age, education_level, socioeconomic_stratum, city_of_residence, 
                                      children_count, salary_multiplier, is_retired, card_type, intent_to_buy_card, 
                                      articles_count, article_type, most_purchased_month, purchase_in_first_half, 
                                      most_wanted_article_type) 
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`,
				u.ID, u.Gender, u.Age, u.EducationLevel, u.SocioeconomicStratum,
				u.CityOfResidence, u.ChildrenCount, u.SalaryMultiplier, u.IsRetired, u.CardType,
				u.IntentToBuyCard, u.ArticlesCount, u.ArticleType,
				u.MostPurchasedMonth, u.PurchaseInFirstHalf, u.MostWantedArticleType)
		}

		br := tx.SendBatch(context.Background(), batch)
		if err := br.Close(); err != nil {
			return err
		}
	}

	return tx.Commit(context.Background())
}
