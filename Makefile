migrate:
	goose -dir=migrations/ postgres "postgres:///app05?sslmode=disable" up
migrate-down:
	goose -dir=migrations/ postgres "postgres:///app05?sslmode=disable" down