# Changelog

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