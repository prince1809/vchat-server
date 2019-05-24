
dist: | check-style test package

build-linux:
	@echo Build Linux
	env GOOS=linux GOARCH=amd64 $(GO) install -i $(GOFLAGS) -ldflags '$(LDFLAGS)' ./...
