rm:
	docker compose stop \
	&& docker compose rm \
	&& sudo rm -rf pgdata \
	&& sudo rm -rf kafkadata/

up:
	docker compose -f docker-compose.yml up --force-recreate

rb:
	docker build . -t logs-app

#&& sudo rm -rf tmp/