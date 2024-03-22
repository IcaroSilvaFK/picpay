up: 
	migrate -path ./migrations -database "postgres://docker:docker@localhost:5432/picpay?sslmode=disable" -verbose up

down: 
	migrate -path ./migrations -database "postgres://docker:docker@localhost:5432/picpay?sslmode=disable" -verbose down

.PHONY: up