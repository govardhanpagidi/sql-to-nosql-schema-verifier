package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/tablewriter"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	clients2 "verifier/src/clients"
	models "verifier/src/models"
)

var dbName = "migration2"
var sqlDBName = "sakila"
var collName = "film"
var tableName = "film"

var schemaPath = "schema.json"
var mysqlConnString = ""
var mongoConnString = ""

const NEW_DOCUMENT = "NEW_DOCUMENT"
const EMBEDDED_DOCUMENT_ARRAY = "EMBEDDED_DOCUMENT_ARRAY"

var dataCounts map[string]models.RecordData

func main() {

	dataCounts = map[string]models.RecordData{}
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
	mongoConnString = schema.ConnectionDetails.Mongodb.ConnectionString
	mysqlConnString = strings.TrimPrefix(schema.ConnectionDetails.Jdbc.Url, "jdbc:mysql://")

	var wg sync.WaitGroup
	//Find relationship id
	for _, relationMappings := range schema.Content.Relationships.Tables {
		//if tabName == sqlDBName+"."+sqlDBName+"."+tableName {
		//Traverse relationsship ids , mostly one
		for _, mapid := range relationMappings.Ids {
			// Compare RelationShip Id with Mappings
			//If embedded path found you need find the Embedded element of the Document
			for id, data := range schema.Content.Mappings {
				if id == mapid {
					//If type is NEW_DOCUMENT it is one-to-one mapping (table to collection)
					if data.Settings.Type == NEW_DOCUMENT {
						wg.Add(1)
						go compareCount(dataCounts, &wg, strings.TrimPrefix(data.Table, "sakila.sakila."), schema.Name, schema.Content.Collections[data.CollectionId].Name, "")
					} else if data.Settings.Type == EMBEDDED_DOCUMENT_ARRAY && data.Settings.EmbeddedPath != "" {
						wg.Add(1)
						//TODO: For Embedded Doc Count we should use diff collectionid
						go compareCount(dataCounts, &wg, strings.TrimPrefix(data.Table, "sakila.sakila."), schema.Name, schema.Content.Collections[data.CollectionId].Name, data.Settings.EmbeddedPath)
					}
				}
			}
		}
	}
	wg.Wait()
	PrintSummary(dataCounts)

}

func compareCount(data map[string]models.RecordData, wg *sync.WaitGroup, tableName, dbName, collName, embeddedPath string) {
	defer wg.Done()
	//get mysql client
	mysqlClient, err := clients2.GetMySqlClient(mysqlConnString)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Get Row Count of a table
	rowCount := clients2.GetRowCount(mysqlClient, tableName)

	//Get mongodb collection count
	collCount := clients2.GetCollectionCount(mongoConnString, dbName, collName, embeddedPath)
	data[tableName] = models.RecordData{RowCount: int64(rowCount), DocCount: collCount}
	//if int64(rowCount) != collCount {
	//	fmt.Printf("Documents count mismatched, mysql:%s table count is :%d but mongodb:%s Colleciton count is :%d", tableName, rowCount, collName, collCount)
	//	fmt.Println()
	//	dataCounts[tableName] = models.RecordData{RowCount: int64(rowCount), DocCount: collCount}
	//	return
	//}
	//fmt.Printf(" Matched, Table:%s RowCount:%d , DocCount:%d", tableName, rowCount, collCount)
	//fmt.Println()
	return
}

func PrintSummary(dataCounts map[string]models.RecordData) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Table Name", "No Of Rows", "No Of Docs", "Count Matched?"})

	for key, data := range dataCounts {
		table.Append([]string{key, strconv.FormatInt(data.RowCount, 10), strconv.FormatInt(data.DocCount, 10), strconv.FormatBool(data.DocCount == data.RowCount)})
	}
	table.Render()
}
