package database

import (
	"context"
	CONFIG "video-api/config"

	"github.com/olivere/elastic/v7"
)

var elasticsearchClient *elastic.Client
var err error

// Connect - connect to elastisearch database with given configuration
func ConnectElastisearch() {
	elasticsearchClient, err = elastic.NewClient(elastic.SetURL(CONFIG.ElastisearchConfig), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}
}

// SearchInElastisearch - serch via query
func SearchInElastisearch(query, index string) (*elastic.MultiSearchResult, error) {
	return elasticsearchClient.MultiSearch().Add(
		elastic.NewSearchRequest().Index(index).Source(query),
	).Do(context.TODO())
}

func AddIndex(index, mapping string) {
	elasticsearchClient.CreateIndex(index).BodyString(mapping).Do(context.TODO())
}
