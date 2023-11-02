build:
	go build -o bin/main

run: build
	./bin/main


proto:
	protoc --proto_path=. --go_out=paths=source_relative:. \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
      		pb/*.proto \
          		pb/types/*.proto

.PHONY: proto