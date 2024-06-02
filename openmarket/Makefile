.PHONY: build

build:
	sam build

local:
	sam build && sam local start-api --env-vars .env.json

deploy:
	sam build && sam package && yes | sam deploy 
