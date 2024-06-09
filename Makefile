build:
	go build -ldflags="-s -w" .

lint: format
	golangci-lint run --fix

format: lines
	gofumpt -w .

lines: imports
	golines --max-len=80 -w .

imports:
	goimports-reviser -rm-unused -use-cache -format ./...

check-deps:
	@go list -u -m -json all | go-mod-outdated -update -direct

changelog:
	@git cliff --bump --unreleased

.PHONY: build lint format lines imports check-deps changelog
