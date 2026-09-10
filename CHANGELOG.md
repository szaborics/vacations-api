# Changelog

## (08/04/2026)

- Add `GET /health` endpoint returning `{"status":"ok"}`
- Add `GET /vacations/{id}` endpoint to fetch a single vacation by its MongoDB ID
- Add `GetVacationByID` to the service layer and wire it through to the controller
- Fix `VacationDTO` missing `id` field — responses now include the MongoDB-generated ID so callers can use it in subsequent requests
- Default `MONGODB_URI` to `mongodb://localhost:27017` when not set, so the server starts without needing the env var during local development
- Fix seed script (`scripts/seedVacations.mongodb.js`) — was defined as an arrow function that never executed and used the wrong database name; rewritten as plain executable statements targeting the correct `vacationsApi` database
- Add `CURRICULUM.md` — personal learning roadmap with staged tasks, re-orientation checklist, and a done-state checklist for the finished API
- Add `//go:build integration` tag to `database/vacation_test.go` so it is excluded from `go test ./...` and only runs when explicitly targeting a live database
- Delete `database/mocks/testifymock_vacation_test.go` — was testing the testify mock framework itself rather than application code

## (01/07/2025)

- restructured tests to use mocks
- update API spec
- Remove GetAll throughout — GetFilteredVacations with an empty filter
covers the same case without the duplication.
- Reorganize mocks to use package mocks (was package database/services),
which is the standard mockery convention and avoids import name
conflicts when both the real package and mock are imported together.
Add a VacationService mock in services/mocks/.
- Rewrite controller and service tests to inject mocks instead of calling
NewVacationController/NewVacationService at init, which required a live
Atlas connection to run. Tests now use seed data from
scripts/seedVacations.mongodb.js as mock return values.
- Update API spec to define the ideal done state: add PUT /vacations/{id}
and GET /health, split Vacation/VacationInput so responses include id
but request bodies don't require it, replace the Error schema with
{code, message} for machine-readable handling, add 422 for validation
failures, extract shared parameters and responses into components.


## (01/07/2025)

- Added CRUD database operations
    - TO DO Next: add mocked interface for mongodb so that i can mock the database connection for tests instead of messing with live data
    - To Do After: add http server 