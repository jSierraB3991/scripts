package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/eliotandelon/gotesting/models"
)

type MysqlRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) MysqlRepository {
	return MysqlRepository{DB: db}
}

func CreateConnection(ctx context.Context, configuration *models.Configuration) *sql.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		configuration.DB.User,
		configuration.DB.Password,
		configuration.DB.Host,
		configuration.DB.Port,
		configuration.DB.Name,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	return db
}
