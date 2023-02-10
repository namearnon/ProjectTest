package model

import "encoding/base64"

type Beer struct {
	ID        int
	BeerName  string
	BeerType  string
	BeerDesc  string
	BearImage string
}

type Log struct {
	ID        int
	LogMethod string
	LogDesc   string
}

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
