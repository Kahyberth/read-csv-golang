package model

type User struct {
	ID                    int `csv:"0" json:"id"`
	Gender                int `csv:"1" json:"gender"`
	Age                   int `csv:"2" json:"age"`
	EducationLevel        int `csv:"3" json:"education_level"`
	SocioeconomicStratum  int `csv:"4" json:"socioeconomic_stratum"`
	CityOfResidence       int `csv:"5" json:"city_of_residence"`
	ChildrenCount         int `csv:"6" json:"children_count"`
	SalaryMultiplier      int `csv:"7" json:"salary_multiplier"`
	IsRetired             int `csv:"8" json:"is_retired"`
	CardType              int `csv:"9" json:"card_type"`
	IntentToBuyCard       int `csv:"10" json:"intent_to_buy_card"`
	ArticlesCount         int `csv:"11" json:"articles_count"`
	ArticleType           int `csv:"12" json:"article_type"`
	MostPurchasedMonth    int `csv:"13" json:"most_purchased_month"`
	PurchaseInFirstHalf   int `csv:"14" json:"purchase_in_first_half"`
	MostWantedArticleType int `csv:"15" json:"most_wanted_article_type"`
}
