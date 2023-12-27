fmt:
	./shell/fmt.sh

gen:
	./shell/gen.sh
	./shell/fmt.sh

build:
	docker compose build

up:
	docker compose up

ps:
	docker compose ps

minio:
	docker container run -d --name minio -p 9000:9000 -p 9001:9001 minio/minio server /data --console-address ":9001"
minio_old:
	docker container run -d --name minio -p 9000:9000 -p 9001:9001 minio/minio:RELEASE.2022-10-08T20-11-00Z server /data --console-address ":9001"