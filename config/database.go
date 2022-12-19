package config

import (
	"FinalProject/config/migrations"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	Err error
)


func Connect() *sql.DB {
	port, _ := strconv.Atoi(os.Getenv("PGPORT"))

	psqlInfo := fmt.Sprintf(
					"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
					os.Getenv("PGHOST"), 
					port, 
					os.Getenv("PGUSER"), 
					os.Getenv("PGPASSWORD"), 
					os.Getenv("PGDATABASE"),
			)

	DB, Err = sql.Open("postgres", psqlInfo)
	if Err != nil {
		panic(Err)
	}

	migrations.DbMigrate(DB)

	return DB
}

