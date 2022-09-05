package main

import (
	"fmt"
	clients2 "verifier/src/clients"
)

var dbName = "migration2"
var collName = "film"
var tableName = "film"

//var connString = "mongodb://localhost:27017"
//var connString = "mongodb+srv://gov-mongopush:HvkzKmiA1shW8YkN@mongodbsource.kn30v.mongodb.net/?retryWrites=true&w=majority"
//var mongoConnString = "mongodb+srv://saviour:saviour@saviour.dwhf9.mongodb.net/gdelt"

var mongoConnString = fmt.Sprintf("mongodb://localhost:27017/%s", dbName)
var mysqlConnString = "application-user:application-user123@tcp(34.66.118.177:3306)/sakila"

func main() {

	//get mysql client
	mysqlClient, err := clients2.GetMySqlClient(mysqlConnString)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Get Row Count of a table
	rowCount := clients2.GetRowCount(mysqlClient, tableName)

	//Get mongodb collection count
	collCount := clients2.GetCollectionCount(mongoConnString, dbName, collName)
	if int64(rowCount) != collCount {
		fmt.Printf("Documents count mismatched, mysql:%s table count is :%d but mongodb:%s Colleciton count is :%d", tableName, rowCount, collName, collCount)
		return
	}
	fmt.Println("Continue checking data comparison... ")

}
