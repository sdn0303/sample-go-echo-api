package adapters

import (
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/sdn0303/sample-go-echo-api/pkg/ent"
)

// Database adapter
type Database struct {
	Client *ent.Client
}

// NewDatabase constructs a new database
func NewDatabase(dbURI string) (*Database, error) {
	drv, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	// add custom sql.DB settings
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client := ent.NewClient(ent.Driver(drv))
	return &Database{Client: client}, nil
}
