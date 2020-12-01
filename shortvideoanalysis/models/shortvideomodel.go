package models

// ShortVideoURL ...
type ShortVideoURL struct {
	FeedID          string `bson:"feed_id"`
	RealURL         string `bson:"real_url"`
	RealURLLossless string `bson:"real_url_lossless"`
}
