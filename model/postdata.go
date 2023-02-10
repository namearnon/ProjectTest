package model

import "mime/multipart"

type PostData struct {
	BeerName string `json:"beerName,omitempty"`
}

type GetBeerData struct {
	ID        string
	BeerName  string
	BeerType  string
	BeerDesc  string
	BeerImage *multipart.FileHeader
}
