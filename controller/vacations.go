package controller

import "github.com/szaborics/vacations-api/services"

//VacationController handles vacation requests
type VacationController interface{

}
//VacationCotrollerImpl implements vacation service
type VacationControllerImpl struct{
	vacationService services.VacationService
}

