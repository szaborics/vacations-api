# Vacations API — Learning Curriculum

This is your personal roadmap for finishing this project. Each stage builds on the last, and every task is small enough to complete in a single sitting. When you come back after a break, start at **Re-orienting yourself** before touching any code.

---

## Development environment (one-time setup)

You need Docker installed. Then run this once to create your local database:

```bash
docker run -d \
  --name vacations-mongo \
  -p 27017:27017 \
  mongo:7
```

Seed it with the project's test data:

```bash
docker exec -i vacations-mongo \
  mongosh --eval "$(cat scripts/seedVacations.mongodb.js)"
```

That's it. The container persists between restarts — you just need to start it again each session (see the checklist below). Use `MONGODB_URI=mongodb://localhost:27017` for local development; keep your Atlas URI for production/cloud testing only.

To reset back to clean seed data at any time:

```bash
docker exec -i vacations-mongo \
  mongosh --eval 'db.getSiblingDB("vacationsApi").vacations.drop()'

docker exec -i vacations-mongo \
  mongosh --eval "$(cat scripts/seedVacations.mongodb.js)"
```

---

## When you come back after a break (start here every time)

Before writing a single line, spend 15 minutes re-orienting:

```bash
# 1. Start the local database (skip if it's already running)
docker start vacations-mongo

# 2. See what's changed since you were last here
git log --oneline -10

# 3. Run the tests — if they pass, nothing is broken
go test ./...

# 4. Run the server and hit it manually
MONGODB_URI=mongodb://localhost:27017 go run main.go
curl http://localhost:8080/vacations
curl "http://localhost:8080/vacations?country=Italy"
```

Then re-read two files:
- **`docs/api_v1.yaml`** — this is the contract. It defines every endpoint, every request shape, every response shape. It is your specification.
- **This file** — find the first unchecked task and start there.

---

## How the project is structured

```
main.go                     ← wires routes to handlers
controller/vacations.go     ← HTTP layer: reads request, calls service, writes response
services/vacations.go       ← business logic: calls database, maps DAO→DTO
database/vacation.go        ← database operations (already mostly complete)
models/vacations.go         ← shared data shapes used across all layers
docs/api_v1.yaml            ← the API specification (your source of truth)
```

Every new feature follows the same path: **database → service → controller → main.go**. When adding a new endpoint, you touch each layer in that order.

---

## Stage 1 — Get a working server (done ✅)

The foundation is already built:

- [x] MongoDB connection
- [x] `GET /vacations` with optional `?country=` and `?city=` filters
- [x] Controller, service, and database layers separated
- [x] Tests for the service and controller using mocks (no live database needed)

If `go test ./...` and `go run main.go` both work, you're ready to move on.

---

## Stage 2 — Health check (your warmup every time you restart)

**Why:** A health endpoint is the simplest possible handler — it takes no input, calls nothing, and returns a fixed response. It is the perfect warm-up task. Every real API has one.

**What to build:** `GET /health` → `200 {"status": "ok"}`

**Step by step:**

1. Add `HandleHealth` to the `VacationController` interface in `controller/vacations.go`
2. Implement it on `VacationControllerImpl` — it only needs to write a JSON response, no service call needed
3. Register the route in `main.go`: `mux.HandleFunc("GET /health", vacationController.HandleHealth)`
4. Test it: `curl http://localhost:8080/health`
5. Write a test in `controller/vacations_test.go` following the same pattern as `TestHandleRoot`

**You're done when:** `go test ./...` passes and `curl http://localhost:8080/health` returns `{"status":"ok"}`.

---

## Stage 3 — Read a single vacation by ID

**Why:** This introduces path parameters (`/vacations/{id}`) and a new type of error — "not found". It also teaches you how to propagate a feature through all three layers.

**What to build:** `GET /vacations/{id}` → `200 Vacation` or `404 {"code":"NOT_FOUND","message":"..."}`

**Step by step:**

1. **Service layer** — `GetVacationByID` already exists in `database/vacation.go`. Add it to the `VacationService` interface in `services/vacations.go` and implement it on `VacationServiceImpl`. It should call `service.database.GetVacationByID` and convert the result from DAO to DTO using `models.DAOToDTO`.

2. **Controller layer** — Add `GetByID` to the `VacationController` interface. Implement it on `VacationControllerImpl`:
   - Get the ID from the URL: `r.PathValue("id")`
   - Call `v.vacationService.GetByID(r.Context(), id)`
   - If the service returns an error, return `404`
   - If it succeeds, return `200` with the vacation JSON

3. **Router** — Add the route in `main.go`: `mux.HandleFunc("GET /vacations/{id}", vacationController.GetByID)`

4. **Test it manually:**
   - Run the server, then hit `GET /vacations` to get an ID from the response
   - `curl http://localhost:8080/vacations/<that-id>`
   - `curl http://localhost:8080/vacations/000000000000000000000000` (should 404)

5. **Write tests** — add cases to `controller/vacations_test.go` using the service mock, just like `TestGetFiltered`

**You're done when:** `go test ./...` passes and both the happy path and the not-found case work with curl.

---

## Stage 4 — Create a vacation

**Why:** This is your first write operation. It introduces reading a JSON request body, the `201 Created` status code, and the distinction between what the client sends (no ID) and what the server returns (with ID).

**What to build:** `POST /vacations` → `201 Vacation` or `400`/`422`

**Step by step:**

1. **Service layer** — `InsertVacation` already exists in `database/vacation.go`. Add `Create` to `VacationService` and implement it. It should accept a `models.VacationDTO`, convert it to a `VacationDAO` (the reverse of `DAOToDTO`), call `service.database.InsertVacation`, and return the created vacation (with the new ID) as a DTO.

   > You'll need to write a `DTOToDAO` helper in `models/vacations.go` — the reverse of `DAOToDTO`.

2. **Controller layer** — Add `Create` to `VacationController`. Implement it:
   - Decode the request body: `json.NewDecoder(r.Body).Decode(&input)`
   - Validate required fields: if `input.City` or `input.Country` is empty, return `422`
   - Call `v.vacationService.Create(r.Context(), input)`
   - Return `201` with the created vacation

3. **Router** — `mux.HandleFunc("POST /vacations", vacationController.Create)`

4. **Test it manually:**
   ```bash
   curl -X POST http://localhost:8080/vacations \
     -H "Content-Type: application/json" \
     -d '{"city":"Split","country":"Croatia","food":"grilled fish"}'
   ```
   The response should include the `id` that was assigned.

5. **Write tests** — test the happy path, missing required field (422), and a service error (500).

**You're done when:** `go test ./...` passes and you can create a vacation and see it appear in `GET /vacations`.

---

## Stage 5 — Delete a vacation

**Why:** Short task that reinforces the pattern. The `204 No Content` response (no body) is a useful HTTP convention to learn.

**What to build:** `DELETE /vacations/{id}` → `204` or `404`

**Step by step:**

1. **Service layer** — Add `Delete` to `VacationService`. Implement it: call `service.database.DeleteVacationByID`. If `DeletedCount` is 0, return a not-found error.

2. **Controller layer** — Add `Delete` to `VacationController`. Get the ID from `r.PathValue("id")`, call the service, return `204` on success with no body (`w.WriteHeader(http.StatusNoContent)`), or `404` if not found.

3. **Router** — `mux.HandleFunc("DELETE /vacations/{id}", vacationController.Delete)`

4. **Test it:** Create a vacation with POST, copy its ID, delete it, verify it's gone with GET.

5. **Write tests.**

**You're done when:** `go test ./...` passes and you can create-then-delete a vacation.

---

## Stage 6 — Update a vacation

**Why:** This is the most complex write operation. `PUT` replaces the entire resource — all fields, not just the ones you send. This teaches you the difference between PUT (full replace) and PATCH (partial update).

**What to build:** `PUT /vacations/{id}` → `200 Vacation` or `404`/`422`

**Step by step:**

1. **Database layer** — Add `UpdateVacation(ctx, id string, vacation VacationDAO) (*VacationDAO, error)` to `VacationRepository` in `database/vacation.go`. Implement it using `collection.ReplaceOne`. Also add it to the mock in `database/mocks/mockDatabase.go`.

2. **Service layer** — Add `Update` to `VacationService`. It accepts an ID and a DTO, converts the DTO to a DAO, calls `database.UpdateVacation`, and returns the updated vacation as a DTO.

3. **Controller layer** — Add `Update` to `VacationController`. Get the ID from the path, decode the body, validate required fields, call the service.

4. **Router** — `mux.HandleFunc("PUT /vacations/{id}", vacationController.Update)`

5. **Test it:** Create a vacation, update its food field, verify the change with GET.

6. **Update the mock** — after adding `UpdateVacation` to the interface, run `go generate ./...` or manually add the mock method to `database/mocks/mockDatabase.go` following the same pattern as the other methods.

**You're done when:** All CRUD operations work end-to-end and `go test ./...` passes.

---

## Stage 7 — Structured error responses

**Why:** Right now `sendJSONError` returns `{"error": "some message"}`. The API spec defines `{"code": "NOT_FOUND", "message": "..."}`. Matching your implementation to your spec is a professional habit — clients depend on consistent error shapes.

**What to do:**

1. Update `sendJSONError` in `controller/vacations.go` to accept a `code` string:
   ```go
   func sendJSONError(w http.ResponseWriter, status int, code, message string) {
       ...
       json.NewEncoder(w).Encode(map[string]string{"code": code, "message": message})
   }
   ```

2. Update all call sites to pass an error code (e.g. `"NOT_FOUND"`, `"INTERNAL_SERVER_ERROR"`, `"UNPROCESSABLE_ENTITY"`).

3. Update your controller tests — the `checkBody` assertions for error cases will need to check for `code` and `message` instead of `error`.

**You're done when:** Every error response has both `code` and `message` fields, matching `docs/api_v1.yaml`.

---

## Stage 8 — Request logging middleware

**Why:** Without logs you have no visibility into what your API is doing. Middleware is a powerful Go pattern — a function that wraps a handler to add behaviour without changing the handler itself.

**What to build:** Log every request: method, path, status code, duration.

**Step by step:**

1. Create `middleware/logging.go`:
   ```go
   package middleware

   func Logger(next http.Handler) http.Handler {
       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           start := time.Now()
           // wrap w to capture the status code
           next.ServeHTTP(w, r)
           log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
       })
   }
   ```

2. Wrap the mux in `main.go`: `http.ListenAndServe(":8080", middleware.Logger(mux))`

**You're done when:** Every request prints a line like `GET /vacations 12.4ms`.

---

## Stage 9 — Graceful shutdown

**Why:** Without this, pressing Ctrl+C while a request is in flight can drop that request mid-response. Graceful shutdown waits for in-flight requests to finish before exiting.

**What to do:** Replace the `http.ListenAndServe` call in `main.go` with an `http.Server` struct that listens for `SIGINT`/`SIGTERM` and calls `server.Shutdown(ctx)`.

Look up: `os/signal`, `signal.NotifyContext`, `srv.Shutdown`.

**You're done when:** Ctrl+C prints "shutting down" and the process exits cleanly.

---

## Done state checklist

When every item below is checked, the project is complete:

- [ ] `GET /health` returns `{"status":"ok"}`
- [ ] `GET /vacations` returns all vacations (optionally filtered)
- [ ] `GET /vacations/{id}` returns one vacation or 404
- [ ] `POST /vacations` creates a vacation and returns it with its ID
- [ ] `PUT /vacations/{id}` replaces a vacation and returns the updated version
- [ ] `DELETE /vacations/{id}` deletes a vacation and returns 204
- [ ] All error responses have `{"code":"...","message":"..."}`
- [ ] `go test ./...` passes with no live database required
- [ ] Every request is logged with method, path, and duration
- [ ] Ctrl+C shuts down cleanly

---

## Learning resources (Go-specific)

These are worth bookmarking rather than reading all at once — come back to them when you hit the relevant stage:

- **[net/http docs](https://pkg.go.dev/net/http)** — handler signatures, ResponseWriter, request body
- **[encoding/json docs](https://pkg.go.dev/encoding/json)** — Decode, Encode, struct tags
- **[Go by Example](https://gobyexample.com)** — short, practical examples for any Go concept
- **[100 Go Mistakes](https://100go.co)** — common traps, worth skimming once you finish the project

---

*Last updated to reflect the project state on this branch. The spec in `docs/api_v1.yaml` is the authoritative source of truth for what "done" looks like.*
