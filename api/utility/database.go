package utility

import (
	"context"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Database_Connect_Write() (context.Context, neo4j.DriverWithContext, neo4j.SessionWithContext, error) {
	// Get Environment Variables
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	// Connect To Database
	driver, err := neo4j.NewDriverWithContext(db_host, neo4j.BasicAuth(db_user, db_pass, ""))
	if err != nil {
		return nil, nil, nil, err
	}
	// Create Context
	ctx := context.Background()
	// Create Session
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	// Return
	return ctx, driver, session, nil
}
