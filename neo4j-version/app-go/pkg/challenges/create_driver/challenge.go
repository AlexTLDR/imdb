package main

// Import the driver ("github.com/neo4j/neo4j-go-driver/v5/neo4j")
import (
	// Import the driver
	"context"
	"fmt"

	. "github.com/neo4j-graphacademy/neoflix/pkg/shared"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	// Neo4j Credentials
	credentials := GetNeo4jCredentials()

	// Cypher Query and Parameters
	cypher := `
      MATCH (p:Person)-[:DIRECTED]->(:Movie {title: $title})
      RETURN p.name AS Director
    `
	params := map[string]any{"title": "Toy Story"}

	// TODO: Create a DriverWithContext Instance

	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		credentials.Uri,
		neo4j.BasicAuth(credentials.Username, credentials.Password, ""),
	)
	PanicOnErr(err)

	// TODO: close the driver with defer PanicOnClosureError(ctx, driver)
	defer PanicOnClosureError(ctx, driver)

	// TODO: Open a new Session

	session := driver.NewSession(ctx, neo4j.SessionConfig{})

	// TODO: close the session with defer PanicOnClosureError(ctx, driver)

	defer PanicOnClosureError(ctx, session)

	// TODO: Run a Cypher statement
	result, err := session.Run(ctx, cypher, params)
	PanicOnErr(err)

	// TODO: Log the Director value of the first record
	director, err := neo4j.SingleTWithContext[string](ctx, result,
		func(record *neo4j.Record) (string, error) {
			director, _, err := neo4j.GetRecordValue[string](record, "Director")
			return director, err
		})
	PanicOnErr(err)
	fmt.Println(director)
}
