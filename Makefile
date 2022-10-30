proto-buff:
	@protoc --go_out=./service/src --proto_path=./proto ./proto/*.proto
	@protoc --go-grpc_out=./service/src --proto_path=./proto ./proto/*.proto

run-service:
	@echo "--> Starting servers"
	@docker-compose up

run-app:
	@echo "--> Starting frontend"
	cd ./frontend
	npm run serve