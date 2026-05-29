package database

import (
	"context"
	"fmt"

	"github.com/szaborics/vacations-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// VacationRepository defines vacation related operations (abstracts database operations allowing for implementation for any database)
type VacationRepository interface {
	GetVacationCount(ctx context.Context, filter *models.VacationFilter) (int64, error)
	GetFilteredVacations(ctx context.Context, filter *models.VacationFilter) ([]models.VacationDAO, error)
	InsertVacation(ctx context.Context, vacation models.VacationDAO) (primitive.ObjectID, error)
	DeleteVacationByID(ctx context.Context, vacationID string) (int64, error)
	GetVacationByID(ctx context.Context, vacationID string) (*models.VacationDAO, error)
}

// DATABASE variable defines the mongo database to connect to
const DATABASE = "vacationsApi"

// COLLECTION variable defines the mongo collection to execute CRUD operations against
const COLLECTION = "vacations"

func (db *MongoDB) getQueryFilter(filter *models.VacationFilter) bson.D {

	query := bson.D{}

	if filter.City != "" {
		query = append(query, bson.E{Key: "city", Value: filter.City})
	}
	if filter.Country != "" {
		query = append(query, bson.E{Key: "country", Value: filter.Country})

	}

	return query
}

// GetVacationCount gets the filtered count of vacations from the database
func (db *MongoDB) GetVacationCount(ctx context.Context, filter *models.VacationFilter) (int64, error) {
	collection, err := GetNoCloseClientCollection(DATABASE, COLLECTION)
	if err != nil {
		return 0, err
	}

	query := db.getQueryFilter(filter)

	count, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return 0, err
	}

	return count, err
}

// GetFilteredVacations gets the filtered vacations from the database
func (db *MongoDB) GetFilteredVacations(ctx context.Context, filter *models.VacationFilter) ([]models.VacationDAO, error) {
	var vacations []models.VacationDAO

	collection, err := GetNoCloseClientCollection(DATABASE, COLLECTION)
	if err != nil {
		return nil, err
	}

	query := db.getQueryFilter(filter)

	cur, err := collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cur.Close((ctx))

	if err := cur.All(ctx, &vacations); err != nil {
		return nil, err
	}

	return vacations, nil
}

// InsertVacation inserts one new vacation to the vacations collection
func (db *MongoDB) InsertVacation(ctx context.Context, vacation models.VacationDAO) (primitive.ObjectID, error) {

	collection, err := GetNoCloseClientCollection(DATABASE, COLLECTION)
	if err != nil {
		return primitive.NilObjectID, err
	}

	vacation.ID = primitive.NewObjectID()

	res, err := collection.InsertOne(ctx, vacation)
	if err != nil {
		return primitive.NilObjectID, err
	}

	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("failed to cast inserted ID to ObjectID")
	}

	return insertedID, nil

}

// DeleteVacationByID deletes one vacation from the vacations collection
func (db *MongoDB) DeleteVacationByID(ctx context.Context, vacationID string) (int64, error) {

	id, err := primitive.ObjectIDFromHex(vacationID)
	if err != nil {
		return 0, fmt.Errorf("failed to initialize objectID: %v", err)
	}
	query := bson.M{"_id": id}

	collection, err := GetNoCloseClientCollection(DATABASE, COLLECTION)
	if err != nil {
		return 0, err
	}
	res, err := collection.DeleteOne(ctx, query)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil

}

// GetVacationByID returns one vacation from the vacations collection using the ID
func (db *MongoDB) GetVacationByID(ctx context.Context, vacationID string) (*models.VacationDAO, error) {
	var vacation models.VacationDAO
	id, err := primitive.ObjectIDFromHex(vacationID)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize objectID: %v", err)
	}
	query := bson.M{"_id": id}

	collection, err := GetNoCloseClientCollection(DATABASE, COLLECTION)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, query).Decode(&vacation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Vacation %s not found", vacationID)
		}
		return nil, err
	}

	return &vacation, nil

}
