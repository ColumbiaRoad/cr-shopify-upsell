package db

import "github.com/ColumbiaRoad/cr-shopify-upsell/backend/lib/database"

// Database provides database implementation for storage/retrieval interfaces
// needed by the core modules
type Database struct {
	*database.Database
}

// NewDatabase creates a new database
func NewDatabase() (*Database, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}
	return &Database{Database: db}, nil
}
