compress:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/main cmd/main.go
		zip -jrm build/main.zip build/main
clean:
		rm build/main.zip

deploy:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main cmd/main.go
		sam deploy --stack-name go-serverless-sam --region us-east-1 --resolve-s3 --capabilities CAPABILITY_IAM  
		rm main