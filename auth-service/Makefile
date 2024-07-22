CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

run:
	go run main.go

migrate_up:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/forum_auth -verbose up

migrate_down:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/forum_auth -verbose down

migrate_force:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/forum_auth -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1

docker-build:
	sudo docker build -t auth .

docker-migrate-up:
	sudo docker run --rm --network="net" -v $(CURRENT_DIR)/migrations:/migrations migrate/migrate -path=/migrations -database postgres://postgres:QodirovCoder@postgres:5432/forum_auth?sslmode=disable -verbose up

docker-migrate-down:
	sudo docker run --rm --network="net" -v $(CURRENT_DIR)/migrations:/migrations migrate/migrate -path=/migrations -database postgres://postgres:QodirovCoder@postgres:5432/forum_auth?sslmode=disable -verbose down

docker-migrate-force:
	sudo docker run --rm --network="net" -v $(CURRENT_DIR)/migrations:/migrations migrate/migrate -path=/migrations -database postgres://postgres:QodirovCoder@postgres:5432/forum_auth?sslmode=disable -verbose force 1
