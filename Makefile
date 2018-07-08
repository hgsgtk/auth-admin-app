dev:
	go run main.go admin.go auth.go config.go template.go staff.go

build:
	docker-compose build

run:
	docker-compose up -d

browse:
	open http://localhost:8080

stop:
	docker stop auth-admin-app