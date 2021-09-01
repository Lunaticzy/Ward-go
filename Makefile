export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct

LDFLAGS := "-s -w"
OUT_DIR := "build"


.PHONY: all
all: clean fmt test build


.PHONY: build
build: 
	@echo "build start..."; \
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux  GIN_MODE=release go build -ldflags $(LDFLAGS)  -o $(OUT_DIR)/ward_linux_arm64 main.go; \
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows  GIN_MODE=release go build -ldflags $(LDFLAGS)  -o $(OUT_DIR)/ward_windows_amd64.exe main.go; \
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux  GIN_MODE=release go build -ldflags $(LDFLAGS)  -o $(OUT_DIR)/ward_linux_amd64 main.go; \
	cp ward.ini $(OUT_DIR); \
	echo "build finish..."

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: clean
clean: 
	@echo "clean..."; \
	rm -rf $(OUT_DIR);

.PHONY: test
test:
	@go test -v
