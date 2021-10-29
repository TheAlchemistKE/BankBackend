.PHONY: migrate-up migrate-down sqlc test

migrate-up:
	@echo "Migrating DB."
	migrate -database 'postgresql://localhost/bank?user=njeri&password=KelynPNjeri@1998' -path db/migration up

migrate-down:
	@echo "Destroying Tables."
	migrate -database 'postgresql://localhost/bank?user=njeri&password=KelynPNjeri@1998' -path db/migration down


sqlc:
	@echo "Generating Queries"
	sqlc generate

test:
	go test -v -cover ./...
