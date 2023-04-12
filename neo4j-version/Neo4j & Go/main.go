package main

import (
	"context"

	. "github.com/neo4j-graphacademy/neoflix/pkg/shared"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		"neo4j+s://dbhash.databases.neo4j.io",    // (1)
		neo4j.BasicAuth("neo4j", "letmein!", ""), // (2)
	)
}
