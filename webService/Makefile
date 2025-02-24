.PHONY: build up down

build:
	docker build -t webservice:latest .
	
up:
	docker-compose up -d
	# docker container logs -f webservice

down:
	docker-compose down --remove-orphans
	rm -rf pgdata
	docker system prune -f