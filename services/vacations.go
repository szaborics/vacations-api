package services

import (
	"context"

	"github.com/szaborics/vacations-api/database"
	"github.com/szaborics/vacations-api/models"
)

// VacationService manages vacations
type VacationService interface {
	Get(ctx context.Context, filter *models.VacationFilter) ([]models.VacationDTO, error)
	GetVacationByID(ctx context.Context, id string) (*models.VacationDTO, error)
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

// GetVacationByID retrieves a vacation by its ID
func (service *VacationServiceImpl) GetVacationByID(ctx context.Context, id string) (*models.VacationDTO, error) {
	vacationDAO, err := service.database.GetVacationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	vacationDTO := models.DAOToDTO(vacationDAO)

	return &vacationDTO, nil
}
