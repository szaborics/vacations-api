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

// VacationDAO ...
type VacationDAO struct {
	ID               primitive.ObjectID `bson:"_id, omitempty"`
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

func vacationsToString(vacations []VacationDTO) []string {
	var vacationStrings []string

	for _, vac := range vacations {
		vacationStrings = append(vacationStrings, vac.String())
	}
	return vacationStrings
}

// DAOToDTO translates the database retreived values to the service object
func DAOToDTO(dao *VacationDAO) VacationDTO {
	return VacationDTO{
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
