package db

import "github.com/kuronosu/animeflv-api/pkg/scrape"

// Serial represents a sequence document
type Serial struct {
	ID  string `bson:"_id"`
	Seq int    `bson:"seq"`
}

type PaginatedAnimeResult struct {
	Page       int
	TotalPages int
	Count      int
	Animes     []scrape.Anime
}
