package models

type Personality struct {
	Id			string    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Personalities []Personality

func GetAllPersonalities() ([]Personality, error) {
	return Personalities, nil
}