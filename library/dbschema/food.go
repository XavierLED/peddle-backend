package dbschema

type Recipe struct {
	Id          string      `json:"id" bson:"_id,omitempty"`
	CreatorId   string      `json:"creator_id" bson:"creator_id"`
	Name        string      `json:"name" bson:"name"`
	Ingredients Ingredients `json:"ingredients" bson:"ingredients"`
	Recipe      string      `json:"recipe" bson:"recipe"`
}

type Ingredients struct {
	Name       string  `json:"name" bson:"name"`
	Measurment string  `json:"measurment" bson:"measurment"`
	Amount     float32 `json:"amount" bson:"amount"`
}
