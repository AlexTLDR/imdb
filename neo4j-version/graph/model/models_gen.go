// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node interface {
	IsNode()
	GetUUID() string
}

type Movie struct {
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Tagline   string    `json:"tagline"`
	Released  int       `json:"released"`
	Directors []*Person `json:"directors"`
	Writers   []*Person `json:"writers"`
	Cast      []*Person `json:"cast"`
}

func (Movie) IsNode()              {}
func (this Movie) GetUUID() string { return this.UUID }

// Participation represents a person's role in a movie
type Participation struct {
	Role  string `json:"role"`
	Movie *Movie `json:"movie"`
}

type Person struct {
	UUID         string           `json:"uuid"`
	Name         string           `json:"name"`
	Born         int              `json:"born"`
	Role         *string          `json:"role,omitempty"`
	Participated []*Participation `json:"participated"`
}

func (Person) IsNode()              {}
func (this Person) GetUUID() string { return this.UUID }
