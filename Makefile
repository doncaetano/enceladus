start:
	docker-compose up -d
stop:
	docker-compose stop
build:
	docker-compose build
dev-start:
	docker-compose --file docker-compose.dev.yml up -d
dev-stop:
	docker-compose --file docker-compose.dev.yml stop
