package main

import (
	"context"
	"github.com/Kahyberth/read-csv-golang/model"
	"github.com/jackc/pgx/v5"
)

func uploadUsersToDB(conn *pgx.Conn, usuarios []model.User, batchSize int) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	for i := 0; i < len(usuarios); i += batchSize {
		end := i + batchSize
		if end > len(usuarios) {
			end = len(usuarios)
		}
		batch := &pgx.Batch{}

		for _, u := range usuarios[i:end] {
			batch.Queue(`INSERT INTO usuarios(nombre, email, edad) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`, u.ID, u.Gender, u.Age, u.EducationLevel, u.SocioeconomicStratum,
				u.CityOfResidence, u.ChildrenCount, u.SalaryMultiplier, u.IsRetired, u.CardType, u.IntentToBuyCard, u.ArticlesCount, u.ArticleType,
				u.MostPurchasedMonth, u.PurchaseInFirstHalf, u.MostWantedArticleType)
		}

		br := tx.SendBatch(context.Background(), batch)
		if err := br.Close(); err != nil {
			return err
		}
	}

	return tx.Commit(context.Background())
}
