.DEFAULT_GOAL := deploy

deploy:
	docker compose -f scripts/docker-compose.yaml up -d

init:
	bash scripts/mongo/replset.sh

down:
	docker compose -f scripts/docker-compose.yaml down
	rm -rf scripts/data

up:
	make deploy
	bash scripts/mongo/replset.sh

restart:
	make down && make up
