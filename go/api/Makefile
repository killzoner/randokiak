.PHONY: default
default: install mockgen;

install:
	go get \
		github.com/golang/mock/mockgen

#example of golan mock generation, not used for now
mockgen:
	mockgen -source=controller/connectors.go -destination=mocks/connectors.go -package=mocks
	mockgen -source=controller/controller.go -destination=mocks/controller.go -package=mocks
