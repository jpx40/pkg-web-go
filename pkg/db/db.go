package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var Migration bool = false

// type db gorm.DB
func Connect() *pgx.Conn {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Listen and Server in 0.0.0.0:3000
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return conn
}

func GetAllPackage(db *pgx.Conn) []Package {
	// result := db.find(&p).rowsaffected

	var resultPackageList []Package

	rows, _ := db.Query(context.Background(), "SELECT * FROM packages")
	for rows.Next() {

		var name string
		var version string
		var size string

		err := rows.Scan(&name, &size, &version)
		if err != nil {
			panic(err)
		}
		resultPackage := Package{Name: name, Version: version, Size: size}

		resultPackageList = append(resultPackageList, resultPackage)
	}

	return resultPackageList
}
