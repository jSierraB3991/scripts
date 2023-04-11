awsbuild:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/main cmd/main.go

compress:
		zip -jrm build/main.zip build/main

deploy:
		cowsay "Deploying"

clean:
		rm build/main.zip

all: awsbuild compress deploy