package dbmodel

import "imdb/graph/model"

type Actor struct {
	ID    uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
	Phone string `gorm:"not null"`
}

type Movie struct {
	ID          uint64       `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	ActorID     []uint64     `gorm:"not null"`
	Name        string       `gorm:"not null"`
	Description string       `gorm:"not null"`
	Status      model.Status `gorm:"not null"`
}
