genSearch:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    rpc/search/search.proto

genParser:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    rpc/parser/parser.proto

genAll: genSearch genParser

build:
	go build -o bin/server cmd/server/main.go

run: build
	./bin/server

jsonParser:
	curl -v --request GET \
  	--url http://localhost:8083/v1/parse/ \
	--header 'authorization: Bearer jwt' \
	--header 'user-agent: vscode-restclient'

search:
	curl -v --request GET \
  	--url http://localhost:8083/v1/search/3 \
	--header 'authorization: Bearer jwt' \
	--header 'user-agent: vscode-restclient'

delete:
	curl -v --request DELETE \
  	--url http://localhost:8083/v1/search/3 \
	--header 'authorization: Bearer jwt' \
	--header 'user-agent: vscode-restclient'

insert:
	curl -v --request POST \
  	--url http://localhost:8083/v1/search/ \
	--header 'authorization: Bearer jwt' \
	--data '{"number": 3}'

checkEnvExists:
    ifeq (,$(wildcard .env))
        $(error .env file does not exist)
    endif

loadEnv:
	export $(xargs < .env)

dockerBuildRun: checkEnvExists
	docker build -t gateway . && \
	docker run --rm -it -p 8083:8083 --env-file .env gateway

dockerPush: checkEnvExists loadEnv
	docker build -t gateway . && \
	docker tag gateway $(DOCKER_REGISTRY) && \
	docker push $(DOCKER_REGISTRY):latest