PROJECT?=github.com/rumyantseva/tenerife
VERSION?=0.0.1

test:
	go test --race ./...

build:
	go build \
		-ldflags "-s -w -X ${PROJECT}/internal/diagnostics.Version=${VERSION} -X ${PROJECT}/internal/diagnostics.Commit=${COMMIT}" \
		-o bin/tenerife ${PROJECT}/cmd/tenerife
