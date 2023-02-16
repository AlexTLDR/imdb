package dbmodel

import "imdb/graph/model"

type Actor struct {
	ID    string `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
	Phone string `gorm:"not null"`
}

type Movie struct {
	ID          string       `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	ActorID     string       `gorm:"not null"`
	Name        string       `gorm:"not null"`
	Description string       `gorm:"not null"`
	Status      model.Status `gorm:"not null"`
}
