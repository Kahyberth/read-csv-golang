package model

type User struct {
	ID                    int `json:"id"`
	Gender                int `json:"gender"`
	Age                   int `json:"age"`
	EducationLevel        int `json:"education_level"`
	SocioeconomicStratum  int `json:"socioeconomic_stratum"`
	CityOfResidence       int `json:"city_of_residence"`
	ChildrenCount         int `json:"children_count"`
	SalaryMultiplier      int `json:"salary_multiplier"`
	IsRetired             int `json:"is_retired"`
	CardType              int `json:"card_type"`
	IntentToBuyCard       int `json:"intent_to_buy_card"`
	ArticlesCount         int `json:"articles_count"`
	ArticleType           int `json:"article_type"`
	MostPurchasedMonth    int `json:"most_purchased_month"`
	PurchaseInFirstHalf   int `json:"purchase_in_first_half"`
	MostWantedArticleType int `json:"most_wanted_article_type"`
}
