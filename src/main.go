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
var configPath = "config.json"
var mysqlConnString = ""
var mongoConnString = ""

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
	//Find relationship id
	for _, relationMappings := range schema.Content.Relationships.Tables {

		for _, mapid := range relationMappings.Ids {
			// Compare RelationShip Id with Mappings
			//If embedded path found you need find the Embedded element of the Document
			for id, data := range schema.Content.Mappings {
				if id == mapid {
					//If type is NEW_DOCUMENT it is one-to-one mapping (table to collection)
					if data.Settings.Type == models.NEW_DOCUMENT {
						wg.Add(1)
						go compareCount(dataCounts, &wg, strings.TrimPrefix(data.Table, "sakila.sakila."), schema.Name, schema.Content.Collections[data.CollectionId].Name, "", models.NEW_DOCUMENT, "")
					} else if data.Settings.Type == models.EMBEDDED_DOCUMENT_ARRAY && data.Settings.EmbeddedPath != "" {
						wg.Add(1)
						//TODO: For Embedded Doc Count we should use diff collectionid
						go compareCount(dataCounts, &wg, strings.TrimPrefix(data.Table, "sakila.sakila."), schema.Name, schema.Content.Collections[data.CollectionId].Name, data.Settings.EmbeddedPath, models.EMBEDDED_DOCUMENT_ARRAY, "")
					} else if data.Settings.Type == models.EMBEDDED_DOCUMENT {
						relKey := ""
						for key, val := range data.Fields {
							if val.Source.IsPrimaryKey {
								relKey = key
							}
						}
						wg.Add(1)
						//TODO: For Embedded Doc Count we should use diff collectionid
						go compareCount(dataCounts, &wg, strings.TrimPrefix(data.Table, "sakila.sakila."), schema.Name, schema.Content.Collections[data.CollectionId].Name, data.Settings.EmbeddedPath, models.EMBEDDED_DOCUMENT, relKey)
					}
				}
			}
		}
	}
	wg.Wait()
	PrintSummary(dataCounts)
}

//Compare SQL vs NoSQL row count
func compareCount(data map[string]models.RecordData, wg *sync.WaitGroup, tableName, dbName, collName, embeddedPath, embeddedType, relKey string) error {
	defer wg.Done()
	//get mysql client
	mysqlClient, err := clients2.GetMySqlClient(mysqlConnString)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//Get Row Count of a table
	rowCount := clients2.GetRowCount(mysqlClient, tableName)

	//Get mongodb collection count
	collCount, err := clients2.GetCollectionCount(mongoConnString, dbName, collName, embeddedPath, embeddedType, relKey)
	if err != nil {
		return err
	}
	//Update map for print summary
	data[tableName] = models.RecordData{RowCount: int64(rowCount), DocCount: collCount, CollectionName: collName}
	return err
}

func PrintSummary(dataCounts map[string]models.RecordData) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Table", "Row Count", "Collection ", "Doc Count", "Count Matched?"})

	for key, data := range dataCounts {
		matched := "yes"
		if data.DocCount != data.RowCount {
			matched = "no"
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
