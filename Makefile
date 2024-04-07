rm:
	docker compose stop \
	&& docker compose rm \
	&& sudo rm -rf pgdata/

up:
	docker compose -f docker-compose.yml up --force-recreate

rb:
	docker build . -t registrationtogames-app

#	docker logs registrationtogames-postgresql-1 >& logs/postgres.log
#	docker logs registrationtogames-app-1 >& logs/app.log