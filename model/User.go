package model

type User struct {
	ID                    int `csv:"id"`
	Gender                int `csv:"gender"`
	Age                   int `csv:"age"`
	EducationLevel        int `csv:"education_level"`
	SocioeconomicStratum  int `csv:"socioeconomic_stratum"`
	CityOfResidence       int `csv:"city_of_residence"`
	ChildrenCount         int `csv:"children_count"`
	SalaryMultiplier      int `csv:"salary_multiplier"`
	IsRetired             int `csv:"is_retired"`
	CardType              int `csv:"card_type"`
	IntentToBuyCard       int `csv:"intent_to_buy_card"`
	ArticlesCount         int `csv:"articles_count"`
	ArticleType           int `csv:"article_type"`
	MostPurchasedMonth    int `csv:"most_purchased_month"`
	PurchaseInFirstHalf   int `csv:"purchase_in_first_half"`
	MostWantedArticleType int `csv:"most_wanted_article_type"`
}
