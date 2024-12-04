package database

import (
	"context"
	"fmt"

	"github.com/szaborics/vacations-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

type vacationRepository interface {
	GetVacationCount(ctx context.Context, filter *models.VacationFilter) (int64, error)
	GetAllVacations(ctx context.Context) ([]models.VacationDAO, error)
	GetFilteredVacations(ctx context.Context, filter *models.VacationFilter) ([]models.VacationDAO, error)
}

func (db *MongoDB) getQueryFilter(filter *models.VacationFilter) bson.M {
	query := bson.M{}

	if filter.City != "" {
		query["city"] = filter.City
	}
	if filter.Country != "" {
		query["country"] = filter.Country
	}

	return query
}

// GetVacationCount ...
func (db *MongoDB) GetVacationCount(ctx context.Context, filter *models.VacationFilter) (int64, error) {
	collection := MongodbConnect()
	query := db.getQueryFilter(filter)

	fmt.Printf("query: %v, filter: %v", query, filter)

	count, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return 0, err
	}

	defer collection.Database().Client().Disconnect(ctx)

	return count, err
}

// GetFilteredVacations ...
func (db *MongoDB) GetFilteredVacations(ctx context.Context, filter *models.VacationFilter) ([]models.VacationDAO, error) {
	var vacations []models.VacationDAO
	collection := MongodbConnect()

	query := db.getQueryFilter(filter)

	cur, err := collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cur.Close((context.Background()))

	cur.All(context.Background(), &vacations)

	return vacations, err
}

// GetAllVacations ...
func (db *MongoDB) GetAllVacations(ctx context.Context) ([]models.VacationDAO, error) {

	var vacations []models.VacationDAO
	collection := MongodbConnect()

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close((context.Background()))

	cur.All(context.Background(), &vacations)

	return vacations, err
}
