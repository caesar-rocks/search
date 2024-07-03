package search

import (
	"github.com/meilisearch/meilisearch-go"
)

type SearchConfig struct {
	Host   string
	APIKey string
}

type Search struct {
	Client *meilisearch.Client
}

func NewSearch(cfg *SearchConfig) *Search {
	db := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   cfg.Host,
		APIKey: cfg.APIKey,
	})
	return &Search{db}
}

func (search *Search) GetOrCreateIndex(index string) *meilisearch.Index {
	return search.Client.Index(index)
}

func (search *Search) UploadDocuments(index string, documents []map[string]interface{}) error {
	idx := search.GetOrCreateIndex(index)
	_, err := idx.AddDocuments(documents)
	return err
}

func (search *Search) Search(index string, query string, limit int64) ([]interface{}, error) {
	if limit <= 0 {
		limit = 10
	}
	searchRes, err := search.Client.Index(index).Search(
		query,
		&meilisearch.SearchRequest{
			Limit: limit,
		})
	if err != nil {
		return nil, err
	}
	return searchRes.Hits, err
}

func (search *Search) SearchWithFilter(index string, query string, filter string, limit int64) ([]interface{}, error) {
	if limit <= 0 {
		limit = 10
	}
	idx := search.GetOrCreateIndex(index)
	searchRes, err := idx.Search(
		query,
		&meilisearch.SearchRequest{
			Filter: filter,
			Limit:  limit,
		})
	if err != nil {
		return nil, err
	}
	return searchRes.Hits, err
}

func (search *Search) DeleteIndex(index string) error {
	_, err := search.Client.DeleteIndex(index)
	return err
}

func (search *Search) DeleteDocument(index string, documentId string) error {
	idx := search.GetOrCreateIndex(index)
	_, err := idx.DeleteDocument(documentId)
	return err
}

func (search *Search) DeleteAllDocuments(index string) error {
	idx := search.GetOrCreateIndex(index)
	_, err := idx.DeleteAllDocuments()
	return err
}
