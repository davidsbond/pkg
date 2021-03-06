.DEFAULT_GOAL := all

# Downloads and vendors dependencies
modules:
	go mod tidy
	go mod vendor

# Formats all go source code
format:
	grep -L -R "Code generated .* DO NOT EDIT" --exclude-dir=.git --exclude-dir=vendor --include="*.go" | \
	xargs -n 1 gofumports -w -local pkg.dsb.dev/

# Runs go tests
test:
	go test -race ./...

# Installs go tooling
install-tools:
	go install \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		mvdan.cc/gofumpt/gofumports \
		github.com/sebdah/markdown-toc \
		github.com/bufbuild/buf/cmd/buf \
		github.com/golang/protobuf/protoc-gen-go

# Lints go source code
lint:
	golangci-lint run --enable-all
	buf lint

# Generates go source code
generate:
	markdown-toc --skip-headers=2 --replace --inline README.md
	go generate -x ./...
	buf generate

# Checks for any changes, including new files
has-changes:
	git add .
	git diff --staged --exit-code
