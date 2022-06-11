build:
	sudo podman build . -f Containerfile -t scraping-transcaribe
build_docker:
	sudo docker . -t scraping-transcaribe

run:
	sudo podman run scraping-transcaribe
run_docker:
	sudo docker run scraping-transcaribe

all_docker: build_docker run_docker
all: build run
