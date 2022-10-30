build:
	protoc --go_out=./src --proto_path=../proto ../proto/*.proto
	protoc --go-grpc_out=./src --proto_path=../proto ../proto/*.proto

clean:
	rm -r src/gen

run:
	$(MAKE) build
	cd ./src && go run main.go