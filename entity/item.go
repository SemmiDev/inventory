package entity

import (
	"net/url"

	"gorm.io/gorm"
)

type Status string

func (s Status) String() string {
	// Clean the string from any escape characters
	// Example: Hello%20World -> Hello World
	cleanedStr, err := url.QueryUnescape(string(s))
	if err != nil {
		return ""
	}
	return cleanedStr
}

const (
	BarangMasuk  Status = "Barang Masuk"
	BarangKeluar Status = "Barang Keluar"
)

type Item struct {
	gorm.Model
	Name   string `json:"name"`
	Status Status `json:"status"`
}
