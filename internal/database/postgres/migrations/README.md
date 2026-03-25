Use `golang-migrate` for schema changes.

Example commands:

```bash
migrate -path internal/database/postgres/migrations -database "$DATABASE_URL" up
migrate -path internal/database/postgres/migrations -database "$DATABASE_URL" down 1
```

Expected environment variables for the app connection:

```bash
DB_HOST
DB_PORT
DB_USER
DB_PASSWORD
DB_NAME
DB_SSLMODE
DB_MAX_CONNS
DB_MIN_CONNS
```
