package services

import (
	"context"

	"github.com/szaborics/vacations-api/database"
	"github.com/szaborics/vacations-api/models"
)

// VacationService manages vacations
type VacationService interface {
	Get(ctx context.Context, filter *models.VacationFilter) ([]models.VacationDTO, error)
	GetAll(ctx context.Context) ([]models.VacationDTO, error)
}

// VacationServiceImpl implements VacationsService
type VacationServiceImpl struct {
	database database.Database
}

var _ VacationService = (*VacationServiceImpl)(nil)

// NewVacationService returns implementation of service
func NewVacationService() VacationService {
	database := database.GetDatabase()
	return &VacationServiceImpl{
		database: database,
	}
}

// Get retreives a list of all vacations, if there is no filter, it returns all
func (service *VacationServiceImpl) Get(ctx context.Context, filter *models.VacationFilter) ([]models.VacationDTO, error) {

	vacationsDAO, err := service.database.GetFilteredVacations(ctx, filter)

	if err != nil {
		return nil, err
	}

	vacationsDTOs := models.DAOsToDTOs(vacationsDAO)

	return vacationsDTOs, nil

}

// GetAll retreives a list of filtered vacations, if there is no filter, it returns all
func (service *VacationServiceImpl) GetAll(ctx context.Context) ([]models.VacationDTO, error) {

	vacationsDAO, err := service.database.GetAllVacations(ctx)

	if err != nil {
		return nil, err
	}

	vacationsDTOs := models.DAOsToDTOs(vacationsDAO)

	return vacationsDTOs, nil

}

// func getVacationsByCountry(country string) []models.VacationDTO {

// 	var filtered []models.VacationDTO

// 	places := readFileVacations().Places

// 	for _, place := range places {
// 		if strings.EqualFold(place.Country, country) {

// 			fmt.Println(place.String())
// 			filtered = append(filtered, place)

// 		}
// 	}
// 	return filtered

// }

// func readFileVacations() models.VacationsDTO {

// 	var filePath = "../../data/vacations.json"
// 	data, err := os.ReadFile(filePath)

// 	if err != nil {
// 		log.Fatalf("Error reading file: %v", err)

// 	}

// 	var vacations models.VacationsDTO
// 	err = json.Unmarshal(data, &vacations)

// 	if err != nil {
// 		log.Fatalf("error parsing JSON %v", err)
// 	}

// 	return vacations

// }
