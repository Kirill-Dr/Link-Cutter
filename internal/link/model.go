package link

import (
	"crypto/rand"
	"encoding/hex"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: GenerateHash(4),
	}
}

func GenerateHash(n int) string {
	bytes := make([]byte, n)

	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}
