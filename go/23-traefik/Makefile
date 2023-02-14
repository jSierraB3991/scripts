network:
	docker network create traefik-net
traefik:
	docker compose up -d
whoami:
	docker compose -f ./whoami.docker-compose.yml up -d

all: network traefik whoami
