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