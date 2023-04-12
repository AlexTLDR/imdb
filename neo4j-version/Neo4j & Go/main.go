package main

import (
	"context"
	"fmt"
	"time"

	. "github.com/neo4j-graphacademy/neoflix/pkg/shared"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	// Create a Driver Instance
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		"neo4j+s://dbhash.databases.neo4j.io",    // (1)
		neo4j.BasicAuth("neo4j", "letmein!", ""), // (2)
	)

	// Verify Connectivity
	PanicOnErr(driver.VerifyConnectivity(ctx))

	// Open a new Session
	session := driver.NewSession(ctx, neo4j.SessionConfig{})

	defer PanicOnClosureError(ctx, driver)
	defer PanicOnClosureError(ctx, session)

	// Execute a Cypher statement in an auto-commit transaction
	result, err := session.Run(
		ctx, // (1)
		`
	MATCH (p:Person)-[:DIRECTED]->(:Movie {title: $title})
	RETURN p
	`, // (2)
		map[string]any{"title": "The Matrix"}, // (3)
		func(txConfig *neo4j.TransactionConfig) {
			txConfig.Timeout = 3 * time.Second // (4)
		},
	)
	PanicOnErr(err)

	// Get all Person nodes
	people, err := neo4j.CollectTWithContext[neo4j.Node](ctx, result,
		func(record *neo4j.Record) (neo4j.Node, error) {
			person, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
			return person, err
		})
	PanicOnErr(err)
	fmt.Println(people)

	// Close the driver to release any resources
	driver.Close(ctx)

}
