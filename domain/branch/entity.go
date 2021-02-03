package branch

import (
	"openbankingcrawler/domain/subentities"

	"github.com/go-bongo/bongo"
)

type branchIdentification struct {
	Type       string `json:"type"`
	Code       string `json:"code"`
	CheckDigit string `json:"checkDigit"`
	Name       string `json:"name"`
}

type branchPostalAddress struct {
	Address               string `json:"address"`
	DistrictName          string `json:"districtName"`
	TownName              string `json:"townName"`
	CountrySubDivision    string `json:"countrySubDivision"`
	PostCode              string `json:"postCode"`
	GeographicCoordinates struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	} `json:"geographicCoordinates"`
}

//Entity branch entity
type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	InstitutionID      string               `json:"institutionid"`
	Identification     branchIdentification `json:"identification"`
	PostalAddress      branchPostalAddress  `json:"postalAddress"`
	Availability       struct {
		Standards []struct {
			Weekday     string `json:"weekday"`
			OpeningTime string `json:"openingTime"`
			ClosingTime string `json:"closingTime"`
		} `json:"standards"`
		Exception         string `json:"exception"`
		AllowPublicAccess bool   `json:"allowPublicAccess"`
	} `json:"availability"`
	Phones   []subentities.Phone `json:"phones"`
	Services struct {
		Codes          []string `json:"codes"`
		AdditionalInfo string   `json:"additionalInfo"`
	} `json:"services"`
}

//NewEntity create a new branch entity
func NewEntity(institutionID string) *Entity {

	return &Entity{
		InstitutionID: institutionID,
	}
}
