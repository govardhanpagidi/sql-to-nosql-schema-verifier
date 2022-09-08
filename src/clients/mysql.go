package clients

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getMySqlClient(connString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)
	//db.SetConnMaxLifetime(time.Second * 10)
	if err != nil {
		panic(err.Error())
	}
	return db, err
}

func getRowCount(db *sql.DB, tableName string) int {
	rows, err := db.Query(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			fmt.Println(err)
		}
	}
	return count
}

var mysqlClient *sql.DB

func GetSqlRowCount(mysqlConnString, tableName string) (int, error) {

	if mysqlClient == nil {
		mysqlClient, err := getMySqlClient(mysqlConnString)

		if err != nil {
			log.Fatal(err)
			return 0, err
		}
		rowCount := getRowCount(mysqlClient, tableName)
		return rowCount, err
	}
	//Get Row Count of a table
	rowCount := getRowCount(mysqlClient, tableName)
	return rowCount, nil
}
