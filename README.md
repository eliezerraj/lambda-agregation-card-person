GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

zip -jrm ../build/main.zip ../build/main

aws lambda update-function-code \
--function-name lambda-agregation-card-person \
--zip-file fileb:///mnt/c/Eliezer/workspace/github.com/lambda-agregation-card-person/build/main.zip \
--publish 

//------------------------

Endpoints

Get /version
Get /personaddress/{004}
Get /agregation/AGREGATION-4444.000.000.010/PERSON:PERSON-010

obs: Data agregated via lambda-agregation-card-person-worker
//-----

APIGW ==> Lambda ==> DynamoDB (agregation-card-person)

//-----