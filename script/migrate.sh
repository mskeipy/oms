#! /bin/sh

rm ./internal/infrastructure/database/migration/.env
cp .env ./internal/infrastructure/database/migration
go run ./internal/infrastructure/database/migration/migrate.go