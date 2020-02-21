.PHONY: default
default: generateAll;

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

PROTOS_OUT_DIR = src/app/models/protos
PROTOS_GO_PATH = go/protocols
PROTOCOLS = $(wildcard $(ROOT_DIR)/../$(PROTOS_GO_PATH)/*.proto) # with path
PROTOCOLS_PATH_OUT = $(notdir $(PROTOCOLS))

root:
	echo $(ROOT_DIR) 
	echo $(PROTOCOLS)
	echo $(PROTOCOLS_PATH_OUT)

generateTest:
	protoc \
		-I $(ROOT_DIR)/../go/protocols \
		--plugin=./node_modules/.bin/protoc-gen-ts_proto \
		--ts_proto_out=./src/app/models/protos \
		randomprofile.proto \
		&& sed -i 's/global.//g' $(ROOT_DIR)/src/app/models/protos/randomprofile.ts
		#remove global node references 

generate:
	protoc \
		-I $(ROOT_DIR)/../$(PROTOS_GO_PATH) \
		--plugin=./node_modules/.bin/protoc-gen-ts_proto \
		--ts_proto_out=./$(PROTOS_OUT_DIR) \
		${file}

generateAll:
	for elt in ${PROTOCOLS}; do \
		make generate file=$$elt; \
	done; \
	sed -i 's/global.//g' $(ROOT_DIR)/$(PROTOS_OUT_DIR)/*.ts

