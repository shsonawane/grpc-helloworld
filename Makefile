proto-buff:
	@protoc --go_out=./service/src --proto_path=./proto ./proto/*.proto
	@protoc --go-grpc_out=./service/src --proto_path=./proto ./proto/*.proto
	@protoc --js_out=import_style=commonjs,binary:./app/src/gen/ --proto_path=./proto ./proto/*.proto
	@protoc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./app/src/gen/ --proto_path=./proto ./proto/*.proto

run-service:
	@echo "--> Starting service"
	@docker-compose up --abort-on-container-exit --build

run-service-prod:
	@echo "--> Starting service"
	@docker-compose -f docker-compose-prod.yaml up  --build

run-app:
	@echo "--> Starting web app"
	@cd ./app && yarn install && yarn clean && yarn build && yarn serve