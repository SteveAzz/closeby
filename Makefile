BINARY = closeby
GOARCH = amd64
BUILD_DIR = builds
CURRENT_DIR = $(shell pwd)

linux:
	docker run \
	-v ${CURRENT_DIR}:/go/src/github.com/steveazz/closeby \
	-w /go/src/github.com/steveazz/closeby \
	-e GOOS=linux \
	-e GOARCH=${GOARCH} \
	golang:1.9.2-alpine3.7 \
	go build -o ${BUILD_DIR}/linux/${BINARY} cmd/closeby/closeby.go

darwin:
	docker run \
	-v ${CURRENT_DIR}:/go/src/github.com/steveazz/closeby \
	-w /go/src/github.com/steveazz/closeby \
	-e GOOS=darwin \
	-e GOARCH=${GOARCH} \
	golang:1.9.2-alpine3.7 \
	go build -o ${BUILD_DIR}/darwin/${BINARY} cmd/closeby/closeby.go

windows:
	docker run \
	-v ${CURRENT_DIR}:/go/src/github.com/steveazz/closeby \
	-w /go/src/github.com/steveazz/closeby \
	-e GOOS=windows \
	-e GOARCH=${GOARCH} \
	golang:1.9.2-alpine3.7 \
	go build -o ${BUILD_DIR}/darwin/${BINARY}.exe cmd/closeby/closeby.go

make test:
	docker run \
	-v ${CURRENT_DIR}:/go/src/github.com/steveazz/closeby \
	-w /go/src/github.com/steveazz/closeby \
	golang:1.9.2-alpine3.7 \
	go test -v ./...

