package database

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// NewNeo4jConnection creates a new neo4j connection
func NewNeo4jConnection() (neo4j.Driver, error) {
	target := "neo4j://localhost:7687"

	driver, err := neo4j.NewDriver(
		target,
		neo4j.BasicAuth("neo4j", "", ""),
		func(c *neo4j.Config) {
			c.Encrypted = false
		})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Neo4j Server")

	return driver, nil
}

// Neo4jRepository is a Neo4j DB repository
type Neo4jRepository struct {
	Connection neo4j.Driver
}

// func GetNeo4jCredentials() Neo4jCredentials {
// 	return Neo4jCredentials{
// 		Uri:      os.Getenv("NEO4J_URI"),
// 		Username: os.Getenv("NEO4J_USERNAME"),
// 		Password: os.Getenv("NEO4J_PASSWORD"),
// 	}
// }
