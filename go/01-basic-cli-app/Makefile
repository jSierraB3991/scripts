dev:
	sudo docker build --target dev . -t "videos-app:1.0"

build:
	sudo docker build . -t "videos-app:1.0"

run-get:
	sudo docker run -it videos-app:1.0 get --all

run-get-help:
	sudo docker run -it videos-app:1.0 get --help

run-add-help:
	sudo docker run -it videos-app:1.0 add --help

run-add-incorrect-format:
	sudo docker run -it videos-app:1.0 add -video "format video json"

#for inser data run next command
# sudo docker run -it videos-app:1.0 add -video "format video json"
# if you not have idea of the json format for video, run "make run-add-incorrect-format", and show a guide format
all: build run-get
