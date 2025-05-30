build-server:
	docker-compose -f docker-compose.production.server.yaml build

build-web:
	docker-compose -f docker-compose.production.web.yaml build

build-all:
	docker-compose -f docker-compose.production.server.yaml build
	docker-compose -f docker-compose.production.web.yaml build

build-all-no-cache:
	docker-compose -f docker-compose.production.server.yaml build --no-cache
	docker-compose -f docker-compose.production.web.yaml build --no-cache


