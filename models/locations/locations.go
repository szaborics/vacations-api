package locations

type Location struct {
	City             string `json:"city"`
	Country          string `json:"country"`
	Food             string `json:"food"`
	Wine             string `json:"wine"`
	PointsOfInterest POI    `json:"pointsOfInterest"`
}

type Locations struct {
	Locations []Location `json:"locations"`
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
