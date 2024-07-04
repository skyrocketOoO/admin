run-postgres:
	docker run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password postgres:latest

build-img:
	docker build -t go-server-template .

run-container:
	docker run -d --name go-server-template go-server-template

backup:
	./scripts/add_gitkeep.sh
	golines . -w -m 93
	gofumpt -w .
	
	git add .
	git commit -m "backup"
	git push

gen-rest-doc:
	swag init -g internal/delivery/rest/*.go -o ./
	
gen-proto:
	protoc --go_out=. --go-grpc_out=. manifest/protobuf/*