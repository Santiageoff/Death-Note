package bd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var ConnectionString string

func init() {
	// Cargar el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Construir la cadena de conexión
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
}

func GetDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}

const AllowedCORSDomain = "http://localhost:5173"
