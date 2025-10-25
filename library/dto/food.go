package dto

type Recipe struct {
	Id          string      `json:"id"`
	CreatorId   string      `json:"creator_id"`
	Name        string      `json:"name"`
	Ingredients Ingredients `json:"ingredients"`
	Recipe      string      `json:"recipe"`
}

type Ingredients struct {
	Name       string  `json:"name"`
	Measurment string  `json:"measurment"`
	Amount     float32 `json:"amount"`
}
