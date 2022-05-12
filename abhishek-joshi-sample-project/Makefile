download:
	go mod download
test-v: 
	go test -v -cover -covermode=atomic ./...
test:
	go test -cover -covermode=atomic ./...
swagger:
	swagger generate spec -o ./docs/static/resources/stickers.yaml
swagger-serve:
	swagger serve ./docs/static/resources/stickers.yaml
swagger-static-serve:
	swagger serve -F=swagger ./docs/static/resources/stickers.yaml
list:
	go list ./... | grep -v /vendor/
lint:
	golint -set_exit_status ./...
vet:
	go vet ./...
race:
	go test -race -short ./...
msan:
	go test -msan -short ./...
mocks:
	mockery --all
build:
	go build -o ${BINARY_NAME} app/*.go
docker-build:
	docker-compose up --build --no-start
start-mysql: docker-build
	docker-compose start mysql
start-redis: docker-build
	docker-compose start redis-server
start-app: docker-build
	docker-compose start app
restart-mysql:
	docker-compose restart mysql
restart-redis:
	docker-compose restart redis-server
restart-app: docker-build
	docker-compose restart app
stop-app:
	docker-compose stop app
stop-redis:
	docker-compose stop redis-server
stop-mysql:
	docker-compose stop mysql
live-reload:
	nodemon --exec go run app/main.go --signal SIGTERM
load-test:
	docker-compose -f locust-docker-compose.yml up locust-master
run: build
	./${BINARY_NAME}
