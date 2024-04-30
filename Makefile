rm:
	docker compose stop \
	&& docker compose rm \
	&& sudo rm -rf pgdata/

up:
	docker compose -f docker-compose.yml up --force-recreate

rb:
	docker build . -t registrationtogames-app
net-wel:
	docker network connect reg-to-games welcome-app-1
net-user:
	docker network connect reg-to-games user-app-1
net-sch:
	docker network connect reg-to-games schedule-app-1
net-media:
	docker network connect reg-to-games media-app-1
net-reg:
	docker network connect reg-to-games registration-app-1
net-set:
	docker network connect reg-to-games settings-app-1
#	docker logs registrationtogames-postgresql-1 >& logs/postgres.log
#	docker logs registrationtogames-app-1 >& logs/app.log
net:
	docker network connect reg-to-games welcome-app-1
	docker network connect reg-to-games user-app-1
	docker network connect reg-to-games schedule-app-1
	docker network connect reg-to-games media-app-1
	docker network connect reg-to-games registration-app-1
	docker network connect reg-to-games settings-app-1