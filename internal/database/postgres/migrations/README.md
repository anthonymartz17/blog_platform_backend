Use `golang-migrate` for schema changes.

Example commands:

```bash
migrate -path internal/database/postgres/migrations -database "$DATABASE_URL" up
migrate -path internal/database/postgres/migrations -database "$DATABASE_URL" down 1
```

Seed development data with:

```bash
psql "$DATABASE_URL" -f internal/database/postgres/seeds/001_dev_seed.sql
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

Seeded users use the same placeholder bcrypt hash and are meant only for local development.
