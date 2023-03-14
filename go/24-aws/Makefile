clean:
		rm main
		rm main.zip
all:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
		zip -r main.zip main