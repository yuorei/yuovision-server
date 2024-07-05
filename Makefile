fmt:
	./shell/fmt.sh

build:
	docker compose build

up:
	docker compose up

ps:
	docker compose ps

test:
	go test -v ./...

minio:
	docker container run -d --name minio -p 9000:9000 -p 9001:9001 minio/minio server /data --console-address ":9001"
minio_old:
	docker container run -d --name minio -p 9000:9000 -p 9001:9001 minio/minio:RELEASE.2022-10-08T20-11-00Z server /data --console-address ":9001"

migration:
	set -a && source .env.prod && set +a&&\
	atlas schema apply \
	-u "mysql://$${MYSQL_USER}:$${MYSQL_PASSWORD}@$${MYSQL_HOST}:$${MYSQL_PORT}/$${MYSQL_DATABASE}" \
	--to file://db/atlas/schema.hcl

schema_output:
	mkdir -p db/atlas &&\
	set -a && source .env.prod && set +a&&\
	atlas schema inspect -u "mysql://$${MYSQL_USER}:$${MYSQL_PASSWORD}@$${MYSQL_HOST}:$${MYSQL_PORT}/$${MYSQL_DATABASE}" > db/atlas/schema.hcl

sql_output:
	mkdir -p db/atlas &&\
	set -a && source .env.prod && set +a&&\
	atlas schema inspect -u "mysql://$${MYSQL_USER}:$${MYSQL_PASSWORD}@$${MYSQL_HOST}:$${MYSQL_PORT}/$${MYSQL_DATABASE}" --format "{{ sql . \" \" }}" > db/atlas/schema.sql

gen:
	./shell/gen_db.sh

lint:
	./shell/lint.sh

prod:
	set -a && source .env.prod && set +a&&\
	PORT=$${PORT} \
	AWS_ACCESS_KEY_ID=$${AWS_ACCESS_KEY_ID} \
	AWS_SECRET_ACCESS_KEY=$${AWS_SECRET_ACCESS_KEY} \
	AWS_S3_ENDPOINT=$${AWS_S3_ENDPOINT} \
	AWS_S3_URL=$${AWS_S3_URL} \
	AUTH_URL=$${AUTH_URL} \
	REDIS_ADDRESS=$${REDIS_ADDRESS} \
	MYSQL_HOST=$${MYSQL_HOST} \
	MYSQL_PORT=$${MYSQL_PORT} \
	MYSQL_DATABASE=$${MYSQL_DATABASE} \
	MYSQL_PASSWORD=$${MYSQL_PASSWORD} \
	MYSQL_USER=$${MYSQL_USER} \
	go run main.go