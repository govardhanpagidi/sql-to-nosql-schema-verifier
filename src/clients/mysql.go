package clients

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetMySqlClient(connString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	return db, err
}

func GetRowCount(db *sql.DB, tableName string) int {
	rows, err := db.Query(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	return count
}
