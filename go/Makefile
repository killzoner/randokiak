.PHONY: default
default: install go2proto generateAll;

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

PROTOCOLS = $(wildcard ./protocols/*.proto)
GOTOPROTOS = $(wildcard ./gotoproto/**/**/*.go)
ROOT_GITHUB = github.com/killzoner/randokiak

root:
	echo $(ROOT_DIR) 
	echo $(PROTOCOLS) 
	echo $(GOTOPROTOS) 
	echo $$GOPATH

generate:
	protoc \
		-I protocols \
		-I gotoproto \
		-I ${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		-I ${GOPATH}/src/github.com/gogo/protobuf/gogoproto \
		--gogoslick_out=plugins=grpc:$(ROOT_DIR) \
		${file}

# generatefast:
# 	protoc \
# 		-I protocols\
# 		-I gotoproto\
# 		-I ${GOPATH}/src/github.com/gogo/protobuf/protobuf\
# 		-I ${GOPATH}/src/github.com/gogo/protobuf/gogoproto\
# 		--gogoslick_out=plugins=grpc:$(ROOT_DIR)\
# 		./protocols/random.proto

# generatefaster:
# 	protoc \
# 		-I protocols \
# 		--gogofaster_out=plugins=grpc:$(ROOT_DIR)\
# 		protocols/helloworld.proto

# generateslick:
# 	protoc \
# 		-I protocols \
# 		--gogoslick_out=plugins=grpc:$(ROOT_DIR)\
# 		protocols/helloworld.proto

generateAll:
	for elt in ${PROTOCOLS}; do \
		make generate file=$$elt; \
	done

go2proto: 
	go2proto -f ./gotoproto -p ./gotoproto \
	&& cd ./gotoproto \
	&& mv output.proto randomprofile.proto \
	&& sed -i 's/package proto;/package randomprofile;\n\noption go_package = "protocols;protocols";/g' randomprofile.proto \
	&& sed -i '/interface/d' randomprofile.proto \
	&& mv randomprofile.proto ../protocols/randomprofile.proto

# to be run with project in GOPATH, since proteus does not support Go modules
# proteus seems to generate proto v2, don't know why, so exit
# proteus:
# 	proteus proto -f ./gotoproto \
# 		-p $(ROOT_GITHUB)/go/gotoproto/basicprotos/randomprofile \
# 		--verbose

install:
	go get \
		github.com/gogo/protobuf/protoc-gen-gogo \
		github.com/gogo/protobuf/protoc-gen-gofast \
		github.com/gogo/protobuf/protoc-gen-gogofaster \
		github.com/gogo/protobuf/protoc-gen-gogoslick \
		github.com/gogo/protobuf/proto \
		github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
		github.com/anjmao/go2proto