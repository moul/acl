DOCKER_IMAGE ?=	moul/acl
SOURCES := cmd/acl/main.go service/service.go

.PHONY: build
build: acl

acl: gen/pb/acl.pb.go gen/.generated $(SOURCES) $(MO_FILES)
	go build -o acl ./cmd/acl

gen/pb/acl.pb.go: pb/acl.proto
	@mkdir -p gen/pb
	cd pb; protoc --gogo_out=plugins=grpc:../gen/pb ./acl.proto
	mv gen/pb/github.com/moul/acl/gen/pb/*.pb.go gen/pb

gen/.generated:	pb/acl.proto
	@mkdir -p gen
	cd pb; protoc --gotemplate_out=destination_dir=../gen,template_dir=../vendor/github.com/moul/protoc-gen-gotemplate/examples/go-kit/templates/{{.File.Package}}/gen:../gen ./acl.proto
	@touch gen/.generated

.PHONY: install
install:
	go install ./cmd/acl

.PHONY: docker.build
docker.build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker.run
docker.run:
	docker run -p 8000:8000 -p 9000:9000 $(DOCKER_IMAGE)
