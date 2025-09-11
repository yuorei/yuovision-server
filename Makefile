fmt:
	./shell/fmt.sh

build:
	docker build -t yuovision-server .

run:
	./docker-run.sh --build

run-dev:
	./docker-run.sh --dev --build

run-prod:
	./docker-run.sh --prod --build

stop:
	docker stop yuovision-server || true

ps:
	docker ps --filter name=yuovision-server

test:
	go test -v ./...

minio:
	docker container run -d --name minio -p 9000:9000 -p 9001:9001 minio/minio server /data --console-address ":9001"
minio_old:
	docker container run -d --name minio -p 9000:9000 -p 9001:9001 minio/minio:RELEASE.2022-10-08T20-11-00Z server /data --console-address ":9001"

gen:
	go generate ./...

lint:
	./shell/lint.sh

prod:
	set -a && source .env.prod && set +a&&\
	go run main.go

dev:
	set -a && source .env.dev && set +a&&\
	go run main.go

setup_dev:
	./shell/setup_dev.sh