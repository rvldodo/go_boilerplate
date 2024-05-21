package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"

	"github.com/rvldodo/boilerplate/db"
	"github.com/rvldodo/boilerplate/domain/model"
	"github.com/rvldodo/boilerplate/lib/log"
)

func init() {
	goose.AddMigrationContext(upUser, downUser)
}

func upUser(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	err := db.DB.WithContext(ctx).
		Set("gorm:table_options", TABLE_OPTIONS).
		AutoMigrate(&model.User{})
	if err != nil {
		log.Errorf("Error migrate user: %v", err)
		return err
	}
	return nil
}

func downUser(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	err := db.DB.WithContext(ctx).Migrator().DropTable(&model.User{})
	if err != nil {
		log.Errorf("Error drop user table: %v", err)
		return err
	}
	return nil
}
