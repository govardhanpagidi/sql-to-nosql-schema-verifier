package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	clients2 "verifier/src/clients"
	models "verifier/src/models"
)

var dbName = "migration2"
var collName = "film"
var tableName = "film"

var schemaPath = "schema.json"

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

	//load migration schema file
	content, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("error in reading %s schema file.", schemaPath)
	}

	var schema models.MigrationSchema
	err = json.Unmarshal(content, &schema)

	if err != nil {
		log.Fatalf("error in unmarshalling %s schema data.", schemaPath)
	}

	for _, value := range schema.Content.Mappings {
		if value.Table == tableName {
			value.CollectionId
			//Get collection id from mappings

			//Find the collection names from relation ships
		}
	}

}
