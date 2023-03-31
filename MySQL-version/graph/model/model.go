package model

type Movie struct {
	ID int `json:"_id"`
	//	Actors      []*Actor `json:"actors"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}
