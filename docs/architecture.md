# Architecture Blueprint

## Product scope

The MVP has two tabs:

- `Engineering Blogs`
- `Startups`

Both are read-only and backed by Postgres. Data is imported by a periodic job outside this phase. The frontend never connects directly to the database; it reads through the Go API.

## Repository structure

```text
frontend/
  app/
    components/
    lib/
    routes/
    styles/
backend/
  cmd/api/
  internal/
    api/
    config/
    domain/
    repository/
    server/
  migrations/
docs/
design/
```

## Request flow

1. Remix route loader receives the request.
2. Remix loader calls the Echo API.
3. Echo handler validates query parameters.
4. Repository uses GORM to query Postgres and map rows into domain records.
5. Echo returns JSON to Remix.
6. Remix renders the response into list pages and cards.

## Design implementation notes

The visual system is based on the existing design references:

- editorial shell with a fixed left rail
- slate and paper surface hierarchy
- `Manrope` for display/headings
- `Inter` for UI/body copy
- filter pills and soft card surfaces instead of hard borders

The first implementation should keep that direction intact and avoid generic dashboard styling.

## Future-safe decisions

The MVP excludes users and saved state, but the architecture leaves clean seams for them:

- imported records use stable string `id` values across storage and API boundaries
- timestamps are present for imported records
- tags are modeled through join tables
- frontend routes are isolated from persistence details
- API responses are shaped so personalization fields can be added later without breaking core list rendering

## Recommended next implementation slices

1. Install frontend and backend dependencies.
2. Wire database access and run migrations.
3. Finish the backend query layer and response tests.
4. Flesh out the Remix shell and responsive layouts.
5. Add infinite-scroll behavior for the blog feed and richer filters for startups.
