clean:
		rm main
		rm main.zip
build:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
		zip -r main.zip main

first-deploy:
		#Create User With Permissions to Deploy way CLI
		#Download Credentials key
		#Create and downloading access key
		#Download AWS CLI
		#Configure AWS CLI <aws configure> key id, access key, region name, output format
		echo "Deploy in AWS"

		echo "Create ROL"
		aws iam create-role --role-name lambda-basic-execution --assume-role-policy-document file://lambda-trust-policy.json
		aws iam attach-role-policy --role-name lambda-basic-execution --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
		aws lambda create-function --function-name helloworld_go --zip-file fileb://main.zip --handler main --runtime go1.x --role "arn:aws:iam::860066115873:role/lambda-basic-execution"

deploy:
		aws lambda update-function-code --function-name helloworld_go --zip-file fileb://main.zip

refresh: clean build deploy
all: build deploy