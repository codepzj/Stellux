web:
	docker-compose -f docker-compose.production.yaml build --no-cache web

server:
	docker-compose -f docker-compose.production.yaml build --no-cache server 

admin:
	docker-compose -f docker-compose.production.yaml build --no-cache admin

all:
	docker-compose -f docker-compose.production.yaml build --no-cache