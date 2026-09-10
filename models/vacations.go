package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VacationsDTO ...
type VacationsDTO struct {
	Places []VacationDTO `json:"places"`
}

// VacationDTO ...
type VacationDTO struct {
	ID               string `json:"id"`
	City             string `json:"city"`
	Country          string `json:"country"`
	Food             string `json:"food"`
	Wine             string `json:"wine"`
	PointsOfInterest []POI  `json:"pointsOfInterest"`
}

// POI ...
type POI struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

// VacationDAO ...
type VacationDAO struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	City             string             `bson:"city"`
	Country          string             `bson:"country"`
	Food             string             `bson:"food"`
	Wine             string             `bson:"wine"`
	PointsOfInterest []POI              `bson:"pointsOfInterest"`
}

// VacationFilter ...
type VacationFilter struct {
	Country string
	City    string
}

func (v VacationDTO) String() string {
	return fmt.Sprintf(" City : %s, Country: %s, Wine: %s, Food: %s, Points of Interest:%s", v.City, v.Country, v.Wine, v.Food, v.PointsOfInterest)
}
func (p POI) String() string {
	return fmt.Sprintf("Name : %s, Description: %s", p.Name, p.Description)
}

// DAOToDTO translates the database retreived values to the service object
func DAOToDTO(dao *VacationDAO) VacationDTO {
	return VacationDTO{
		ID:               dao.ID.Hex(),
		City:             dao.City,
		Country:          dao.Country,
		Food:             dao.Food,
		Wine:             dao.Wine,
		PointsOfInterest: dao.PointsOfInterest,
	}
}

// DAOsToDTOs translates multiple DAOs to DTOs
func DAOsToDTOs(daos []VacationDAO) []VacationDTO {
	dtos := make([]VacationDTO, len(daos))
	for i, dao := range daos {
		dtos[i] = DAOToDTO(&dao)
	}
	return dtos
}
