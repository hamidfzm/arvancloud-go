
.PHONY: install
install:
	yarn install

.PHONY: build
build:
	yarn build
	$(MAKE) go-tidy
	$(MAKE) go-fmt

MODULES := $(shell find . -type f -name 'go.mod' -printf '%h\n')

.PHONY: go-tidy
go-tidy:
	@for dir in $(MODULES); do \
		cd $$dir && go mod tidy && cd - > /dev/null; \
	done

.PHONY: go-fmt
go-fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./cdn/test ./iaas/test ./live/test ./paas/test ./storage/test ./vads/test ./vod/test
