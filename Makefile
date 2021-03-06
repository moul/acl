DOCKER_IMAGE ?=	moul/acl
SOURCES := $(shell find . -type f -name "*.go")

.PHONY: build
build: acl

acl: gen/pb/acl.pb.go $(SOURCES)
	go build -o acl ./cmd/acl

gen/pb/acl.pb.go: pb/acl.proto
	@mkdir -p gen/pb
	cd pb; protoc --gotemplate_out=destination_dir=../gen,template_dir=../vendor/github.com/moul/protoc-gen-gotemplate/examples/go-kit/templates/{{.File.Package}}/gen:../gen ./acl.proto
	gofmt -w ./gen
	cd pb; protoc --gogo_out=plugins=grpc:../gen/pb ./acl.proto
	mv gen/pb/github.com/moul/acl/gen/pb/*.pb.go gen/pb

.PHONY: test
test:
	go test -v $(shell go list ./... | grep -v /vendor/)


.PHONY: install
install:
	go install ./cmd/acl

.PHONY: docker.build
docker.build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker.run
docker.run:
	docker run -p 8000:8000 -p 9000:9000 $(DOCKER_IMAGE)

.PHONY: docker.test
docker.test: docker.build
	docker run $(DOCKER_IMAGE) make test
