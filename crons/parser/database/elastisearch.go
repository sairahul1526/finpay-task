package database

import (
	"context"
	CONFIG "video-parser/config"

	"github.com/olivere/elastic/v7"
)

var elasticsearchClient *elastic.Client
var err error

// Connect - connect to elastisearch database with given configuration
func ConnectElastisearch() {
	elasticsearchClient, err = elastic.NewClient(elastic.SetURL(CONFIG.ElastisearchConfig), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
}

// InsertIntoElastisearch - bulk add documents into elasti search
func InsertIntoElastisearch(index string, documents []interface{}, ids []string) error {
	bulkRequest := elasticsearchClient.Bulk().Index(index)
	for i, document := range documents {
		req := elastic.NewBulkCreateRequest().Id(ids[i]).UseEasyJSON(true).Doc(document)
		bulkRequest = bulkRequest.Add(req)
	}
	_, err := bulkRequest.Do(context.TODO())
	return err
}
