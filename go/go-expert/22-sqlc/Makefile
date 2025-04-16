create-migration:
	migrate create -ext sql -dir sql/migrations -seq $(name)

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

migrate-down:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down


.PHONY: migrate create-migration