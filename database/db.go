package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
	"os"
	"strconv"
)

// Db is to be able to use it outside of the function.
var Db *sql.DB

func SetupDatabase() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred on .env file")
	}
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT")) // don't forget to convert int since port is int type.
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	// set up postgres sql to open it.
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database")
		panic(err)
	}
	Db = db
	fmt.Println("Successfully connected!")

}
