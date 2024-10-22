package models

import (
	"fmt"
)

// Locations ...
type Locations struct {
	Places []Location `json:"places"`
}

// Location ...
type Location struct {
	City             string `json:"city"`
	Country          string `json:"country"`
	Food             string `json:"food"`
	Wine             string `json:"wine"`
	PointsOfInterest []POI  `json:"pointsOfInterest"`
}

// POI ...
type POI struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (l Location) String() string {
	return fmt.Sprintf(" City : %s, Country: %s, Wine: %s, Food: %s, Points of Interest:%s", l.City, l.Country, l.Wine, l.Food, l.PointsOfInterest)
}
func (p POI) String() string {
	return fmt.Sprintf("Name : %s, Description: %s", p.Name, p.Description)
}

func locationsToString(locations []Location) []string {
	var locationStrings []string

	for _, loc := range locations {
		locationStrings = append(locationStrings, loc.String())
	}
	return locationStrings
}
