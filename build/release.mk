
dist: | check-style test package

build-linux:
	@echo Build Linux
	env GOOS=linux GOARCH=amd64 $(GO) install -i $(GOFLAGS) -ldflags '$(LDFLAGS)' ./...




ifeq ($(BUILDER_GOOS_GOARCH),"linux_amd64")
	cp $(GOPATH)/bin/mattermost $(DIST_PATH)/bin # from native bin dir, not cross-compiled
	cp $(GOPATH)/bin/mattermost $(DIST_PATH)/bin # From native bin dir, not cross-compiled
else
	cp $(GOPATH)/bin/linux_amd64/mattermost $(DIST_PATH)/bin # from cross-compiled bin dir
	cp $(GOPATH)/bin/linux_amd64/platform $(DIST_PATH)/bin # from cross-compiled
endif
	@# Package
	tar -C dist -czf $(DIST_PATH)-$(BUILD_TYPE_NAME)-linux-amd64.tar.gz mattermost
	@# Don't clean up native package so dev machines will have an unzipped package available
	@# rm -f $(DIST_PATH)/bin/mattermost

