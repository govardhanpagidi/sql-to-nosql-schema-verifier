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
	"time"
	"verifier/src/clients"
	"verifier/src/models"
)

var dbName = "migration2"
var sqlDBName = "sakila"
var collName = "film"
var tableName = "film"

var schemaPath = "schema.json"
var configPath = "config.json"
var mysqlConnString = ""
var mongoConnString = ""

var dataCounts map[string]models.RecordData

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	dataCounts = map[string]models.RecordData{}
	//load migration schema file
	content, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("error in reading %s schema file.", schemaPath)
	}
	var schema models.MigrationSchema
	err = json.Unmarshal(content, &schema)

	//Load config file
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("error in reading %s schema file.", schemaPath)
	}
	var config models.Config
	err = json.Unmarshal(data, &config)

	if err != nil {
		log.Fatalf("error in unmarshalling %s schema data.", schemaPath)
	}
	mongoConnString = config.Mongodb.ConnectionString
	mysqlConnString = strings.TrimPrefix(config.Jdbc.Url, "jdbc:mysql://")

	var wg sync.WaitGroup

	//Iterate the tables
	for tabName, relationMappings := range schema.Content.Relationships.Tables {
		if len(relationMappings.Ids) > 1 {

		}
		collCount := int64(0)
		tableName := strings.TrimPrefix(tabName, "sakila.sakila.")
		collectionName := ""
		fmt.Println("Comparing : ", tableName)
		wg.Add(1)
		//Traverse all the tables and compare the data counts
		go compareData(&wg, dataCounts, relationMappings, schema, collectionName, collCount, tableName)
		time.Sleep(time.Millisecond * 200)
	}
	wg.Wait()
	PrintSummary(dataCounts)
}

func compareData(s *sync.WaitGroup, dataCounts map[string]models.RecordData, relationMappings models.Mappings, schema models.MigrationSchema, collectionName string, collCount int64, tableName string) {
	defer s.Done()
	collections := ""
	for _, mapid := range relationMappings.Ids {
		// Compare RelationShip Id with Mappings
		//If embedded path found you need find the Embedded element of the Document
		for id, data := range schema.Content.Mappings {
			relKey := ""
			if id == mapid {
				//If type is NEW_DOCUMENT it is one-to-one mapping (table to collection)
				if data.Settings.Type == models.EMBEDDED_DOCUMENT {
					//If type is EMBEDDED_DOCUMENT, the table is embedded document or inside the embedded document
					for key, val := range data.Fields {
						if val.Source.IsPrimaryKey {
							relKey = key
						}
					}
				}

				collectionName = schema.Content.Collections[data.CollectionId].Name

				//Get mongodb collection count
				count, err := clients.GetCollectionCount(mongoConnString, schema.Name, schema.Content.Collections[data.CollectionId].Name, data.Settings.EmbeddedPath, data.Settings.Type, relKey)
				if err != nil {
					fmt.Println(err)
					return
				}
				if data.Settings.EmbeddedPath != "" {
					collections += collectionName + "." + data.Settings.EmbeddedPath + "\n"
				} else {

					collections += collectionName + "\n"
				}

				collCount += count
			}
		}
	}

	//Table Count
	sqlRowCount, err := clients.GetSqlRowCount(mysqlConnString, tableName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//Update map for print summary
	dataCounts[tableName] = models.RecordData{RowCount: int64(sqlRowCount), DocCount: collCount, CollectionName: collections}
	return
}

func PrintSummary(dataCounts map[string]models.RecordData) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Table", "Row_Count", "Collection ", "Doc_Count", "Matched?"})
	for key, data := range dataCounts {
		matched := "YES"
		if data.DocCount < data.RowCount {
			matched = "NO"
		}
		table.Append(
			[]string{key,
				strconv.FormatInt(data.RowCount, 10),
				data.CollectionName,
				strconv.FormatInt(data.DocCount, 10),
				matched})

	}
	table.Render()

}
