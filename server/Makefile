.DEFAULT_GOAL := deploy

deploy:
	docker compose -f scripts/docker-compose.yaml up -d
	bash scripts/mongo/replset.sh

down:
	docker compose -f scripts/docker-compose.yaml down

