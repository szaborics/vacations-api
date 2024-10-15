package models

import (
	"fmt"
	"strings"
)

type Location struct {
	City             string `json:"city"`
	Country          string `json:"country"`
	Food             string `json:"food"`
	Wine             string `json:"wine"`
	PointsOfInterest POI    `json:"pointsOfInterest"`
}

type Locations struct {
	Places []Location `json:"places"`
}

type POI struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type POIs struct {
	PointsOfInterest []POIs `json:"pointsOfInterest"`
}

func (ls *Locations) GetLocations(country string) string {
	return "i don't work yet"

}

func (l Location) String() string {
	return fmt.Sprintf(" City : %s, Country: %s, Wine: %s, Food: %s, Points of Interest:%s", l.City, l.Country, l.Wine, l.Food, l.PointsOfInterest)
}

func locationsToString(locations []Location) string {
	var locationStrings []string

	for _, loc := range locations {
		locationStrings = append(locationStrings, loc.String())
	}
	return strings.Join(locationStrings, "; ")
}
