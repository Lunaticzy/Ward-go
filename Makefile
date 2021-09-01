LDFLAGS="-s -w"
OUT_DIR="build"

.PHONY: all
all: clean fmt test build

.PHONY: build
build: 
	@echo "build start..."; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(OUT_DIR)/ward_linux_amd64 main.go; \
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags $(LDFLAGS) -o $(OUT_DIR)/ward_linux_arm64 main.go; \
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -ldflags $(LDFLAGS) -o $(OUT_DIR)/ward_windows_amd64.exe main.go; \
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

.PHONY: upx
upx:
	@upx -9 $(OUT_DIR)/ward_*