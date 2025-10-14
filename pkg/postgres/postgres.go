package postgres

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	err := godotenv.Load()
	var (
		host     = "localhost"
		port     = "5432"
		user     = "obras"
		password = "7894"
		dbname   = "obrasdb"
	)

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("conectado ao: " + dbname)

	return db, nil
}
