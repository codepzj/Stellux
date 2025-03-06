.DEFAULT_GOAL := restart

deploy:
	openssl rand -base64 756 > ./scripts/mongo/keyFile
	chmod 400 ./scripts/mongo/keyFile
	chown 999:999 ./scripts/mongo/keyFile

	docker compose -f scripts/docker-compose.yaml up -d

init:
	bash scripts/mongo/replset.sh

down:
	docker compose -f scripts/docker-compose.yaml down
	rm -rf scripts/data
	rm -rf static/
	rm -rf scripts/mongo/keyFile
	
up:
	make deploy
	bash scripts/mongo/replset.sh

restart:
	make down && make up
