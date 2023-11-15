# workdir info
PACKAGE=gtstart
PREFIX=$(shell pwd)
CMD_PACKAGE=${PACKAGE}/command
OUTPUT_DIR=${PREFIX}/bin
OUTPUT_FILE=${OUTPUT_DIR}/gtstart
COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell git describe --tags || echo "v0.0.0")
VERSION_IMPORT_PATH=main
BUILD_TIME=$(shell date '+%Y-%m-%dT%H:%M:%S%Z')
VCS_BRANCH=$(shell git symbolic-ref --short -q HEAD)

# which golint
GOLINT=$(shell which golangci-lint || echo '')

# build args
BUILD_ARGS := \
    -ldflags "-X $(VERSION_IMPORT_PATH).appName=$(PACKAGE) \
    -X $(VERSION_IMPORT_PATH).version=$(VERSION) \
    -X $(VERSION_IMPORT_PATH).revision=$(COMMIT_ID) \
    -X $(VERSION_IMPORT_PATH).branch=$(VCS_BRANCH) \
    -X $(VERSION_IMPORT_PATH).buildDate=$(BUILD_TIME)"
EXTRA_BUILD_ARGS=

export GOPRIVATE=git.yottacloud.cn
export GOCACHE=

.PONY: lint test
default: dep generate lint test build

dep:
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/fatih/gomodifytags@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/golang/mock/mockgen@v1.6.0
	#go install git.yottacloud.cn/yottacloud/chaos/chaos@latest
	go install github.com/zeromicro/go-zero/tools/goctl@latest

generate:
	@echo "+ $@"
	go generate ./...

lint:
	@echo "+ $@"
	@$(if $(GOLINT), , \
		$(error Please install golint: `go install github.com/golangci/golangci-lint/cmd/golangci-lint`))
	golangci-lint --build-tags convenience run --deadline=10m ./...

test:
	@echo "+ test"
	go test -tags convenience -covermode=count -coverprofile=count.out $(EXTRA_BUILD_ARGS) ./...

coverage:
	@echo "+ $@"
	@if [ ! -f "count.out" ]; then make test ; fi
	go tool cover -func=count.out

build:
	@echo "+ build"
	go build -tags convenience $(BUILD_ARGS) $(EXTRA_BUILD_ARGS) -o ${OUTPUT_FILE} $(CMD_PACKAGE)

clean:
	@echo "+ $@"
	@if [ -d "bin/" ]; then rm -rf bin/ ; fi
	@if [ -d "design/gen/internal/" ]; then rm -rf design/gen/internal/ ; fi

release:
	@echo "+ $@"
	@make build
	@mkdir -p dist/
	@tar -zcvf dist/gtstart-${VERSION}.tar.gz README.md \
	    -C bin/ gtstart \
	    -C ../config/ config.sample.toml \
	    gtstart.service

goctl:
	@echo "+ $@"
	goctl api go --api ./design/*.api -dir ./design/gen/
	@cp -rf ./design/gen/internal/types ./design/gen/
	@rm -rf ./design/gen/etc

clean-goctl:
	@echo "+ $@"
	@if [ -d "design/gen/internal/" ]; then rm -rf design/gen/internal/ ; fi

goa-doc:
	@echo "+ $@"
	redoc-cli bundle ./design/gen/test.json -o ${OUTPUT_DIR}/apidoc.html


